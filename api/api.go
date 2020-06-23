package api

import (
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"

	"github.com/InVisionApp/rye"
	"github.com/pkg/errors"
	"github.com/sirupsen/logrus"
	"github.rakops.com/BNP/DisplayInvoiceGen/config"
	"github.rakops.com/BNP/DisplayInvoiceGen/deps"
	"github.rakops.com/BNP/DisplayInvoiceGen/log"
)

type contextKey int

const ContextBillingDate contextKey = 1
const ContextBillingTime contextKey = 2

type Api struct {
	Config  *config.Config
	Version string
	Deps    *deps.Dependencies
}

func New(cfg *config.Config, version string, deps *deps.Dependencies) (*Api, error) {

	return &Api{
		Config:  cfg,
		Version: version,
		Deps:    deps,
	}, nil
}

func (a *Api) Run() {
	srv := &http.Server{
		Addr:         a.Config.ListenAddress,
		Handler:      a.CreateRouter(),
		WriteTimeout: 60 * time.Second,

		ReadTimeout:    5 * time.Second,
		MaxHeaderBytes: 1024 * 1024,
	}

	go func() {
		logrus.Info("Api: running...")
		if err := srv.ListenAndServe(); err != nil {
			logrus.Errorf("HTTP server exit: %v", err)
		}
	}()

	signalChan := make(chan os.Signal, 1)
	cleanupDone := make(chan bool)
	signal.Notify(signalChan, syscall.SIGINT, syscall.SIGTERM, syscall.SIGQUIT, syscall.SIGKILL)
	go func() {
		var once sync.Once
		for range signalChan {
			once.Do(func() {

				a.Deps.Postgres.Close()
				a.Deps.Consumer.Shutdown()
				log.Warnf("Tax calculator stopped successfully")

				cleanupDone <- true
			})
		}
	}()
	log.Info("Tax calculator started successfully")
	<-cleanupDone
}

type loggedResponseWriter struct {
	w          http.ResponseWriter
	statusCode int
}

func (w *loggedResponseWriter) Header() http.Header         { return w.w.Header() }
func (w *loggedResponseWriter) Write(d []byte) (int, error) { return w.w.Write(d) }
func (w *loggedResponseWriter) WriteHeader(status int) {
	w.w.WriteHeader(status)
	w.statusCode = status
}

func wrapHandle(path string, handler http.Handler) (string, http.Handler) {
	return path, http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		startedAt := time.Now()
		loggedRW := &loggedResponseWriter{w: w}

		defer func() {
			if err := recover(); err != nil {
				statusCode := http.StatusInternalServerError
				rye.WriteJSONStatus(w, "error", fmt.Sprintf("internal server error: %v", err), statusCode)
			}
		}()

		handler.ServeHTTP(loggedRW, r)

		requestID := r.Header.Get("request-id")
		requestSource := r.Header.Get("request-source")
		callingService := r.Header.Get("calling-service")

		logrus.WithFields(logrus.Fields{
			"request_id":      requestID,
			"request_source":  requestSource,
			"calling_service": callingService,
			"duration":        time.Since(startedAt).Seconds() * 1000,
			"statusCode":      loggedRW.statusCode,
		}).Info("Request completed")
	})
}

func (a *Api) middlewareParseBillingDate(rw http.ResponseWriter, r *http.Request) *rye.Response {
	ctx := r.Context()

	billingDate := r.URL.Query().Get("billing_date")

	if len(billingDate) != len("0000-00-00") {
		return &rye.Response{
			Err:        fmt.Errorf("invalid biling data length: %v, expected date format '2020-02-20'", billingDate),
			StatusCode: http.StatusBadRequest,
		}
	}

	billingTime, err := time.Parse("2006-01-02", billingDate)
	if err != nil {
		return &rye.Response{
			Err:        err,
			StatusCode: http.StatusBadRequest,
		}
	}

	ctx = context.WithValue(ctx, ContextBillingDate, billingDate)
	ctx = context.WithValue(ctx, ContextBillingTime, billingTime)

	return &rye.Response{Context: ctx}
}

func readBody(message interface{}, r *http.Request) *rye.Response {
	defer r.Body.Close()
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		return BadRequestResponse(err, "unable to read request body")
	}

	if err = json.Unmarshal(body, message); err != nil {
		return BadRequestResponse(err, "invalid json message")
	}
	return nil
}

func ServerErrorResponse(err error, message string) *rye.Response {
	return &rye.Response{
		Err:        errors.Wrap(err, message),
		StatusCode: http.StatusInternalServerError,
	}
}

func BadRequestResponse(err error, message string) *rye.Response {
	return &rye.Response{
		Err:        errors.Wrap(err, message),
		StatusCode: http.StatusBadRequest,
	}
}

func NotFoundResponse(err error, message string) *rye.Response {
	return &rye.Response{
		Err:        errors.Wrap(err, message),
		StatusCode: http.StatusNotFound,
	}
}

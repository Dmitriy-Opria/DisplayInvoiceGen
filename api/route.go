package api

import (
	"net/http"

	"github.com/InVisionApp/rye"
	"github.com/gorilla/mux"
	"github.com/sirupsen/logrus"
)

func (a *Api) CreateRouter() http.Handler {
	middlewareHandler := rye.NewMWHandler(rye.Config{})

	routes := mux.NewRouter().StrictSlash(true)

	routes.Handle(wrapHandle("/healthcheck", middlewareHandler.Handle([]rye.Handler{
		a.healthHandler,
	}))).Methods("GET")

	//////////////////////////////////////////////////////////////////////////////
	// Exposed
	//////////////////////////////////////////////////////////////////////////////

	serviceName := "/invoice_gen"

	// healthcheck

	routes.Handle(wrapHandle(serviceName+"/healthcheck", middlewareHandler.Handle([]rye.Handler{
		a.healthHandler,
	}))).Methods("GET")

	// version
	routes.Handle(wrapHandle(serviceName+"/version", middlewareHandler.Handle([]rye.Handler{
		MiddlewareRouteLogger(),
		a.versionHandler,
	}))).Methods("GET")

	// create_invoice
	routes.Handle(wrapHandle(serviceName+"/create_invoice", middlewareHandler.Handle([]rye.Handler{
		MiddlewareRouteLogger(),
		a.middlewareParseBillingDate,
		a.createInvoice,
	}))).Methods("POST")

	return routes
}

func MiddlewareRouteLogger() func(rw http.ResponseWriter, req *http.Request) *rye.Response {
	return func(rw http.ResponseWriter, r *http.Request) *rye.Response {
		logrus.Infof("%s \"%s %s %s\"", r.RemoteAddr, r.Method, r.RequestURI, r.Proto)
		return &rye.Response{Context: r.Context()}
	}
}

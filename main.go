package main

import (
	"fmt"
	"net"
	"net/http"
	"os"
	"time"

	"github.rakops.com/BNP/DisplayInvoiceGen/api"
	"github.rakops.com/BNP/DisplayInvoiceGen/config"
	"github.rakops.com/BNP/DisplayInvoiceGen/deps"
	"github.rakops.com/BNP/DisplayInvoiceGen/log"
	"github.rakops.com/BNP/DisplayInvoiceGen/postgres"
	"github.rakops.com/BNP/DisplayInvoiceGen/queue"
	"github.rakops.com/BNP/DisplayInvoiceGen/rabbit/consumer"
	"github.rakops.com/BNP/DisplayInvoiceGen/rabbit/producer"
	"github.rakops.com/BNP/DisplayInvoiceGen/services"
)

var (
	user  string
	pass  string
	addr  string
	name  string
	token string
)

func main() {
	initLogger()
	cfg := initConfig()
	pg := initPostgres(cfg)
	consumer := initConsumer(cfg)
	producer := initProducer(cfg)

	httpClient := &http.Client{
		Timeout: time.Duration(60 * time.Second),
		Transport: &http.Transport{
			Proxy: http.ProxyFromEnvironment,
			DialContext: (&net.Dialer{
				Timeout:   60 * time.Second,
				KeepAlive: 60 * time.Second,
				DualStack: true,
			}).DialContext,
			MaxIdleConns:          100,
			MaxIdleConnsPerHost:   100,
			IdleConnTimeout:       180 * time.Second,
			TLSHandshakeTimeout:   10 * time.Second,
			ExpectContinueTimeout: 5 * time.Second,
		},
	}

	externalService := services.NewExternalService(cfg, httpClient)

	dependencies := &deps.Dependencies{
		Postgres:        pg,
		ExternalService: externalService,
		Consumer:        consumer,
		Producer:        producer,
		Config:          cfg,
		Version:         cfg.Version,
	}

	queue := queue.NewQueueHandler(dependencies)

	err := queue.Run()
	if err != nil {
		log.Fatalf("Unable to run Queue: %v", err)
		return
	}

	api, err := api.New(cfg, cfg.Version, dependencies)
	if err != nil {
		log.Fatalf("Unable to start API: %v", err)
		return
	}

	api.Run()
}

func initConfig() *config.Config {
	cfg := config.InitConfig()
	if user != "" &&
		pass != "" &&
		addr != "" &&
		name != "" &&
		token != "" {
		cfg.Postgres.User = user
		cfg.Postgres.Pass = pass
		cfg.Postgres.Addr = addr
		cfg.Postgres.Database = name
		cfg.TaxCalculationService.AuthToken = token
	}

	log.Info("Config was successfully initialized")
	return cfg
}
func initLogger() {
	logger, err := log.NewServiceLogger(log.Fields{
		log.FieldApplication: "invoice_gen",
		log.FieldService:     "invoice_gen",
		log.FieldCategory:    "VAT/GST",
	}, log.LevelDebug)

	if err != nil {
		fmt.Fprint(os.Stdout, err)
		os.Exit(1)
	}
	log.SetLogger(logger)
	log.Info("Logger was successfully initialized")
}

func initPostgres(cfg *config.Config) postgres.IConnection {
	conn := postgres.NewConnection(cfg.Postgres)
	if err := conn.Connect(); err != nil {
		log.Fatalf("Can't connect db, %v", err.Error())
	}
	log.Info("Postgres was successfully initialized")
	return conn
}

func initConsumer(cfg *config.Config) consumer.RabbitConsumer {
	conn := consumer.NewConsumer(cfg)
	if err := conn.Connect(); err != nil {
		log.Fatalf("Can't connect db, %v", err.Error())
	}
	log.Info("RabbitMQ consumer was successfully initialized")
	return conn
}

func initProducer(cfg *config.Config) producer.RabbitProducer {
	conn := producer.NewProducer(cfg)
	if err := conn.Connect(); err != nil {
		log.Fatalf("Can't connect db, %v", err.Error())
	}
	log.Info("RabbitMQ producer was successfully initialized")
	return conn
}

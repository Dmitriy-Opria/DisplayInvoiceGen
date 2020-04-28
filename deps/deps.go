package deps

import (
	"github.rakops.com/BNP/DisplayInvoiceGen/config"
	"github.rakops.com/BNP/DisplayInvoiceGen/postgres"
	"github.rakops.com/BNP/DisplayInvoiceGen/rabbit/consumer"
	"github.rakops.com/BNP/DisplayInvoiceGen/rabbit/producer"
	"github.rakops.com/BNP/DisplayInvoiceGen/services"
)

type Dependencies struct {
	Postgres        postgres.IConnection
	ExternalService services.IExternalService
	Consumer        consumer.RabbitConsumer
	Producer        producer.RabbitProducer
	Config          *config.Config
	Version         string
}

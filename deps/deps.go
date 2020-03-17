package deps

import (
	"github.rakops.com/BNP/DisplayInvoiceGen/config"
	"github.rakops.com/BNP/DisplayInvoiceGen/postgres"
	"github.rakops.com/BNP/DisplayInvoiceGen/services"
)

type Dependencies struct {
	Postgres        postgres.IConnection
	ExternalService services.IExternalService
	Config          *config.Config
	Version         string
}

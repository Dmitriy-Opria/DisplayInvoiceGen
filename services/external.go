package services

import (
	"github.rakops.com/BNP/DisplayInvoiceGen/config"
	"net/http"
)

const DefaultTaxRate = 20

type IExternalService interface {
	GetTaxRate() int64
}

type ExternalService struct {
	config *config.Config
	client *http.Client
}

func NewExternalService(config *config.Config, client *http.Client) IExternalService {
	return &ExternalService{
		config: config,
		client: client,
	}
}

func (e *ExternalService) GetTaxRate() int64 {
	return DefaultTaxRate
}

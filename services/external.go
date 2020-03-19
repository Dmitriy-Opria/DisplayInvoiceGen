package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"github.rakops.com/BNP/DisplayInvoiceGen/config"
	"github.rakops.com/BNP/DisplayInvoiceGen/model"
	"io/ioutil"
	"net/http"
)

const DefaultTaxRate = 20

type IExternalService interface {
	GetTaxRate() (int64, error)
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

func (e *ExternalService) GetTaxRate() (int64, error) {
	// TODO add structure fields
	payload := model.TaxRequest{}

	requestBody, err := json.Marshal(payload)
	if err != nil {
		return 0, errors.Wrap(err, "creating request body")
	}

	req, err := http.NewRequest("POST", fmt.Sprintf("%v/v1/CalculateTax", e.config.TaxCalculationService.Address), bytes.NewBuffer(requestBody))
	if err != nil {
		return 0, errors.Wrap(err, "posting tax request failed")
	}

	req.Header.Add("Authorization", e.config.TaxCalculationService.AuthToken)
	req.Header.Add("Content-Type", "application/json")

	response, err := e.client.Do(req)
	if response != nil && response.Body != nil {
		defer response.Body.Close()
	}

	if err != nil {
		return 0, errors.Wrap(err, "posting tax api failed")
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return 0, errors.Wrap(err, "reading index body failed")
	}

	if response.StatusCode != 200 {
		return 0, fmt.Errorf("index request failed with code %v and body: %v", response.StatusCode, string(body))
	}

	result := &model.TaxRequest{}
	if err := json.Unmarshal(body, &result); err != nil {
		return 0, errors.Wrap(err, "parsing result failed")
	}
	return 0, nil
}

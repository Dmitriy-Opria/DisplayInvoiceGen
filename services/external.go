package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"github.com/pkg/errors"
	"github.rakops.com/BNP/DisplayInvoiceGen/config"
	"github.rakops.com/BNP/DisplayInvoiceGen/log"
	"github.rakops.com/BNP/DisplayInvoiceGen/model"
	"github.rakops.com/BNP/DisplayInvoiceGen/postgres"
	"io/ioutil"
	"net/http"
	"strconv"
)

const DefaultTaxRate = 20

type IExternalService interface {
	GetTaxResponse(charges []*postgres.Charge, billingDate string) (*model.Response, error)
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

func (e *ExternalService) GetTaxResponse(charges []*postgres.Charge, billingDate string) (*model.Response, error) {

	payload := model.Invoice{
		InvoiceRequest: model.InvoiceRequest{
			Customer: model.Customer{
				Destination: model.Destination{
					Country: charges[0].BillingCountryCode,
				},
				Tax: []model.TaxRegistrationInvoice{
					{
						TaxRegistrationNumber:        charges[0].VATRegistrationNumber,
						HasPhysicalPresenceIndicator: "true",
						IsoCountryCode:               charges[0].BillingCountryCode,
					},
				},
			},
			Seller: model.Seller{
				Company: e.config.TaxCalculationParams.CompanyName,
				PhysicalOrigin: model.Destination{
					Country: charges[0].BillingCountryCode,
				},
				TaxRegistrationType: []model.TaxRegistrationInvoice{
					{
						TaxRegistrationNumber:        charges[0].VATRegistrationNumber,
						HasPhysicalPresenceIndicator: "true",
						IsoCountryCode:               charges[0].BillingCountryCode,
					},
				},
			},

			DocumentDate:    billingDate,
			TransactionType: e.config.TaxCalculationParams.TransactionType,
		},
	}

	if _, ok := e.config.TaxCalculationParams.DivisionAU[charges[0].Country]; ok {
		payload.InvoiceRequest.Seller.Division = "RM_AU"

	} else if _, ok := e.config.TaxCalculationParams.DivisionUS[charges[0].Country]; ok {
		payload.InvoiceRequest.Seller.Division = "RM_US"

	} else if _, ok := e.config.TaxCalculationParams.DivisionEU[charges[0].Country]; ok {
		payload.InvoiceRequest.Seller.Division = "RM_EU"

	} else {
		payload.InvoiceRequest.Seller.Division = ""

	}

	for index, line := range charges {
		payload.InvoiceRequest.Lines = append(payload.InvoiceRequest.Lines,
			model.LineItemInvoice{
				Product: model.ProductInvoice{
					ProductClass: e.config.TaxCalculationParams.ProductClass,
				},
				ExtendedPrice:  fmt.Sprintf("%.2f", line.ChargeAmount),
				LineItemNumber: strconv.Itoa(index + 1),
			})
	}

	requestBody, err := json.Marshal(payload)
	if err != nil {
		return nil, errors.Wrap(err, "creating request body")
	}

	log.Infof("request: %v", string(requestBody))

	req, err := http.NewRequest("POST", fmt.Sprintf("%v/v1/CalculateTax90", e.config.TaxCalculationService.Address), bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, errors.Wrap(err, "posting tax request failed")
	}

	req.Header.Add("Authorization", e.config.TaxCalculationService.AuthToken)
	req.Header.Add("Content-Type", "application/json")

	response, err := e.client.Do(req)
	if response != nil && response.Body != nil {
		defer response.Body.Close()
	}

	if err != nil {
		return nil, errors.Wrap(err, "posting tax api failed")
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, errors.Wrap(err, "reading index body failed")
	}

	if response.StatusCode != 200 {
		return nil, fmt.Errorf("index request failed with code %v and body: %v", response.StatusCode, string(body))
	}

	result := &model.Response{}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, errors.Wrap(err, "parsing result failed")
	}
	log.Infof("response: %v", string(body))
	return result, nil
}

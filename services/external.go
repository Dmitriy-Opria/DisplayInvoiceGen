package services

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
	"strings"
	"time"

	"github.com/jinzhu/now"
	"github.com/pkg/errors"
	"github.rakops.com/BNP/DisplayInvoiceGen/config"
	"github.rakops.com/BNP/DisplayInvoiceGen/log"
	"github.rakops.com/BNP/DisplayInvoiceGen/model"
	"github.rakops.com/BNP/DisplayInvoiceGen/postgres"
)

const Layout = "2006-01-02"

type IExternalService interface {
	GetTaxResponse(charges []*postgres.Charge) (*model.Response, error)
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

func (e *ExternalService) GetTaxResponse(charges []*postgres.Charge) (*model.Response, error) {

	payload := model.Invoice{
		InvoiceRequest: model.InvoiceRequest{
			Customer: model.Customer{
				Destination: model.Destination{
					Country: charges[0].BillingCountryCode,
				},
			},
			Seller: model.Seller{
				Company: e.config.TaxCalculationParams.CompanyName,
				PhysicalOrigin: model.Destination{
					Country: charges[0].RakutenCountry,
				},
			},

			DocumentDate:    now.New(time.Now()).BeginningOfMonth().Format(Layout), //time.now, first day of month
			TransactionType: e.config.TaxCalculationParams.TransactionType,
		},
	}

	if charges[0].CODAVATRegistrationNumber != "" {
		payload.InvoiceRequest.Customer.Tax = []model.TaxRegistrationInvoice{ //don't send if number empty!
			{
				TaxRegistrationNumber:        charges[0].CODAVATRegistrationNumber,
				HasPhysicalPresenceIndicator: "true", //if exist
				IsoCountryCode:               charges[0].BillingCountryCode,
			},
		}
	}

	if _, ok := e.config.TaxCalculationParams.RegistrationAU[charges[0].RakutenCountry]; ok {
		payload.InvoiceRequest.Seller.Division = "RM_AU"
		payload.InvoiceRequest.Seller.TaxRegistration = []model.TaxRegistrationInvoice{
			{
				TaxRegistrationNumber:        e.config.TaxCalculationParams.TaxIdAU,
				HasPhysicalPresenceIndicator: "true",
				IsoCountryCode:               e.config.TaxCalculationParams.RegistrationIsoAU,
			},
		}

	} else if _, ok := e.config.TaxCalculationParams.RegistrationUS[charges[0].RakutenCountry]; ok {
		payload.InvoiceRequest.Seller.Division = "RM_US"
		payload.InvoiceRequest.Seller.TaxRegistration = []model.TaxRegistrationInvoice{
			{
				TaxRegistrationNumber:        e.config.TaxCalculationParams.TaxIdUS,
				HasPhysicalPresenceIndicator: "true",
				IsoCountryCode:               e.config.TaxCalculationParams.RegistrationIsoUS,
			},
		}
	} else if _, ok := e.config.TaxCalculationParams.RegistrationEU[charges[0].RakutenCountry]; ok {
		payload.InvoiceRequest.Seller.Division = "RM_EU"
		payload.InvoiceRequest.Seller.TaxRegistration = []model.TaxRegistrationInvoice{
			{
				TaxRegistrationNumber:        e.config.TaxCalculationParams.TaxIdEU,
				HasPhysicalPresenceIndicator: "true",
				IsoCountryCode:               e.config.TaxCalculationParams.RegistrationIsoEU,
			},
		}

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
				LineItemID:     line.ChargeID,
			})
	}

	requestBody, err := json.Marshal(payload)
	if err != nil {
		return nil, errors.Wrap(err, "creating request body")
	}

	log.Infof("request: %v", string(requestBody))

	log.WithFields(log.Fields{
		"request": strings.Replace(string(requestBody), `\`, "", -1),
	}).Infof("billing_settings_request: %v", charges[0].BillingSettings)

	req, err := http.NewRequest("POST", fmt.Sprintf("%v/v1/CalculateTax90", e.config.TaxCalculationService.Address), bytes.NewBuffer(requestBody))
	if err != nil {
		return nil, errors.Wrap(err, "posting tax request failed")
	}

	req.Header.Add("Authorization", e.config.TaxCalculationService.AuthToken)
	req.Header.Add("Content-Type", "application/json")

	try := 0

Request:
	response, err := e.client.Do(req)
	if response != nil && response.Body != nil {
		defer response.Body.Close()
	}

	if (err != nil || (response != nil && response.StatusCode != http.StatusOK)) &&
		try < e.config.TaxCalculationService.Retry {
		try++
		log.Warnf("get to retry count: %v", try)
		goto Request
	}

	if err != nil {
		return nil, errors.Wrap(err, "posting tax api failed")
	}

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return nil, errors.Wrap(err, "reading index body failed")
	}

	if response.StatusCode != http.StatusOK {
		return nil, errors.New(fmt.Sprintf("invalid status code, expected 200, got: %v, cody %v", response.StatusCode, string(body)))
	}
	if response.StatusCode != 200 {
		return nil, fmt.Errorf("index request failed with code %v", response.StatusCode)
	}

	result := &model.Response{}
	if err := json.Unmarshal(body, &result); err != nil {
		return nil, errors.Wrap(err, "parsing result failed")
	}

	log.WithFields(log.Fields{
		"request_response": strings.Replace(string(body), `\`, "", -1),
	}).Infof("billing_settings_response: %v", charges[0].BillingSettings)

	return result, nil
}

package salesforce

import (
	"fmt"

	"github.com/pkg/errors"
	"github.rakops.com/BNP/DisplayInvoiceGen/log"
)

type InvoiceLineItem struct {
	InvoiceNumber       int64   `json:"-"`
	InvoiceID           string  `json:"Invoice__c"`
	LineItemTaxAmount   float64 `json:"Line_Item_Tax_Amount__c"`
	LineItemAmount      float64 `json:"Line_Item_Amount__c"`
	LineItemTotalAmount float64 `json:"Amount__c"`
	CurrencyIsoCode     string  `json:"CurrencyIsoCode"`
	Description         string  `json:"Description__c"`
	LineItemNumber      int64   `json:"Line_Number__c"`
}

func (s *SalesForce) GetInvoiceLineItemGobID() (string, error) {
	payload := map[string]string{
		"operation":       "insert",
		"object":          "Sales_Invoice_Line_Items__c",
		"contentType":     "JSON",
		"concurrencyMode": "Parallel",
	}

	out := make(map[string]interface{})
	err := s.request("POST", fmt.Sprintf("/services/async/%v.0/job", s.config.SalesForce.ApiVersion), payload, &out)

	if err != nil {
		return "", errors.Wrap(err, "can't execute invoice line item job request")
	}

	id, ok := out["id"]
	if !ok {
		return "", errors.New(fmt.Sprintf("can't get invoice line item job ID: %v", out))
	}
	return id.(string), err
}

func (s *SalesForce) PushInvoiceLineItemBatch(jobID string, payload []*InvoiceLineItem) (string, error) {
	out := make(map[string]interface{})
	err := s.request("POST", fmt.Sprintf("/services/async/%v.0/job/%v/batch", s.config.SalesForce.ApiVersion, jobID), payload, &out)

	if err != nil {
		return "", errors.Wrap(err, "can't post invoice line item push request")
	}
	log.Infof("output: %v", out)
	id, ok := out["id"]
	if !ok {
		return "", errors.New(fmt.Sprintf("can't get invoice line item batch id: %v", out))
	}
	return id.(string), nil
}

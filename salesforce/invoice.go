package salesforce

import (
	"fmt"

	"github.com/pkg/errors"
	"github.rakops.com/BNP/DisplayInvoiceGen/log"
)

type Invoice struct {
	Name            string  `json:"Name,omitempty"`
	InvoiceNumber   int64   `json:"Invoice_Number__c,omitempty"`
	BillingSettings string  `json:"GBS_Billing_Setting__c,omitempty"`
	InvoiceAmount   float64 `json:"Invoice_Amount__c,omitempty"`
	TaxTotal        float64 `json:"Tax_Total__c,omitempty"`
	Account         string  `json:"Account__c,omitempty"`
	SAPCustomerID   string  `json:"SAP_Customer_ID__c,omitempty"`
	CurrencyIsoCode string  `json:"CurrencyIsoCode,omitempty"`
	InvoiceDate     string  `json:"Invoiced_Date__c,omitempty"`
	InvoiceDueDate  string  `json:"Invoice_Due_Date__c,omitempty"`
	Company         string  `json:"Company__c,omitempty"`
	RecordTypeId    string  `json:"RecordTypeId,omitempty"`
	PONumber        string  `json:"PO_Number__c,omitempty"`
}

func (s *SalesForce) GetInvoiceGobID() (string, error) {
	payload := map[string]string{
		"operation":       "insert",
		"object":          "Invoice__c",
		"contentType":     "JSON",
		"concurrencyMode": "Parallel",
	}

	out := make(map[string]interface{})
	err := s.request("POST", fmt.Sprintf("/services/async/%v.0/job", s.config.SalesForce.ApiVersion), payload, &out)

	if err != nil {
		return "", errors.Wrap(err, "can't execute Job request")
	}

	id, ok := out["id"]
	if !ok {
		return "", errors.New(fmt.Sprintf("can't get Invoice job ID: %v", out))
	}
	return id.(string), err
}

func (s *SalesForce) PushInvoiceBatch(jobID string, payload []*Invoice) (string, error) {
	out := make(map[string]interface{})
	err := s.request("POST", fmt.Sprintf("/services/async/%v.0/job/%v/batch", s.config.SalesForce.ApiVersion, jobID), payload, &out)

	if err != nil {
		return "", errors.Wrap(err, "can't post Invoice job request")
	}
	log.Infof("output: %v", out)
	id, ok := out["id"]
	if !ok {
		return "", errors.New(fmt.Sprintf("can't get Invoice job ID: %v", out))
	}
	return id.(string), nil
}

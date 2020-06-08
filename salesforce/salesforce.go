package salesforce

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/nimajalali/go-force/force"
	"github.com/pkg/errors"
	"github.rakops.com/BNP/DisplayInvoiceGen/config"
	"github.rakops.com/BNP/DisplayInvoiceGen/log"
)

const (
	environment = "production"
)

type ISalesForceUploader interface {
	Auth() error

	CloseJob(jobID string) error
	GetStatus(jobID, batchID string) (Status, error)
	GetResult(jobID, batchID string) ([]*Result, error)
	PushToSalesForce(invoices []*Invoice, invoiceLineItems []*InvoiceLineItem) error

	GetInvoiceGobID() (string, error)
	PushInvoiceBatch(jobID string, payload []*Invoice) (string, error)

	GetInvoiceLineItemGobID() (string, error)
	PushInvoiceLineItemBatch(jobID string, payload []*InvoiceLineItem) (string, error)
}

type SalesForce struct {
	config *config.Config
	client *http.Client
	api    *force.ForceApi
	tm     time.Time
}

func NewSalesForce(config *config.Config, client *http.Client) ISalesForceUploader {
	return &SalesForce{
		config: config,
		client: client,
	}
}

func (s *SalesForce) Auth() error {

	env := "sandbox"
	if s.config.Production {
		env = environment
	}
	forceApi, err := force.Create(
		fmt.Sprintf("v%v.0", s.config.SalesForce.ApiVersion),
		s.config.SalesForce.ClientID,
		s.config.SalesForce.ClientSecret,
		s.config.SalesForce.UserName,
		s.config.SalesForce.Pass,
		s.config.SalesForce.SecurityToken,
		env,
	)
	if err != nil {
		log.Warnf("can't authenticate: %v", err)
		return err
	}
	s.api = forceApi
	return nil
}

type Result struct {
	Id      string        `json:"id"`
	Success bool          `json:"success"`
	Created bool          `json:"created"`
	Errors  []interface{} `json:"errors"`
}

type Status int

const (
	Unknown    = Status(0)
	Completed  = Status(1)
	InProgress = Status(2)
	Closed     = Status(3)
)

func (s *SalesForce) GetStatus(jobID, batchID string) (Status, error) {
	out := make(map[string]interface{})

	err := s.request("GET", fmt.Sprintf("/services/async/%v.0/job/%v/batch/%v", s.config.SalesForce.ApiVersion, jobID, batchID), nil, &out)

	if err != nil {
		return Unknown, errors.Wrap(err, "can't get response from Status salesforce")
	}
	state, ok := out["state"]
	if ok {
		switch state.(string) {
		case "Completed":
			return Completed, nil
		case "InProgress":
			return InProgress, nil
		case "Closed":
			return Closed, nil
		}

	}
	log.Infof("output: %v", out)

	return Unknown, nil
}

func (s *SalesForce) GetResult(jobID, batchID string) ([]*Result, error) {
	out := []*Result{}

	err := s.request("GET", fmt.Sprintf("/services/async/%v.0/job/%v/batch/%v/result", s.config.SalesForce.ApiVersion, jobID, batchID), nil, &out)

	if err != nil {
		return nil, errors.Wrap(err, "can't get response from Result salesforce")
	}
	log.Infof("output: %v", out)

	return out, nil
}

func (s *SalesForce) CloseJob(jobID string) error {

	payload := map[string]string{
		"state": "Closed",
	}

	out := make(map[string]interface{})
	err := s.request("POST", fmt.Sprintf("/services/async/%v.0/job/%v", s.config.SalesForce.ApiVersion, jobID), payload, &out)

	if err != nil {
		return errors.Wrap(err, "can't post Invoice job request")
	}
	log.Infof("output: %v", out)

	state, ok := out["state"]
	if !ok || state.(string) != "Closed" {
		return errors.New(fmt.Sprintf("can't get state from Close response: %v", out))
	}

	return nil
}

func (s *SalesForce) request(method, url string, payload, out interface{}) error {
	var reqBody io.Reader
	var try = 0
	if payload != nil {

		jsonBytes, err := json.Marshal(payload)
		if err != nil {
			return errors.Wrap(err, "can't marshal json body")
		}

		reqBody = bytes.NewReader(jsonBytes)
	}

	request, err := http.NewRequest(method, fmt.Sprintf("%v%v", s.api.GetInstanceURL(), url), reqBody)
	if err != nil {
		return errors.New("can't create request")
	}

	request.Header.Add("Content-type", "application/json")
	request.Header.Add("Accept", "application/json")
	request.Header.Add("X-SFDC-Session", s.api.GetAccessToken())

Restart:
	response, err := s.client.Do(request)
	if err != nil {
		return errors.Wrap(err, "can't get response from salesforce")
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)
	if err != nil {
		return errors.Wrap(err, "can't read body from response")
	}

	if response.StatusCode == http.StatusBadRequest && try < s.config.SalesForce.Retry {
		try++
		s.api.RefreshToken()
		goto Restart
	}

	if response.StatusCode != http.StatusCreated && response.StatusCode != http.StatusOK {
		return errors.New(fmt.Sprintf("can't execute request, statusCode: %v, url: %v, body: %v", response.StatusCode, url, string(body)))
	}
	log.Info(string(body))

	err = json.Unmarshal(body, out)
	if err != nil {
		return errors.Wrap(err, "can't unmarshal Job response")
	}
	return nil
}

func (s *SalesForce) PushToSalesForce(invoices []*Invoice, invoiceLineItems []*InvoiceLineItem) error {

	invoiceJobID, err := s.GetInvoiceGobID()
	if err != nil {
		return errors.Wrap(err, "can't create invoice job id")
	}

	invoiceBatchID, err := s.PushInvoiceBatch(invoiceJobID, invoices)
	if err != nil {
		return errors.Wrap(err, "can't push invoice batch")
	}

	if err := s.waitForCompletion(invoiceJobID, invoiceBatchID); err != nil {
		return err
	}

	result, err := s.GetResult(invoiceJobID, invoiceBatchID)
	if err != nil {
		return errors.Wrap(err, "can't get invoice result")
	}

	if len(result) != len(invoices) {
		return errors.New("bad result of invoices pushing, length of a slices are different")
	}

	invoiceMapping := map[int64]string{}

	for index := range invoices {
		if !result[index].Success {
			return errors.New(fmt.Sprintf("failed invoice publishing result, sequence: %v, invoiceID: %v, status:%v", invoices[index].InvoiceNumber, result[index].Id, result[index].Success))
		}
		invoiceMapping[invoices[index].InvoiceNumber] = result[index].Id
	}

	for index := range invoiceLineItems {
		invoiceLineItems[index].InvoiceID = invoiceMapping[invoiceLineItems[index].InvoiceNumber]
	}

	invoiceLineItemJobID, err := s.GetInvoiceLineItemGobID()
	if err != nil {
		return errors.Wrap(err, "can't create invoice line item job id")
	}

	invoiceLineItemBatchID, err := s.PushInvoiceLineItemBatch(invoiceLineItemJobID, invoiceLineItems)
	if err != nil {
		return errors.Wrap(err, "can't push invoice batch")
	}

	if err := s.waitForCompletion(invoiceLineItemJobID, invoiceLineItemBatchID); err != nil {
		return err
	}

	lineItemResult, err := s.GetResult(invoiceLineItemJobID, invoiceLineItemBatchID)
	if err != nil {
		return errors.Wrap(err, "can't get invoice line item result")
	}
	if len(lineItemResult) != len(invoiceLineItems) {
		return errors.New("bad result of invoices line item pushing, length of a slices are different")
	}

	return nil
}

func (s *SalesForce) waitForCompletion(jobID, batchID string) error {

	status := Unknown
	var err error
	var try int
CheckStatus:
	if status != Completed && try < 10 {
		try++
		time.Sleep(time.Second)
		status, err = s.GetStatus(jobID, batchID)
		log.Infof("status: %v", status)
		if err != nil {
			return errors.Wrap(err, "can't get result of invoice pushing")
		}
		goto CheckStatus
	}
	return nil
}

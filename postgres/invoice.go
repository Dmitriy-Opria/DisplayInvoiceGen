package postgres

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
	"time"

	"github.com/jinzhu/now"
	"github.rakops.com/BNP/DisplayInvoiceGen/log"
)

type Invoice struct {
	tableName       struct{}   `sql:"public.invoice"`
	InvoiceNumber   int64      `json:"InvoiceNumber"          sql:"invoicenumber"`  // InvoiceSequence
	BillingSetting  string     `json:"BillingSetting"         sql:"billingsetting"` // GBS_Billing_Setting__c
	InvoiceAmount   float64    `json:"InvoiceAmount"          sql:"invoiceamount"`
	TaxRate         float64    `json:"TaxRate"                sql:"taxrate,notnull,type:numeric"`
	AccountNumber   string     `json:"AccountNumber"          sql:"accountnumber"`   // Account__c
	SAPCustomerID   string     `json:"SAPCustomerID"          sql:"sapcustomerid"`   // SAP_Customer_ID__c
	PaymentTerms    string     `json:"PaymentTerms"           sql:"paymentterms"`    // Payment_Terms_SAP__c
	InvoiceCurrency string     `json:"InvoiceCurrency"        sql:"invoicecurrency"` // charge_currency
	InvoiceEmailed  bool       `json:"InvoiceEmailed"         sql:"invoiceemailed"`
	InsertDate      time.Time  `json:"InsertDate"             sql:"insertdate"`      // System Date
	UpdateDate      *time.Time `json:"UpdateDate"             sql:"updatedate"`      // Null
	EmailDate       *time.Time `json:"EmailDate"              sql:"emaildate"`       // Null
	InvoiceDueDate  string     `json:"InvoiceDueDate"         sql:"invoiceduedate"`  // If a invoice is generated in March, then base on "Payment_Terms_SAP__c" the InvoiceDueDate should be calculated -- 31 MAR 2020
	InvoicePeriod   string     `json:"InvoicePeriod"          sql:"invoiceperiod"`   // InvoiceDueDate - 1 Month -- Feb 2020
	InvoicePostDate string     `json:"InvoicePostDate"        sql:"invoicepostdate"` // Invoice Run Date -- 10 MAR 2020
	BillingDate     string     `json:"BillingDate"            sql:"billingdate"`     // InvoiceDueDate + 1 or Invoice Run Month -- 1 MAR 2020
	PDFnumber       string     `json:"PDFnumber"              sql:"pdfnumber"`       // PDFnumber SIN002022
	CompanyCountry  string     `json:"CompanyCountry"         sql:"companycountry"`  // GB/US/AU
	CustomerCountry string     `json:"CustomerCountry"        sql:"customercountry"` // UK/USA/AU
}

type InvoiceUp struct {
	Invoice

	TaxTotal    float64 `json:"TaxTotal"               sql:"taxtotal"`
	IsApproved  bool    `json:"IsApproved"             sql:"isapproved"` // true/false
	PONumber    string  `json:"PoNumber"              sql:"po_number"`   // true/false
	CompanyName string  `json:"CompanyName"            sql:"name"`       // Rakuten Marketing LLC
}

func (p *ConnectionWrapper) AddInvoice(invoice *Invoice) error {
	_, err := p.client.Model(invoice).
		Set("invoicenumber = EXCLUDED.invoicenumber").
		Set("billingsetting = EXCLUDED.billingsetting").
		Set("invoiceamount = EXCLUDED.invoiceamount").
		Set("taxrate = EXCLUDED.taxrate").
		Set("accountnumber = EXCLUDED.accountnumber").
		Set("sapcustomerid = EXCLUDED.sapcustomerid").
		Set("paymentterms = EXCLUDED.paymentterms").
		Set("invoicecurrency = EXCLUDED.invoicecurrency").
		Set("invoiceemailed = EXCLUDED.invoiceemailed").
		Set("insertdate = EXCLUDED.insertdate").
		Set("updatedate = EXCLUDED.updatedate").
		Set("emaildate = EXCLUDED.emaildate").
		Set("invoiceduedate = EXCLUDED.invoiceduedate").
		Set("invoiceperiod = EXCLUDED.invoiceperiod").
		Set("invoicepostdate = EXCLUDED.invoicepostdate").
		Set("billingdate = EXCLUDED.billingdate").
		Set("pdfnumber = EXCLUDED.pdfnumber").
		Set("companycountry = EXCLUDED.companycountry").
		Set("customercountry = EXCLUDED.customercountry").Insert()

	if err != nil {
		log.Warn("Can't add invoice: %v", err)
	}

	return err
}

func (p *ConnectionWrapper) GetLastMonthTaxRate(billingSettings, companyCountry, customerCountry string, billingTime time.Time) (float64, error) {
	var invoice []*Invoice

	layout := "2006-01-02"

	query := fmt.Sprintf(
		`SELECT  * from public.invoice WHERE 
					billingsetting  ='%s' AND 
					companycountry  ='%s' AND 
					customercountry ='%s' AND
					billingdate    <='%s' AND 
					billingdate    >='%s'
				ORDER BY billingdate DESC`,

		billingSettings,
		companyCountry,
		customerCountry,
		now.New(billingTime).BeginningOfMonth().Format(layout),
		now.New(now.New(billingTime).BeginningOfMonth().Add(-24*365*time.Hour)).BeginningOfMonth().Format(layout))

	log.Warnf("get last month query: %v", query)
	_, err := p.client.Query(&invoice, query)

	if err != nil {
		log.Warnf("can't execute pg query: %s", err)
		return 0, err
	}
	if len(invoice) > 0 {
		return invoice[0].TaxRate, nil
	}
	return 0, errors.New("not found record for previous month")
}

func (p *ConnectionWrapper) GetInvoices(billingDate string, approved bool) ([]*InvoiceUp, error) {
	str := time.Now()
	defer func() {
		log.Infof("GetApprovedInvoices query: %v seconds", time.Since(str).Seconds()*1000)
	}()
	var invoices []*InvoiceUp
	var notSuffix string

	if !approved {
		notSuffix = "NOT"
	}
	query := fmt.Sprintf(`SELECT i.pdfnumber,
		i.invoicenumber,
		i.billingsetting,
		i.invoiceamount,
		SUM(il.lineitemtaxamount) AS taxtotal,
		i.taxrate,
		i.accountnumber,
		i.sapcustomerid,
		i.invoicecurrency,
		i.billingdate,
		i.invoiceduedate,
		i.companycountry,
		i.isapproved,
		(SELECT po_number from charge WHERE charge_id IN (SELECT charge_id FROM invoicelineitem WHERE invoicenumber = i.invoicenumber) AND po_number NOTNULL LIMIT 1) AS po_number,
		comp."Name" as name
	FROM public.invoice i
       JOIN invoicelineitem il ON il.invoicenumber = i.invoicenumber
       JOIN sfdc."BillingSetting" b ON i.billingsetting = b."Id"
       JOIN sfdc."Company" comp ON b."Company__c" = comp."Id"
	WHERE %v i.isapproved 
	AND i.isUploadedToSalesforce = false
	AND i.billingdate='%v'
	GROUP BY i.invoicenumber, comp."Name"`,
		notSuffix,
		billingDate)

	_, err := p.client.Query(&invoices, query)

	if err != nil {
		log.Warnf("can't execute pg query: %s", err)
		return nil, err
	}

	return invoices, nil
}

func (p *ConnectionWrapper) MarkInvoiceAsPublished(invoices []*InvoiceUp) error {
	invoiceNumbers := make([]string, 0, len(invoices))
	for index := range invoices {
		invoiceNumbers = append(invoiceNumbers, strconv.Itoa(int(invoices[index].InvoiceNumber)))
	}
	query := fmt.Sprintf(`
		UPDATE public.invoice
		SET isUploadedToSalesforce=true
		WHERE invoicenumber in (%v)`, strings.Join(invoiceNumbers, ","))

	log.Infof(query)
	_, err := p.client.Exec(query)

	if err != nil {
		log.Warnf("can't execute pg query: %s", err)
		return err
	}
	return nil
}

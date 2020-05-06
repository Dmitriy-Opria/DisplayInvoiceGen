package postgres

import (
	"errors"
	"fmt"
	"github.com/jinzhu/now"
	"github.rakops.com/BNP/DisplayInvoiceGen/log"
	"time"
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
// NOT USED
func (p *ConnectionWrapper) CheckInvoiceExist(billingDate string) (bool, error) {
	var invoice []*Invoice
	query := fmt.Sprintf(`SELECT  * from public.invoice where billingdate='%s'`, billingDate)
	_, err := p.client.Query(&invoice, query)

	if err != nil {
		log.Warnf("can't execute pg query: %s", err)
		return false, err
	}
	if len(invoice) > 0 {
		return true, nil
	}
	return false, nil
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

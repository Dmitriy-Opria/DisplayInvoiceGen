package postgres

import (
	"fmt"
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
		Set("pdfnumber = EXCLUDED.pdfnumber").Insert()

	if err != nil {
		log.Warn("Can't add invoice: %v", err)
	}

	return err
}

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

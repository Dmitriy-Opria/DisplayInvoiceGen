package utils

import (
	"fmt"
	"math"
	"time"

	"github.com/jinzhu/now"
	"github.rakops.com/BNP/DisplayInvoiceGen/model"
	"github.rakops.com/BNP/DisplayInvoiceGen/postgres"
)

func GroupCharges(chargedList []*postgres.Charge) map[string][]*postgres.Charge {
	if len(chargedList) == 0 {
		return nil
	}
	chargedMap := map[string][]*postgres.Charge{}
	for _, charged := range chargedList {

		if chList, ok := chargedMap[charged.BillingSettings]; ok {
			chargedMap[charged.BillingSettings] = append(chList, charged)
			continue
		}
		chargedMap[charged.BillingSettings] = []*postgres.Charge{charged}
	}
	return chargedMap
}

func CombineInvoice(id int64, invoiceResponse *model.Response, tm time.Time, charges []*postgres.Charge) *postgres.Invoice {

	if len(charges) == 0 {
		return nil
	}

	layout := "2006-01-02"

	invoice := &postgres.Invoice{
		InvoiceNumber:   id,
		BillingSetting:  charges[0].BillingSettings,
		AccountNumber:   charges[0].Account,
		SAPCustomerID:   charges[0].SapCustomerID,
		PaymentTerms:    charges[0].PaymentsTermsSap,
		InvoiceCurrency: charges[0].ChangeCurrency,
		InsertDate:      time.Now(),
		InvoiceDueDate:  now.New(tm).EndOfMonth().Format(layout),
		InvoicePeriod:   now.New(now.New(tm).BeginningOfMonth().Add(-24 * time.Hour)).BeginningOfMonth().Format(layout),
		InvoicePostDate: time.Now().Format(layout),
		BillingDate:     now.New(tm).Format(layout),
		PDFnumber:       fmt.Sprintf("SIN%06d", id),
	}

	for index := range charges {
		invoice.InvoiceAmount += charges[index].ChargeAmount
	}

	invoice.InvoiceAmount = math.Round(invoice.InvoiceAmount*1000) / 1000

	tax := 0.0
	totalTax := 0.0
	for _, line := range invoiceResponse.InvoiceResponse.LineItem {
		totalTax += line.TotalTax
	}
	if invoiceResponse.InvoiceResponse.LineItem[0].TotalTax != 0 {
		tax = (totalTax / invoice.InvoiceAmount) * 100
		tax = math.Round(tax*1000) / 1000
	}

	invoice.TaxRate = tax

	return invoice
}

func CombineInvoiceLineItem(id int64, invoiceResponse *model.Response) []*postgres.InvoiceLineItem {

	if len(invoiceResponse.InvoiceResponse.LineItem) == 0 {
		return nil
	}
	invoiceLineItems := make([]*postgres.InvoiceLineItem, 0, len(invoiceResponse.InvoiceResponse.LineItem))

	for _, lineItem := range invoiceResponse.InvoiceResponse.LineItem {
		if len(lineItem.Taxes) == 0 {
			return nil
		}
		invoiceLineItem := &postgres.InvoiceLineItem{
			InvoiceNumber:  id,
			LineItemNumber: int64(lineItem.LineItemNumber),
		}

		invoiceLineItem.LineItemAmount = lineItem.ExtendedPrice
		invoiceLineItem.LineItemTaxAmount = lineItem.Taxes[0].CalculatedTax
		invoiceLineItem.LineItemTotalAmount = math.Round((invoiceLineItem.LineItemAmount+invoiceLineItem.LineItemTaxAmount)*100) / 100

		invoiceLineItems = append(invoiceLineItems, invoiceLineItem)
	}

	return invoiceLineItems
}

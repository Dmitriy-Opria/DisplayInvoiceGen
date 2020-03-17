package api

import (
	"fmt"
	"github.com/pkg/errors"
	"github.rakops.com/BNP/DisplayInvoiceGen/log"
	"math"
	"net/http"
	"time"

	"github.com/InVisionApp/rye"
	"github.com/jinzhu/now"
	"github.rakops.com/BNP/DisplayInvoiceGen/postgres"
)

func (a *Api) createInvoice(w http.ResponseWriter, r *http.Request) *rye.Response {
	ctx := r.Context()

	billingDate := ctx.Value(ContextBillingDate).(string)

	exist, err := a.Deps.Postgres.CheckInvoiceExist(billingDate)
	if err != nil || len(billingDate) == 0 {
		return &rye.Response{
			Err:        err,
			StatusCode: http.StatusInternalServerError,
		}
	}

	chargedList, err := a.Deps.Postgres.GetChargedList(billingDate)
	if err != nil {
		return &rye.Response{
			Err:        err,
			StatusCode: http.StatusInternalServerError,
		}
	}

	if !exist {

		groupedList := groupCharges(chargedList)
		taxRate := a.Deps.ExternalService.GetTaxRate()

		billingTime := ctx.Value(ContextBillingTime).(time.Time)

		for _, list := range groupedList {
			id, err := a.Deps.Postgres.GetInvoiceSequence()
			if err != nil {
				return &rye.Response{
					Err:        err,
					StatusCode: http.StatusInternalServerError,
				}
			}

			err = a.Deps.Postgres.AddInvoice(combineInvoice(id, taxRate, billingTime, list))
			if err != nil {
				return &rye.Response{
					Err:        errors.Wrap(err, "unable to insert invoice"),
					StatusCode: http.StatusInternalServerError,
				}
			}

			err = a.Deps.Postgres.AddInvoiceLineItem(combineInvoiceLineItem(id, taxRate, list))
			if err != nil {
				return &rye.Response{
					Err:        errors.Wrap(err, "unable to insert invoice line item"),
					StatusCode: http.StatusInternalServerError,
				}
			}

			log.Printf("inserted id: %v", id)
		}
	} else {
		return &rye.Response{
			Err:        errors.New("invoice already exist"),
			StatusCode: http.StatusBadRequest,
		}
	}
	return nil
}

func groupCharges(chargedList []*postgres.Charge) map[string][]*postgres.Charge {
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

func combineInvoice(id, taxRate int64, tm time.Time, charges []*postgres.Charge) *postgres.Invoice {

	if len(charges) == 0 {
		return nil
	}

	layout := "2006-01-02"

	invoice := &postgres.Invoice{
		InvoiceNumber:   id,
		BillingSetting:  charges[0].BillingSettings,
		TaxRate:         taxRate,
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

	return invoice
}

func combineInvoiceLineItem(id, taxRate int64, charges []*postgres.Charge) []*postgres.InvoiceLineItem {

	if len(charges) == 0 {
		return nil
	}
	invoiceLineItems := make([]*postgres.InvoiceLineItem, 0, len(charges))

	for index := range charges {
		invoiceLineItem := &postgres.InvoiceLineItem{
			InvoiceNumber:  id,
			LineItemNumber: int64(index + 1),
		}
		invoiceLineItem.LineItemAmount = math.Round(charges[index].ChargeAmount*100) / 100
		invoiceLineItem.LineItemTaxAmount = math.Round(charges[index].ChargeAmount*(float64(taxRate)*0.01)*100) / 100
		invoiceLineItem.LineItemTotalAmount = math.Round(invoiceLineItem.LineItemAmount+invoiceLineItem.LineItemTaxAmount*100) / 100

		invoiceLineItems = append(invoiceLineItems, invoiceLineItem)
	}

	return invoiceLineItems
}

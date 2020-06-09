package api

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/InVisionApp/rye"
	"github.rakops.com/BNP/DisplayInvoiceGen/utils"
)

func (a *Api) uploadInvoice(w http.ResponseWriter, r *http.Request) *rye.Response {
	ctx := r.Context()

	billingDate := ctx.Value(ContextBillingDate).(string)

	invoices, err := a.Deps.Postgres.GetInvoices(billingDate, true)
	if err != nil {
		return ServerErrorResponse(err, "can't select approved invoices")
	}
	
	if len(invoices) > 0 {
		invoiceLineItems, err := a.Deps.Postgres.GetInvoicesLineItems(billingDate, true)
		if err != nil {
			return ServerErrorResponse(err, "can't select approved invoices line items")
		}

		err = a.Deps.SalesForce.PushToSalesForce(utils.ConvertInvoice(invoices), utils.ConvertInvoiceLineItems(invoiceLineItems))
		if err != nil {
			return ServerErrorResponse(err, "can'push to salse force")
		}

		err = a.Deps.Postgres.MarkInvoiceAsPublished([]int64{})
		if err != nil {
			return ServerErrorResponse(err, "can't mark invoices as published")
		}

		jsonData, _ := json.Marshal(&map[string]string{
			"Message":          "Successfully uploaded",
			"Status":           "OK",
			"Invoices":         strconv.Itoa(len(invoices)),
			"InvoiceLineItems": strconv.Itoa(len(invoiceLineItems)),
		})

		rye.WriteJSONResponse(w, http.StatusOK, jsonData)
	}
	return nil
}

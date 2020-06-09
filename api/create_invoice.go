package api

import (
	"encoding/json"
	"net/http"
	"strconv"
	"time"

	"github.com/InVisionApp/rye"
	"github.com/pkg/errors"
	"github.rakops.com/BNP/DisplayInvoiceGen/log"
	"github.rakops.com/BNP/DisplayInvoiceGen/utils"
)

func (a *Api) createInvoice(w http.ResponseWriter, r *http.Request) *rye.Response {
	ctx := r.Context()

	billingDate := ctx.Value(ContextBillingDate).(string)

	chargedList, err := a.Deps.Postgres.GetNotProcessedChargedList(billingDate)
	if err != nil {
		return ServerErrorResponse(err, "can't get charged list")
	}

	if len(chargedList) > 0 {

		groupedList := utils.GroupCharges(chargedList)

		billingTime := ctx.Value(ContextBillingTime).(time.Time)

		for _, list := range groupedList {

			id, err := a.Deps.Postgres.GetInvoiceSequence()
			if err != nil {
				return ServerErrorResponse(err, "can't get invoice sequence")
			}

			taxResponse, err := a.Deps.ExternalService.GetTaxResponse(list)
			if err != nil {
				log.WithFields(log.Fields{
					"invoiceNumber":  id,
					"billingSetting": list[0].BillingSettings,
					"billingDate":    billingDate,
				}).Warnf("error_calling, can't get tax response: %v", err)
				return BadRequestResponse(err, "can't get tax service response")
			}

			switch err {
			case nil:
				err := a.Deps.Postgres.AddInvoice(utils.CombineInvoice(id, taxResponse, billingTime, list))
				if err != nil {
					return ServerErrorResponse(err, "unable to insert invoice")
				}

				err = a.Deps.Postgres.AddInvoiceLineItem(utils.CombineInvoiceLineItem(id, taxResponse))
				if err != nil {
					return ServerErrorResponse(err, "unable to insert invoice line item")
				}
			default:
				taxRate, err := a.Deps.Postgres.GetLastMonthTaxRate(list[0].BillingSettings, list[0].RakutenCountry, list[0].BillingCountryCode, billingTime)
				if err != nil {
					log.WithFields(log.Fields{
						"invoiceNumber":  id,
						"billingSetting": list[0].BillingSettings,
						"billingDate":    billingDate,
					}).Warnf("error during finding previous taxrate: %v", err)
					return NotFoundResponse(err, "can't found last month tax rate")
				}

				err = a.Deps.Postgres.AddInvoice(utils.CombineCalculatedInvoice(id, taxRate, billingTime, list))
				if err != nil {
					return ServerErrorResponse(err, "unable to insert invoice")
				}

				err = a.Deps.Postgres.AddInvoiceLineItem(utils.CombineCalculatedInvoiceLineItem(id, taxRate, list))
				if err != nil {
					return ServerErrorResponse(err, "unable to insert invoice line item")
				}
			}

			jsonData, _ := json.Marshal(&map[string]string{
				"Message": "Successfully created",
				"Status":  "OK",
				"ID":      strconv.Itoa(int(id)),
			})

			rye.WriteJSONResponse(w, http.StatusOK, jsonData)
			log.Printf("inserted id: %v", id)
		}
	} else {
		return BadRequestResponse(errors.New("invoice already exist"), "")
	}
	return nil
}

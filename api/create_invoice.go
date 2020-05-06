package api

import (
	"net/http"
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
		return &rye.Response{
			Err:        err,
			StatusCode: http.StatusInternalServerError,
		}
	}

	if len(chargedList) > 0 {

		groupedList := utils.GroupCharges(chargedList)

		billingTime := ctx.Value(ContextBillingTime).(time.Time)

		for _, list := range groupedList {

			id, err := a.Deps.Postgres.GetInvoiceSequence()
			if err != nil {
				return &rye.Response{
					Err:        err,
					StatusCode: http.StatusInternalServerError,
				}
			}

			taxResponse, err := a.Deps.ExternalService.GetTaxResponse(list)
			if err != nil {
				log.WithFields(log.Fields{
					"invoiceNumber":  id,
					"billingSetting": list[0].BillingSettings,
					"billingDate":    billingDate,
				}).Warnf("error_calling, can't get tax response: %v", err)
			}

			switch err {
			case nil:
				err := a.Deps.Postgres.AddInvoice(utils.CombineInvoice(id, taxResponse, billingTime, list))
				if err != nil {
					return &rye.Response{
						Err:        errors.Wrap(err, "unable to insert invoice"),
						StatusCode: http.StatusInternalServerError,
					}
				}

				err = a.Deps.Postgres.AddInvoiceLineItem(utils.CombineInvoiceLineItem(id, taxResponse))
				if err != nil {
					return &rye.Response{
						Err:        errors.Wrap(err, "unable to insert invoice line item"),
						StatusCode: http.StatusInternalServerError,
					}
				}
			default:
				taxRate, err := a.Deps.Postgres.GetLastMonthTaxRate(list[0].BillingSettings, list[0].RakutenCountry, list[0].BillingCountryCode, billingTime)
				if err != nil {
					log.WithFields(log.Fields{
						"invoiceNumber":  id,
						"billingSetting": list[0].BillingSettings,
						"billingDate":    billingDate,
					}).Warnf("error during finding previous taxrate: %v", err)
				}

				err = a.Deps.Postgres.AddInvoice(utils.CombineCalculatedInvoice(id, taxRate, billingTime, list))
				if err != nil {
					return &rye.Response{
						Err:        errors.Wrap(err, "unable to insert invoice"),
						StatusCode: http.StatusInternalServerError,
					}
				}

				err = a.Deps.Postgres.AddInvoiceLineItem(utils.CombineCalculatedInvoiceLineItem(id, taxRate, list))
				if err != nil {
					return &rye.Response{
						Err:        errors.Wrap(err, "unable to insert invoice line item"),
						StatusCode: http.StatusInternalServerError,
					}
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

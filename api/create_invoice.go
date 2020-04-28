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

		groupedList := utils.GroupCharges(chargedList)

		billingTime := ctx.Value(ContextBillingTime).(time.Time)

		for _, list := range groupedList {

			taxResponse, err := a.Deps.ExternalService.GetTaxResponse(list)
			if err != nil {
				return &rye.Response{
					Err:        err,
					StatusCode: http.StatusExpectationFailed,
				}
			}

			id, err := a.Deps.Postgres.GetInvoiceSequence()
			if err != nil {
				return &rye.Response{
					Err:        err,
					StatusCode: http.StatusInternalServerError,
				}
			}

			err = a.Deps.Postgres.AddInvoice(utils.CombineInvoice(id, taxResponse, billingTime, list))
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

package queue

import (
	"time"

	"github.com/pkg/errors"
	"github.rakops.com/BNP/DisplayInvoiceGen/log"
	"github.rakops.com/BNP/DisplayInvoiceGen/utils"
)

var (
	AlreadyExist = errors.New("invoice already exist")
)

func (q *Wrapper) CreateInvoice(billingDate string) error {

	chargedList, err := q.deps.Postgres.GetNotProcessedChargedList(billingDate)
	if err != nil {
		return errors.Wrap(err, "can't get charged list")
	}

	if len(chargedList) > 0 {

		groupedList := utils.GroupCharges(chargedList)

		billingTime, err := time.Parse("2006-01-02", billingDate)
		if err != nil {
			return errors.Wrap(err, "can't parse billing time date")
		}

		for _, list := range groupedList {

			id, err := q.deps.Postgres.GetInvoiceSequence()
			if err != nil {
				return errors.Wrap(err, "can't get invoice sequence")
			}
			taxResponse, err := q.deps.ExternalService.GetTaxResponse(list)
			if err != nil {
				log.WithFields(log.Fields{
					"invoiceNumber":  id,
					"billingSetting": list[0].BillingSettings,
					"billingDate":    billingDate,
				}).Warnf("error_calling, can't get tax response: %v", err)
			}

			switch err {
			case nil:
				err := q.deps.Postgres.AddInvoice(utils.CombineInvoice(id, taxResponse, billingTime, list))
				if err != nil {
					return errors.Wrap(err, "unable to insert invoice")
				}

				err = q.deps.Postgres.AddInvoiceLineItem(utils.CombineInvoiceLineItem(id, taxResponse))
				if err != nil {
					return errors.Wrap(err, "unable to insert invoice line item")
				}
			default:
				taxRate, err := q.deps.Postgres.GetLastMonthTaxRate(list[0].BillingSettings, list[0].RakutenCountry, list[0].BillingCountryCode, billingTime)
				if err != nil {
					log.WithFields(log.Fields{
						"invoiceNumber":  id,
						"billingSetting": list[0].BillingSettings,
						"billingDate":    billingDate,
					}).Warnf("error_finding during finding previous tax rate: %v", err)
					return errors.Wrap(err,"can't find previous tax rate")
				}

				err = q.deps.Postgres.AddInvoice(utils.CombineCalculatedInvoice(id, taxRate, billingTime, list))
				if err != nil {
					return errors.Wrap(err, "unable to insert invoice")
				}

				err = q.deps.Postgres.AddInvoiceLineItem(utils.CombineCalculatedInvoiceLineItem(id, taxRate, list))
				if err != nil {
					return errors.Wrap(err, "unable to insert invoice line item")
				}
			}
			log.Printf("inserted id: %v", id)
		}
	} else {
		return AlreadyExist
	}
	return nil
}

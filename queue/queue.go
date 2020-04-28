package queue

import (
	"time"

	"github.com/pkg/errors"
	"github.rakops.com/BNP/DisplayInvoiceGen/deps"
	"github.rakops.com/BNP/DisplayInvoiceGen/log"
	"github.rakops.com/BNP/DisplayInvoiceGen/utils"
)

type Handler interface {
	Run() error
	CreateInvoice(billingDate string) error
}

func NewQueueHandler(deps *deps.Dependencies) Handler {
	return &Wrapper{
		deps: deps,
	}
}

type Wrapper struct {
	deps *deps.Dependencies
}

func (q *Wrapper) Run() error {

	msg, err := q.deps.Consumer.GetChannel()

	if err != nil {
		return errors.Wrap(err, "can't get channel")
	}

	go func() {
		for d := range msg {
			log.Warnf(
				"got %dB delivery: [%v] %q",
				len(d.Body),
				d.DeliveryTag,
				d.Body,
			)

			if err := q.CreateInvoice(string(d.Body)); err != nil {
				log.Warnf("can't create invoice: %v", err)
				continue
			}

			d.Ack(false)
		}
		log.Warnf("handle: deliveries channel closed")
		q.deps.Consumer.Done(nil)
	}()

	return nil
}

func (q *Wrapper) CreateInvoice(billingDate string) error {

	exist, err := q.deps.Postgres.CheckInvoiceExist(billingDate)
	if err != nil || len(billingDate) == 0 {
		return errors.Wrap(err, "can't check invoice exist")
	}

	chargedList, err := q.deps.Postgres.GetChargedList(billingDate)
	if err != nil {
		return errors.Wrap(err, "can't get charged list")
	}

	if !exist {

		groupedList := utils.GroupCharges(chargedList)

		billingTime, err := time.Parse("2006-01-02", billingDate)
		if err != nil {
			return errors.Wrap(err, "can't parse billing time date")
		}

		for _, list := range groupedList {

			taxResponse, err := q.deps.ExternalService.GetTaxResponse(list)
			if err != nil {
				return errors.Wrap(err, "can't get tax response")
			}

			id, err := q.deps.Postgres.GetInvoiceSequence()
			if err != nil {
				return errors.Wrap(err, "can't get invoice sequence")
			}

			err = q.deps.Postgres.AddInvoice(utils.CombineInvoice(id, taxResponse, billingTime, list))
			if err != nil {
				return errors.Wrap(err, "unable to insert invoice")
			}

			err = q.deps.Postgres.AddInvoiceLineItem(utils.CombineInvoiceLineItem(id, taxResponse))
			if err != nil {
				return errors.Wrap(err, "unable to insert invoice line item")
			}

			log.Printf("inserted id: %v", id)
		}
	} else {
		return errors.New("invoice already exist")
	}
	return nil
}

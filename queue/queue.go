package queue

import (
	"sync"
	"time"

	"github.com/pkg/errors"
	"github.com/streadway/amqp"
	"github.rakops.com/BNP/DisplayInvoiceGen/deps"
	"github.rakops.com/BNP/DisplayInvoiceGen/log"
	"github.rakops.com/BNP/DisplayInvoiceGen/rabbit/consumer"
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
	deps                 *deps.Dependencies
	salesForceTimeouts   map[string]time.Time
	salesForceTimeoutsMx sync.Mutex
}

func (q *Wrapper) Run() error {

	invoiceChannel, err := q.deps.Consumer.GetChannel(consumer.Invoice)
	if err != nil {
		return errors.Wrap(err, "can't get channel")
	}

	salesForceChannel, err := q.deps.Consumer.GetChannel(consumer.SalesForce)
	if err != nil {
		return errors.Wrap(err, "can't get channel")
	}

	// Generate invoices worker
	go q.InvoiceProcessor(invoiceChannel)

	// Push approved invoice to salesforce worker
	go q.SalesForceProcessor(salesForceChannel)
	return nil
}

func (q *Wrapper) InvoiceProcessor(channel <-chan amqp.Delivery) {

Restart:
	for d := range channel {
		log.Warnf(
			"got invoice %dB delivery: [%v] %q",
			len(d.Body),
			d.DeliveryTag,
			d.Body,
		)

		if err := q.CreateInvoice(string(d.Body)); err == AlreadyExist {
			d.Ack(false)
			log.Warnf("ack existed invoice: %v", err)
			continue
		} else if err != nil {
			log.Warnf("can't create invoice: %v", err)
			d.Nack(false, false)
			continue
		}

		q.deps.Producer.Send(string(d.Body))
		d.Ack(false)
		log.Info("invoice ack")
	}
	log.Warnf("handle: deliveries channel closed")
	q.deps.Consumer.Done(nil)

	time.Sleep(10 * time.Minute)

	var err error
	channel, err = q.deps.Consumer.GetChannel(consumer.Invoice)
	if err != nil {
		log.Warnf("can't reconnect to rabbitMQ, will try after 10 minutes: %v", err)
	}

	goto Restart
}

func (q *Wrapper) SalesForceProcessor(channel <-chan amqp.Delivery) {

Restart:
	for d := range channel {
		log.Warnf(
			"got salesforce %dB delivery: [%v] %q",
			len(d.Body),
			d.DeliveryTag,
			d.Body,
		)

		func() {
			str := time.Now()
			defer func() {
				log.Infof("ticker time processing: %v seconds", time.Since(str).Seconds()*1000)
			}()
			// process msg for 5 hours, if it still not approved
			if q.TimedOut(string(d.Body)) {
				log.Infof("salesforce sending timed out, billing date: %v", string(d.Body))
				d.Ack(false)
				return
			}

			invoicesApproved, err := q.deps.Postgres.GetInvoices(string(d.Body), true)
			if err != nil {
				log.Warnf("can't select approved invoices: %v", err)
				d.Nack(false, false)
				return
			}
			if len(invoicesApproved) > 0 {
				invoiceLineItems, err := q.deps.Postgres.GetInvoicesLineItems(string(d.Body), true)
				if err != nil {
					log.Warnf("can't select approved invoices line items: %v", err)
					d.Nack(false, false)
					return
				}
				err = q.deps.SalesForce.PushToSalesForce(utils.ConvertInvoice(invoicesApproved, q.deps.Config.SalesForce.RecordTypeId), utils.ConvertInvoiceLineItems(invoiceLineItems))
				if err != nil {
					log.Warnf("can't publish invoices to salesforce: %v", err)
					d.Nack(false, false)
					return
				}
				err = q.deps.Postgres.MarkInvoiceAsPublished(invoicesApproved)
				if err != nil {
					log.Warnf("can't mark invoices as published: %v", err)
					d.Nack(false, false)
					return
				}
				log.Infof("processed invoices: %v", len(invoicesApproved))
				log.Infof("processed invoice line items: %v", len(invoiceLineItems))
				log.Info("salesforce ack")
				d.Ack(false)
			} else {
				log.Info("retry salesforce processing")
				time.Sleep(time.Duration(q.deps.Config.SalesForce.TickerPeriod) * time.Second)
				d.Reject(true)
			}
		}()
	}
	log.Warnf("handle: deliveries channel closed")
	q.deps.Consumer.Done(nil)

	time.Sleep(10 * time.Minute)

	var err error
	channel, err = q.deps.Consumer.GetChannel(consumer.SalesForce)
	if err != nil {
		log.Warnf("can't reconnect to rabbitMQ, will try after 10 minutes: %v", err)
	}

	goto Restart
}

func (q *Wrapper) TimedOut(billingDate string) bool {
	q.salesForceTimeoutsMx.Lock()
	defer q.salesForceTimeoutsMx.Unlock()

	if q.salesForceTimeouts == nil {
		q.salesForceTimeouts = map[string]time.Time{}
	}

	if tm, ok := q.salesForceTimeouts[billingDate]; ok {
		// if trying to process msg is over 5 hours
		if time.Now().After(tm.Add(time.Duration(q.deps.Config.SalesForce.ValidationPeriod) * time.Minute)) {
			return true
		}
	} else {
		q.salesForceTimeouts[billingDate] = time.Now()
		return false
	}
	return false
}

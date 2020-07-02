package postgres

import (
	"time"

	"github.com/go-pg/pg"
	"github.rakops.com/BNP/DisplayInvoiceGen/log"
)

type IConnection interface {
	Connect() error
	Close()
	IsAlive() bool

	AddInvoice(invoice *Invoice) error
	GetInvoices(billingDate string, approved bool) ([]*InvoiceUp, error)
	MarkInvoiceAsPublished(invoices []*InvoiceUp) error

	AddInvoiceLineItem(invoiceLineItem []*InvoiceLineItem) error
	GetInvoicesLineItems(billingDate string, approved bool) ([]*InvoiceLineItemUP, error)

	GetInvoiceSequence() (int64, error)
	GetLastMonthTaxRate(billingSettings, companyCountry, customerCountry string, billingTime time.Time) (float64, error)

	GetNotProcessedChargedList(billingDate string) ([]*Charge, error)
}

type IConf interface {
	GetAddr() string
	GetUser() string
	GetPass() string
	GetDB() string
	GetDebug() bool
}

// ConnectionWrapper wrapper on Postgres connection
type ConnectionWrapper struct {
	client  *pg.DB
	options *pg.Options
	debug   bool
}

// NewConnection establish postgres connection
func NewConnection(conf IConf) IConnection {
	conn := new(ConnectionWrapper)

	conn.options = &pg.Options{
		Addr:         conf.GetAddr(),
		User:         conf.GetUser(),
		Password:     conf.GetPass(),
		Database:     conf.GetDB(),
		MaxRetries:   100,
		DialTimeout:  time.Minute * 5,
		ReadTimeout:  time.Minute * 5,
		MinIdleConns: 10,
		WriteTimeout: time.Minute * 5,
	}
	conn.debug = conf.GetDebug()
	return conn
}

func (p *ConnectionWrapper) Connect() error {
	p.client = pg.Connect(p.options)
	return nil
}

func (p *ConnectionWrapper) IsAlive() bool {
	var n int
	_, err := p.client.QueryOne(pg.Scan(&n), "SELECT 1")
	if err == nil {
		return true
	}
	log.Fatal("IConnection to Postgres lost. Trying to reconnect", err)

	p.client.Close()

	for {
		if err = p.Connect(); err == nil {
			return true
		}

		log.Warn("Failed connect to Postgres. Trying to reconnect", err)
		time.Sleep(time.Second)
	}
}

func (p *ConnectionWrapper) Close() {
	err := p.client.Close()
	if err != nil {
		log.Warn(err)
	}
}

package postgres

import (
	"fmt"
	"time"

	"github.com/go-pg/pg"
	"github.rakops.com/BNP/DisplayInvoiceGen/log"
)

type InvoiceLineItem struct {
	tableName           struct{} `sql:"InvoiceLineItem"`
	InvoiceNumber       int64    `json:"InvoiceNumber"                      sql:"invoicenumber"`
	LineItemNumber      int64    `json:"LineItemNumber"                     sql:"lineitemnumber"`
	LineItemAmount      float64  `json:"LineItemAmount"                     sql:"lineitemamount,notnull,type:numeric"`
	LineItemTaxAmount   float64  `json:"LineItemTaxAmount"                  sql:"lineitemtaxamount,notnull,type:numeric"`
	LineItemTotalAmount float64  `json:"LineItemTotalAmount"                sql:"lineitemtotalamount,notnull,type:numeric"`
	ChargeID            int64    `json:"ChargeID"                           sql:"charge_id,notnull,type:bigint"`
}

type InvoiceLineItemUP struct {
	InvoiceLineItem
	InvoiceID       string `json:"InvoiceID"                          sql:"-"`
	InvoiceCurrency string `json:"InvoiceCurrency"                    sql:"invoicecurrency"`
	ChargeID        int64  `json:"ChargeID"                           sql:"charge_id,notnull,type:bigint"`
	Description     string `json:"Description"                        sql:"note"`
}

func (p *ConnectionWrapper) AddInvoiceLineItem(invoiceLineItems []*InvoiceLineItem) error {

	var err error
	var tx *pg.Tx
	defer func(tx *pg.Tx, err error) {
		if err != nil {
			err = tx.Rollback()
			if err != nil {
				log.Errorln(err)
			}
		}
	}(tx, err)

	tx, err = p.client.Begin()
	if err != nil {
		return err
	}
	for _, invoiceLineItem := range invoiceLineItems {
		_, err := p.client.Model(invoiceLineItem).
			Set("invoicenumber = EXCLUDED.invoicenumber").
			Set("lineitemnumber = EXCLUDED.lineitemnumber").
			Set("lineitemamount = EXCLUDED.lineitemamount").
			Set("lineitemtaxamount = EXCLUDED.lineitemtaxamount").
			Set("lineitemtotalamount = EXCLUDED.lineitemtotalamount").
			Set("charge_id = EXCLUDED.charge_id").Insert()

		if err != nil {
			log.Errorln(err)
			return err
		}
	}
	err = tx.Commit()
	if err != nil {
		log.Errorln(err)
		return err
	}
	return nil
}

func (p *ConnectionWrapper) GetInvoicesLineItems(billingDate string, approved bool) ([]*InvoiceLineItemUP, error) {
	str := time.Now()
	defer func() {
		log.Infof("GetInvoicesLineItems query: %v seconds", time.Since(str).Seconds()*1000)
	}()
	var invoiceLineItems []*InvoiceLineItemUP
	var notSuffix string

	if !approved {
		notSuffix = "NOT"
	}
	query := fmt.Sprintf(`SELECT 
		il.invoicenumber,
		il.lineitemtotalamount,
		il.lineitemtaxamount,
		il.lineitemamount,
		il.lineitemnumber,
		invoicecurrency,
		c.charge_id,
		c.note
	FROM invoice i
		JOIN invoicelineitem il on il.invoicenumber = i.invoicenumber
		JOIN charge c on il.charge_id = c.charge_id
	WHERE %v i.isapproved and i.billingdate='%v'`,
		notSuffix,
		billingDate)

	_, err := p.client.Query(&invoiceLineItems, query)

	if err != nil {
		log.Warnf("can't execute pg query: %s", err)
		return nil, err
	}

	return invoiceLineItems, nil
}

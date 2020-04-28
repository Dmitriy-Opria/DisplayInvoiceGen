package postgres

import (
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
	ChargeID 			int64  	 `json:"ChargeID"                sql:"charge_id,notnull,type:bigint"`
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

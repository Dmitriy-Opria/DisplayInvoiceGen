package postgres

import (
	"fmt"
	"github.com/go-pg/pg"
	"github.rakops.com/BNP/DisplayInvoiceGen/log"
	"time"
)

type Charge struct {
	tableName                 struct{} `sql:"charge"`
	BillingSettings           string   `json:"BillingSettings"           sql:"GBS_Billing_Setting__c"`
	BillingName               string   `json:"BillingName"               sql:"Billing_Name__c"`
	Account                   string   `json:"Account"                   sql:"Account__c"`
	ChargeAmount              float64  `json:"ChargeAmount"              sql:"charge_amount"`
	ChangeCurrency            string   `json:"ChangeCurrency"            sql:"charge_currency"`
	SapCustomerID             string   `json:"SapCustomerID"             sql:"sap_customer_id__c"`
	PaymentsTermsSap          string   `json:"PaymentsTermsSap"          sql:"payment_terms_sap__c"`
	CODAVATRegistrationNumber string   `json:"CODAVATRegistrationNumber" sql:"c2g__CODAVATRegistrationNumber__c"`
	BillingCountryCode        string   `json:"BillingCountryCode"        sql:"BillingCountryCode"`
	VATRegistrationNumber     string   `json:"VATRegistrationNumber"     sql:"c2g__VATRegistrationNumber__c"`
	RakutenCountry            string   `json:"RakutenCountry"            sql:"c2g__Country__c"`
	ChargeId				  int64	   `json:"ChargeId" 				 sql:"charge_id"`
}

func (p *ConnectionWrapper) GetChargedList(billingDate string) ([]*Charge, error) {
	str := time.Now()
	defer func() {
		log.Infof("postgres query: %v seconds", time.Since(str).Seconds()*1000)
	}()
	var chargers []*Charge
	query := fmt.Sprintf(`SELECT  p."GBS_Billing_Setting__c",	
				b."Billing_Name__c",
				b."Account__c",
				c.charge_amount,
				c.charge_currency,
				acct.sap_customer_id__c,
				acct.payment_terms_sap__c,
				--Customer VAT and RakutenCountry Code
				a."c2g__CODAVATRegistrationNumber__c", 
				a."BillingCountryCode",
				--Seller VAT and RakutenCountry Code
				comp."c2g__VATRegistrationNumber__c",
				comp."c2g__Country__c",
				c.charge_id
			from public.charge c
			join sfdc."Program" p on c.program_id = p."Id" and billing_date = '%s'
			join sfdc."BillingSetting" b on p."GBS_Billing_Setting__c" = b."Id"
			join sfdc."Company" comp on b."Company__c" = comp."Id"
			join public.Account acct on b."Account__c" = acct.Id	
			join sfdc."Account" a on b."Account__c" = a."Id"
			order by p."GBS_Billing_Setting__c"`, billingDate)

	_, err := p.client.Query(&chargers, query)

	if err != nil {
		log.Warnf("can't execute pg query: %s", err)
		return nil, err
	}

	return chargers, nil
}

func (p *ConnectionWrapper) GetInvoiceSequence() (int64, error) {
	var id int64
	_, err := p.client.QueryOne(pg.Scan(&id), "SELECT nextval('Invoicesequence')")
	if err != nil {
		log.Warn("Can't select invoice sequence: %v", err)
		return 0, err
	}
	return id, nil
}

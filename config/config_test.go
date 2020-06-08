package config

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var testBody = []byte(`#CALCULATION APP
VERSION: 0.0.1
API_LISTEN_ADDRESS: 0.0.0.0:9065
PRODUCTION: false

#POSTGRES DB
PG_ADDR: 127.0.0.22:4000  
PG_USER: gbs_flow
PG_PASS: pass
PG_NAME: gbs
PG_DEBUG: false

#PABBIT_MQ QUEUE
RABBIT_HOST: 127.0.0.43
RABBIT_PORT: 5672
RABBIT_USER: admin
RABBIT_PASS: pass
RABBIT_VHOST: Vhost
RABBIT_EXCHANGE_NAME: exchange.display.invoice

RABBIT_CONSUMER_INVOICE_QUEUE_NAME: queue.display.invoice
RABBIT_CONSUMER_INVOICE_ROUTE_KEY: queue.display.invoice.routeKey

RABBIT_CONSUMER_SF_QUEUE_NAME: queue.display.data.salesforce.upload
RABBIT_CONSUMER_SF_ROUTE_KEY: queue.display.data.salesforce.upload.routeKey

RABBIT_PRODUCER_PDF_QUEUE_NAME: queue.display.pdf
RABBIT_PRODUCER_PDF_ROUTE_KEY: queue.display.pdf.routeKey

RABBIT_PRODUCER_SF_QUEUE_NAME: queue.display.data.salesforce.upload
RABBIT_PRODUCER_SF_ROUTE_KEY: queue.display.data.salesforce.upload.routeKey


#EXTERNAL SERVICE
API_ADDRESS: https://sloth-qa.private.linksynergy.com
API_AUTHORIZATION_TOKEN: 'Bearer'
ARI_RETRY: 3

#REQUEST PARAMS
COMPANY_NAME: 'Rakuten'
US_REGISTRATION:
  - 'USA'
AU_REGISTRATION:
  - 'AU'
EU_REGISTRATION:
  - 'UK'
TAX_ID_US: '111'
TAX_ID_AU: '222'
TAX_ID_EU: '333'

US_ISO_REGISTRATION: 'US'
AU_ISO_REGISTRATION: 'AU'
EU_ISO_REGISTRATION: 'GB'

PRODUCT_CLASS: 'ASDE'
TRANSACTION_TYPE: 'SALE'

#EXTERNAL SERVICE
SALES_FORCE_VERSION: v48.0
SALES_FORCE_CLIENT_ID: client_id
SALES_FORCE_CLIENT_SECRET: client_secret
SALES_FORCE_USER_NAME: test@gmail.com
SALES_FORCE_PASSWORD: client_pass
SALES_FORCE_SECURITY_TOKEN: client_token
TICKER_PERIOD_SEC: 30
VALIDATION_PERIOD_MIN: 300

RECORD_TYPE_ID: '0122f0000008t9iAAA'
`)

var _ = Describe("Invoice config test", func() {
	Describe("Config", func() {
		It("Invoice Unmarshal", func() {
			conf := initConfig(testBody)
			Expect(conf.Version).To(Equal("0.0.1"))
			Expect(conf.ListenAddress).To(Equal("0.0.0.0:9065"))
			Expect(conf.Production).To(Equal(false))

			Expect(conf.Postgres.Addr).To(Equal("127.0.0.22:4000"))
			Expect(conf.Postgres.User).To(Equal("gbs_flow"))
			Expect(conf.Postgres.Pass).To(Equal("pass"))
			Expect(conf.Postgres.Database).To(Equal("gbs"))

			Expect(conf.Rabbit.Host).To(Equal("127.0.0.43"))
			Expect(conf.Rabbit.Port).To(Equal(5672))
			Expect(conf.Rabbit.User).To(Equal("admin"))
			Expect(conf.Rabbit.Pass).To(Equal("pass"))
			Expect(conf.Rabbit.VHost).To(Equal("Vhost"))

			Expect(conf.Rabbit.ExchangeName).To(Equal("exchange.display.invoice"))

			Expect(conf.Rabbit.ConsumerInvoiceQueueName).To(Equal("queue.display.invoice"))
			Expect(conf.Rabbit.ConsumerInvoiceRouteKey).To(Equal("queue.display.invoice.routeKey"))

			Expect(conf.Rabbit.ConsumerSFQueueName).To(Equal("queue.display.data.salesforce.upload"))
			Expect(conf.Rabbit.ConsumerSFRouteKey).To(Equal("queue.display.data.salesforce.upload.routeKey"))

			Expect(conf.Rabbit.ProducerPDFQueueName).To(Equal("queue.display.pdf"))
			Expect(conf.Rabbit.ProducerPDFRouteKey).To(Equal("queue.display.pdf.routeKey"))

			Expect(conf.Rabbit.ProducerSFQueueName).To(Equal("queue.display.data.salesforce.upload"))
			Expect(conf.Rabbit.ProducerSFRouteKey).To(Equal("queue.display.data.salesforce.upload.routeKey"))

			Expect(conf.TaxCalculationService.AuthToken).To(Equal("Bearer"))
			Expect(conf.TaxCalculationService.Address).To(Equal("https://sloth-qa.private.linksynergy.com"))
			Expect(conf.TaxCalculationService.Retry).To(Equal(3))

			Expect(conf.TaxCalculationParams.CompanyName).To(Equal("Rakuten"))
			Expect(conf.TaxCalculationParams.TransactionType).To(Equal("SALE"))
			Expect(conf.TaxCalculationParams.ProductClass).To(Equal("ASDE"))
			Expect(conf.TaxCalculationParams.TaxIdUS).To(Equal("111"))
			Expect(conf.TaxCalculationParams.TaxIdAU).To(Equal("222"))
			Expect(conf.TaxCalculationParams.TaxIdEU).To(Equal("333"))

			Expect(conf.TaxCalculationParams.RegistrationAU).To(Equal(map[string]struct{}{"AU": {}}))
			Expect(conf.TaxCalculationParams.RegistrationUS).To(Equal(map[string]struct{}{"USA": {}}))
			Expect(conf.TaxCalculationParams.RegistrationEU).To(Equal(map[string]struct{}{"UK": {}}))

			Expect(conf.TaxCalculationParams.RegistrationIsoAU).To(Equal("AU"))
			Expect(conf.TaxCalculationParams.RegistrationIsoUS).To(Equal("US"))
			Expect(conf.TaxCalculationParams.RegistrationIsoEU).To(Equal("GB"))

			Expect(conf.SalesForce.ApiVersion).To(Equal("v48.0"))
			Expect(conf.SalesForce.ClientID).To(Equal("client_id"))
			Expect(conf.SalesForce.ClientSecret).To(Equal("client_secret"))
			Expect(conf.SalesForce.UserName).To(Equal("test@gmail.com"))
			Expect(conf.SalesForce.Pass).To(Equal("client_pass"))
			Expect(conf.SalesForce.SecurityToken).To(Equal("client_token"))
			Expect(conf.SalesForce.RecordTypeId).To(Equal("0122f0000008t9iAAA"))
			Expect(conf.SalesForce.TickerPeriod).To(Equal(int64(30)))
			Expect(conf.SalesForce.ValidationPeriod).To(Equal(int64(300)))
		})
	})
})

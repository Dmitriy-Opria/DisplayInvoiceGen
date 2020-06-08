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

RABBIT_CONSUMER_EXCHANGE_NAME: exchange.display.invoice
RABBIT_CONSUMER_QUEUE_NAME: queue.display.invoice.init
RABBIT_CONSUMER_ROUTE_KEY: queue.display.invoice.routeKey

RABBIT_PRODUCER_QUEUE_NAME: queue.display.pdf
RABBIT_PRODUCER_ROUTE_KEY: queue.display.pdf.routeKey

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

			Expect(conf.Rabbit.ConsumerExchangeName).To(Equal("exchange.display.invoice"))
			Expect(conf.Rabbit.ConsumerQueueName).To(Equal("queue.display.invoice.init"))
			Expect(conf.Rabbit.ConsumerRouteKey).To(Equal("queue.display.invoice.routeKey"))

			Expect(conf.Rabbit.ProducerQueueName).To(Equal("queue.display.pdf"))
			Expect(conf.Rabbit.ProducerRouteKey).To(Equal("queue.display.pdf.routeKey"))

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
		})
	})
})

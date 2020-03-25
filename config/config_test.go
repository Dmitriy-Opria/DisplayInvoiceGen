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

#EXTERNAL SERVICE
API_ADDRESS: https://sloth-qa.private.linksynergy.com
API_AUTHORIZATION_TOKEN: 'Bearer'

#REQUEST PARAMS
COMPANY_NAME: 'Rakuten'
DIVISION_US:
  - 'USA'
DIVISION_AU:
  - 'AU'
DIVISION_EU:
  - 'UK'
TAX_ID_US: '111'
TAX_ID_AU: '222'
TAX_ID_EU: '333'
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
			Expect(conf.Postgres.Debug).To(Equal(false))

			Expect(conf.TaxCalculationService.AuthToken).To(Equal("Bearer"))
			Expect(conf.TaxCalculationService.Address).To(Equal("https://sloth-qa.private.linksynergy.com"))

			Expect(conf.TaxCalculationParams.CompanyName).To(Equal("Rakuten"))
			Expect(conf.TaxCalculationParams.TransactionType).To(Equal("SALE"))
			Expect(conf.TaxCalculationParams.ProductClass).To(Equal("ASDE"))
			Expect(conf.TaxCalculationParams.TaxIdUS).To(Equal("111"))
			Expect(conf.TaxCalculationParams.TaxIdAU).To(Equal("222"))
			Expect(conf.TaxCalculationParams.TaxIdEU).To(Equal("333"))
			Expect(conf.TaxCalculationParams.DivisionAU).To(Equal(map[string]struct{}{"AU": {}}))
			Expect(conf.TaxCalculationParams.DivisionUS).To(Equal(map[string]struct{}{"USA": {}}))
			Expect(conf.TaxCalculationParams.DivisionEU).To(Equal(map[string]struct{}{"UK": {}}))
		})
	})
})

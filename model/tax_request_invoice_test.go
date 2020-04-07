package model

import (
	"encoding/json"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var invoiceRequest = `{
  "invoiceRequest": {
    "customer": {
      "destination": {
        "streetAddress1": "",
        "streetAddress2": "",
        "city": "",
        "mainDivision": "",
        "subDivision": "",
        "postalCode": "",
        "country": "GB"
      },
      "taxRegistration": [
        {
          "taxRegistrationNumber": "",
          "hasPhysicalPresenceIndicator": "true",
          "isoCountryCode": "GB"
        }
      ]
    },
    "seller": {
      "company": "Rakuten",
      "division": "RM_EU",
      "physicalOrigin": {
        "streetAddress1": "",
        "streetAddress2": "",
        "city": "",
        "mainDivision": "",
        "subDivision": "",
        "postalCode": "",
        "country": "GB"
      },
      "taxRegistration": [
        {
          "taxRegistrationNumber": "GB123456789",
          "hasPhysicalPresenceIndicator": "true",
          "isoCountryCode": "GB"
        }
      ]
    },
    "lineItem": [
      {
        "product": {
          "productClass": "ASDE"
        },
        "extendedPrice": "100",
        "lineItemNumber": "1"
      }
    ],
    "documentDate": "2020-02-27",
    "transactionType": "SALE"
  }
}`
var _ = Describe("Request Invoice API", func() {
	Describe("Invoice", func() {
		It("Invoice Unmarshal", func() {
			request := &Invoice{}
			err := json.Unmarshal([]byte(invoiceRequest), request)
			Expect(err).Should(BeNil())
		})

		It("Invoice Marshal", func() {
			request := Invoice{
				InvoiceRequest: InvoiceRequest{
					Customer: Customer{
						Destination: Destination{
							Country: "GB",
						},
						Tax: []TaxRegistrationInvoice{
							{
								TaxRegistrationNumber:        "",
								HasPhysicalPresenceIndicator: "true",
								IsoCountryCode:               "GB",
							},
						},
					},
					Seller: Seller{
						Company:  "Rakuten",
						Division: "RM_EU",
						PhysicalOrigin: Destination{
							Country: "GB",
						},
						TaxRegistration: []TaxRegistrationInvoice{
							{
								TaxRegistrationNumber:        "GB123456789",
								HasPhysicalPresenceIndicator: "true",
								IsoCountryCode:               "GB",
							},
						},
					},
					Lines: []LineItemInvoice{
						{
							Product: ProductInvoice{
								ProductClass: "ASDE",
							},
							ExtendedPrice:  "100",
							LineItemNumber: "1",
						},
					},
					DocumentDate:    "2020-02-27",
					TransactionType: "SALE",
				},
			}
			body, err := json.Marshal(request)

			Expect(err).Should(BeNil())
			Expect(string(body)).To(MatchJSON(invoiceRequest))
		})
	})
})

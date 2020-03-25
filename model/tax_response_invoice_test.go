package model

import (
	"encoding/json"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var invoiceResponse = `{
    "login": {
        "userName": null,
        "password": null,
        "trustedId": null
    },
    "accrualRequest": null,
    "accrualResponse": null,
    "accrualSyncRequest": null,
    "accrualSyncResponse": null,
    "assetMovementRequest": null,
    "assetMovementResponse": null,
    "buyerInputTaxRequest": null,
    "buyerInputTaxResponse": null,
    "deleteRequest": null,
    "deleteResponse": null,
    "distributeTaxProcurementRequest": null,
    "distributeTaxProcurementResponse": null,
    "distributeTaxRequest": null,
    "distributeTaxResponse": null,
    "inventoryRemovalRequest": null,
    "inventoryRemovalResponse": null,
    "invoiceRequest": null,
    "invoiceResponse": {
        "currency": null,
        "originalCurrency": null,
        "companyCodeCurrency": null,
        "seller": {
            "company": "Rakuten",
            "division": "RM_EU",
            "department": null,
            "utilityProvider": null,
            "dispatcher": null,
            "physicalOrigin": {
                "streetAddress1": null,
                "streetAddress2": null,
                "city": null,
                "mainDivision": null,
                "subDivision": null,
                "postalCode": null,
                "country": "GB",
                "currencyConversion": null,
                "taxAreaId": 808260000,
                "latitude": null,
                "longitude": null,
                "locationCustomsStatus": null,
                "locationCode": null,
                "externalJurisdictionCode": null
            },
            "administrativeOrigin": null,
            "taxRegistration": [],
            "nexusIndicator": null,
            "nexusReasonCode": null
        },
        "customer": {
            "customerCode": {
                "value": "",
                "classCode": null,
                "isBusinessIndicator": true
            },
            "destination": {
                "streetAddress1": null,
                "streetAddress2": null,
                "city": null,
                "mainDivision": null,
                "subDivision": null,
                "postalCode": null,
                "country": "GB",
                "currencyConversion": null,
                "taxAreaId": 808260000,
                "latitude": null,
                "longitude": null,
                "locationCustomsStatus": null,
                "locationCode": null,
                "externalJurisdictionCode": null
            },
            "administrativeDestination": null,
            "exemptionCertificate": null,
            "taxRegistration": [
                {
                    "taxRegistrationNumber": "GB",
                    "nexusOverride": [],
                    "physicalLocation": [],
                    "impositionType": null,
                    "isoCountryCode": "GB",
                    "mainDivision": null,
                    "hasPhysicalPresenceIndicator": true,
                    "jurisdictionId": null
                }
            ],
            "isTaxExempt": false,
            "exemptionReasonCode": null
        },
        "taxOverride": null,
        "impositionToProcess": [],
        "jurisdictionOverride": [],
        "situsOverride": null,
        "discount": null,
        "proratePercentage": null,
        "subTotal": 100,
        "total": 120,
        "totalTax": 20,
        "currencyConversionFactors": null,
        "lineItem": [
            {
                "seller": null,
                "customer": null,
                "taxOverride": null,
                "impositionToProcess": [],
                "jurisdictionOverride": [],
                "situsOverride": null,
                "product": {
                    "value": "",
                    "productClass": "ASDE"
                },
                "lineType": null,
                "commodityCode": null,
                "quantity": {
                    "value": 1,
                    "unitOfMeasure": null
                },
                "weight": null,
                "volume": null,
                "supplementaryUnit": null,
                "statisticalValue": null,
                "freight": null,
                "fairMarketValue": 100,
                "unitPrice": 0,
                "extendedPrice": 100,
                "discount": null,
                "amountBilledToDate": null,
                "companyCodeCurrencyTaxableAmount": null,
                "companyCodeCurrencyTaxAmount": null,
                "taxes": [
                    {
                        "jurisdiction": {
                            "value": "UNITED KINGDOM",
                            "jurisdictionLevel": "COUNTRY",
                            "jurisdictionId": 78295,
                            "effectiveDate": null,
                            "expirationDate": null,
                            "externalJurisdictionCode": null
                        },
                        "accumulateAsJurisdiction": null,
                        "calculatedTax": 20,
                        "effectiveRate": 0.2,
                        "taxApportionmentRate": null,
                        "basisReductionPercentage": null,
                        "exempt": null,
                        "nonTaxable": null,
                        "taxable": {
                            "value": 100,
                            "unitOfMeasure": null
                        },
                        "reportingBasis": null,
                        "imposition": {
                            "value": "VAT",
                            "userDefined": null,
                            "impositionId": 1
                        },
                        "impositionType": {
                            "value": "VAT",
                            "userDefined": null,
                            "impositionTypeId": 19,
                            "withholdingType": null
                        },
                        "accumulateAsImposition": null,
                        "accumulateAsImpositionType": null,
                        "taxRuleId": {
                            "value": 339387,
                            "userDefined": null,
                            "salesTaxHolidayIndicator": null,
                            "taxRuleType": null
                        },
                        "basisRuleId": null,
                        "inclusionRuleId": null,
                        "maxTaxRuleId": null,
                        "recoverableRuleId": null,
                        "postCalculationEvaluationRuleId": null,
                        "creditRuleId": null,
                        "basisApportionmentRuleId": null,
                        "taxRateAdjustmentRuleId": null,
                        "taxApportionmentRuleId": null,
                        "accumulationRuleId": null,
                        "telecomUnitConversionRuleId": null,
                        "reportingBasisRuleId": null,
                        "certificateNumber": null,
                        "recoverableAmount": null,
                        "recoverablePercent": null,
                        "blockingRecoverablePercent": null,
                        "partialExemptRecoverablePercent": null,
                        "unrecoverableAmount": null,
                        "filingCurrencyAmounts": null,
                        "sellerRegistrationId": "GB 882.7552.85",
                        "buyerRegistrationId": "GB",
                        "ownerRegistrationId": null,
                        "dispatcherRegistrationId": null,
                        "recipientRegistrationId": null,
                        "invoiceTextCode": [
                            21
                        ],
                        "summaryInvoiceText": null,
                        "invoiceTexts": null,
                        "assistedParameters": null,
                        "taxRuleCurrencyConversionFactors": null,
                        "originalTax": null,
                        "includedTax": null,
                        "nominalRate": 0.2,
                        "markUpRate": null,
                        "accumulateAsLineType": null,
                        "taxResult": "TAXABLE",
                        "taxType": "VAT",
                        "maxTaxIndicator": null,
                        "situs": "DESTINATION",
                        "notRegisteredIndicator": null,
                        "inputOutputType": "OUTPUT",
                        "taxCode": null,
                        "vertexTaxCode": null,
                        "reasonCode": null,
                        "filingCategoryCode": null,
                        "isService": null,
                        "rateClassification": "Standard Rate",
                        "taxCollectedFromParty": "BUYER",
                        "taxStructure": "SINGLE_RATE"
                    }
                ],
                "totalTax": 20,
                "flexibleFields": null,
                "returnsFields": null,
                "assistedParameters": {
                    "assistedParameter": [
                        {
                            "value": "GB",
                            "paramName": "customer.taxRegistrationId",
                            "phase": "PRE",
                            "ruleName": "Set Customer Registration ID",
                            "originalValue": ""
                        }
                    ]
                },
                "lineItem": [],
                "lineItemNumber": 1,
                "taxDate": null,
                "isMulticomponent": null,
                "locationCode": null,
                "deliveryTerm": null,
                "postingDate": null,
                "costCenter": null,
                "departmentCode": null,
                "generalLedgerAccount": null,
                "materialCode": null,
                "projectNumber": null,
                "usage": null,
                "usageClass": null,
                "vendorSKU": null,
                "countryOfOriginISOCode": null,
                "modeOfTransport": null,
                "natureOfTransaction": null,
                "intrastatCommodityCode": null,
                "netMassKilograms": null,
                "lineItemId": null,
                "taxIncludedIndicator": null,
                "transactionType": null,
                "simplificationCode": null,
                "titleTransfer": null,
                "chainTransactionPhase": null,
                "exportProcedure": null,
                "materialOrigin": null
            }
        ],
        "documentNumber": null,
        "accumulationDocumentNumber": null,
        "accumulationCustomerNumber": null,
        "documentType": null,
        "billingType": null,
        "orderType": null,
        "postingDate": null,
        "locationCode": null,
        "returnAssistedParametersIndicator": null,
        "returnGeneratedLineItemsIndicator": null,
        "deliveryTerm": null,
        "documentDate": 1582761600000,
        "transactionId": null,
        "transactionType": "SALE",
        "simplificationCode": null,
        "roundAtLineLevel": null,
        "paymentDate": null,
        "documentSequenceId": null,
        "taxPointDate": null
    },
    "invoiceVerificationRequest": null,
    "invoiceVerificationResponse": null,
    "purchaseOrderRequest": null,
    "purchaseOrderResponse": null,
    "quotationRequest": null,
    "quotationResponse": null,
    "reversalRequest": null,
    "reversalResponse": null,
    "rollbackRequest": null,
    "rollbackResponse": null,
    "transactionExistsRequest": null,
    "transactionExistsResponse": null,
    "findChangedTaxAreaIdsRequest": null,
    "findChangedTaxAreaIdsResponse": null,
    "isTaxAreaChangedRequest": null,
    "isTaxAreaChangedResponse": null,
    "taxAreaRequest": null,
    "taxAreaResponse": null,
    "findTaxAreasRequest": null,
    "findTaxAreasResponse": null,
    "versionRequest": null,
    "versionResponse": null,
    "applicationData": {
        "sender": null,
        "applicationProperty": [],
        "messageLogging": null,
        "logEntry": [],
        "responseTimeMS": 54.9
    },
    "ersrequest": null,
    "ersresponse": null,
    "apinvoiceSyncRequest": null,
    "apinvoiceSyncResponse": null,
    "arbillingSyncRequest": null,
    "arbillingSyncResponse": null
}`

var _ = Describe("Request Invoice API", func() {
	Describe("Invoice", func() {
		It("Invoice Unmarshal", func() {
			response := &Response{}
			err := json.Unmarshal([]byte(invoiceResponse), response)
			Expect(err).Should(BeNil())
		})

		It("Invoice Marshal", func() {
			request := Response{}
			request.InvoiceResponse.Seller.Company = "Rakuten"
			request.InvoiceResponse.Seller.Division = "RM_EU"
			request.InvoiceResponse.Seller.PhysicalOrigin.Country = "GB"
			request.InvoiceResponse.Seller.PhysicalOrigin.TaxAreaID = 808260000
			request.InvoiceResponse.Seller.TaxRegistration = []interface{}{}

			request.InvoiceResponse.Customer.CustomerCode.IsBusinessIndicator = true
			request.InvoiceResponse.Customer.Destination.Country = "GB"
			request.InvoiceResponse.Customer.Destination.TaxAreaID = 808260000

			request.InvoiceResponse.Customer.TaxRegistration = []TaxRegistration{{
				TaxRegistrationNumber:        "GB",
				NexusOverride:                []interface{}{},
				PhysicalLocation:             []interface{}{},
				ImpositionType:               nil,
				IsoCountryCode:               "GB",
				MainDivision:                 nil,
				HasPhysicalPresenceIndicator: true,
				JurisdictionID:               nil,
			}}
			request.InvoiceResponse.LineItem = []LineItem{{
				ImpositionToProcess:  []interface{}{},
				JurisdictionOverride: []interface{}{},
				LineItemNumber:       1,
				TotalTax:             20,
				Product:              Product{Value: "", ProductClass: "ASDE"},
				Quantity: Quantity{
					Value:         1,
					UnitOfMeasure: nil,
				},
				LineItem:        []interface{}{},
				FairMarketValue: 100,
				ExtendedPrice:   100,
				Taxes: []Taxes{{
					Jurisdiction: Jurisdiction{
						Value:             "UNITED KINGDOM",
						JurisdictionLevel: "COUNTRY",
						JurisdictionID:    78295,
					},
					CalculatedTax: 20,
					EffectiveRate: 0.2,
					Taxable: Taxable{
						Value: 100,
					},
					Imposition: Imposition{
						Value:        "VAT",
						ImpositionID: 1,
					},
					ImpositionType: ImpositionType{
						Value:            "VAT",
						ImpositionTypeID: 19,
					},
					TaxRuleID:             TaxRuleID{Value: 339387},
					SellerRegistrationID:  "GB 882.7552.85",
					BuyerRegistrationID:   "GB",
					Situs:                 "DESTINATION",
					RateClassification:    "Standard Rate",
					NominalRate:           0.2,
					TaxCollectedFromParty: "BUYER",
					TaxStructure:          "SINGLE_RATE",
					InputOutputType:       "OUTPUT",
					TaxType:               "VAT",
					TaxResult:             "TAXABLE",
					InvoiceTextCode:       []int{21},
				}},
				AssistedParameters: AssistedParameters{[]AssistedParameter{
					{Value: "GB",
						ParamName: "customer.taxRegistrationId",
						Phase:     "PRE",
						RuleName:  "Set Customer Registration ID"},
				}},
			}}

			request.InvoiceResponse.ImpositionToProcess = []interface{}{}
			request.InvoiceResponse.JurisdictionOverride = []interface{}{}
			request.InvoiceResponse.SubTotal = 100
			request.InvoiceResponse.Total = 120
			request.InvoiceResponse.TotalTax = 20
			request.InvoiceResponse.DocumentDate = 1582761600000
			request.InvoiceResponse.TransactionType = "SALE"

			request.ApplicationData.ResponseTimeMS = 54.9
			request.ApplicationData.ApplicationProperty = []interface{}{}
			request.ApplicationData.LogEntry = []interface{}{}

			body, err := json.Marshal(request)

			Expect(err).Should(BeNil())
			Expect(string(body)).To(MatchJSON(invoiceResponse))
		})
	})
})

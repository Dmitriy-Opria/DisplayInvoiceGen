package model

import (
	"encoding/json"
	"testing"
)

var requestBody = `{
  "indata": {
    "callingSystemNumber": "string",
    "companyId": 0,
    "companyName": "string",
    "companyRole": "string",
    "externalCompanyId": "string",
    "hostRequestInfo": {
      "hostRequestId": "string",
      "hostRequestLogEntryId": "string"
    },
    "hostSystem": "string",
    "invoice": [
      {
        "allocationGroupName": "string",
        "allocationGroupOwner": "string",
        "allocationName": "string",
        "autoCreateCertificates": "string",
        "autoCreateCustomers": "string",
        "basisPercent": "string",
        "billTo": {
          "address1": "string",
          "address2": "string",
          "address3": "string",
          "addressValidationMode": "string",
          "city": "string",
          "companyBranchId": "string",
          "country": "string",
          "county": "string",
          "defaultAddressValidationMode": "string",
          "district": "string",
          "geocode": "string",
          "isBonded": "string",
          "locationTaxCategory": "string",
          "postcode": "string",
          "province": "string",
          "state": "string"
        },
        "buyerPrimary": {
          "address1": "string",
          "address2": "string",
          "address3": "string",
          "addressValidationMode": "string",
          "city": "string",
          "companyBranchId": "string",
          "country": "string",
          "county": "string",
          "defaultAddressValidationMode": "string",
          "district": "string",
          "geocode": "string",
          "isBonded": "string",
          "locationTaxCategory": "string",
          "postcode": "string",
          "province": "string",
          "state": "string"
        },
        "calculationDirection": "string",
        "callingSystemNumber": "string",
        "companyId": 0,
        "companyName": "string",
        "companyRole": "string",
        "countryOfOrigin": "string",
        "currencyCode": "string",
        "customerGroupName": "string",
        "customerGroupOwner": "string",
        "customerName": "string",
        "customerNumber": "string",
        "deliveryTerms": "string",
        "deptOfConsign": "string",
        "documentType": "string",
        "endUse": [
          "string"
        ],
        "endUserName": "string",
        "establishments": {
          "buyerRole": {
            "billTo": "string",
            "buyerPrimary": "string",
            "middleman": "string",
            "orderAcceptance": "string",
            "orderOrigin": "string",
            "sellerPrimary": "string",
            "shipFrom": "string",
            "shipTo": "string",
            "supply": "string"
          },
          "middlemanRole": {
            "billTo": "string",
            "buyerPrimary": "string",
            "middleman": "string",
            "orderAcceptance": "string",
            "orderOrigin": "string",
            "sellerPrimary": "string",
            "shipFrom": "string",
            "shipTo": "string",
            "supply": "string"
          },
          "sellerRole": {
            "billTo": "string",
            "buyerPrimary": "string",
            "middleman": "string",
            "orderAcceptance": "string",
            "orderOrigin": "string",
            "sellerPrimary": "string",
            "shipFrom": "string",
            "shipTo": "string",
            "supply": "string"
          }
        },
        "exemptAmount": {
          "city": "string",
          "country": "string",
          "county": "string",
          "district": "string",
          "geocode": "string",
          "postcode": "string",
          "province": "string",
          "state": "string"
        },
        "exemptCertificate": {
          "city": "string",
          "country": "string",
          "county": "string",
          "district": "string",
          "geocode": "string",
          "postcode": "string",
          "province": "string",
          "state": "string"
        },
        "exemptReason": {
          "city": "string",
          "country": "string",
          "county": "string",
          "district": "string",
          "geocode": "string",
          "postcode": "string",
          "province": "string",
          "state": "string"
        },
        "externalCompanyId": "string",
        "filterGroupName": "string",
        "filterGroupOwner": "string",
        "fiscalDate": "string",
        "functionalCurrencyCode": "string",
        "hostSystem": "string",
        "inclusiveTaxIndicators": {
          "authorityType": [
            "string"
          ],
          "fullyInclusive": "string"
        },
        "inputRecoveryType": "string",
        "invoiceDate": "string",
        "invoiceNumber": "string",
        "isAuditUpdate": "string",
        "isAudited": "string",
        "isBusinessSupply": "string",
        "isCredit": "string",
        "isExempt": {
          "all": "string",
          "city": "string",
          "country": "string",
          "county": "string",
          "district": "string",
          "geocode": "string",
          "postcode": "string",
          "province": "string",
          "state": "string"
        },
        "isNoTax": {
          "all": "string",
          "city": "string",
          "country": "string",
          "county": "string",
          "district": "string",
          "geocode": "string",
          "postcode": "string",
          "province": "string",
          "state": "string"
        },
        "isReported": "string",
        "isReversed": "string",
        "isRounding": "string",
        "isSimplification": "string",
        "licenses": {
          "customerLicense": [
            {
              "number": "string",
              "type": "string"
            }
          ]
        },
        "line": [
          {
            "accountingCode": "string",
            "allocationGroupName": "string",
            "allocationGroupOwner": "string",
            "allocationName": "string",
            "basisPercent": "string",
            "billTo": {
              "address1": "string",
              "address2": "string",
              "address3": "string",
              "addressValidationMode": "string",
              "city": "string",
              "companyBranchId": "string",
              "country": "string",
              "county": "string",
              "defaultAddressValidationMode": "string",
              "district": "string",
              "geocode": "string",
              "isBonded": "string",
              "locationTaxCategory": "string",
              "postcode": "string",
              "province": "string",
              "state": "string"
            },
            "buyerPrimary": {
              "address1": "string",
              "address2": "string",
              "address3": "string",
              "addressValidationMode": "string",
              "city": "string",
              "companyBranchId": "string",
              "country": "string",
              "county": "string",
              "defaultAddressValidationMode": "string",
              "district": "string",
              "geocode": "string",
              "isBonded": "string",
              "locationTaxCategory": "string",
              "postcode": "string",
              "province": "string",
              "state": "string"
            },
            "commodityCode": "string",
            "countryOfOrigin": "string",
            "customerGroupName": "string",
            "customerGroupOwner": "string",
            "customerName": "string",
            "customerNumber": "string",
            "deliveryTerms": "string",
            "deptOfConsign": "string",
            "description": "string",
            "discountAmount": "string",
            "endUse": [
              "string"
            ],
            "endUserName": "string",
            "establishments": {
              "buyerRole": {
                "billTo": "string",
                "buyerPrimary": "string",
                "middleman": "string",
                "orderAcceptance": "string",
                "orderOrigin": "string",
                "sellerPrimary": "string",
                "shipFrom": "string",
                "shipTo": "string",
                "supply": "string"
              },
              "middlemanRole": {
                "billTo": "string",
                "buyerPrimary": "string",
                "middleman": "string",
                "orderAcceptance": "string",
                "orderOrigin": "string",
                "sellerPrimary": "string",
                "shipFrom": "string",
                "shipTo": "string",
                "supply": "string"
              },
              "sellerRole": {
                "billTo": "string",
                "buyerPrimary": "string",
                "middleman": "string",
                "orderAcceptance": "string",
                "orderOrigin": "string",
                "sellerPrimary": "string",
                "shipFrom": "string",
                "shipTo": "string",
                "supply": "string"
              }
            },
            "exemptAmount": {
              "city": "string",
              "country": "string",
              "county": "string",
              "district": "string",
              "geocode": "string",
              "postcode": "string",
              "province": "string",
              "state": "string"
            },
            "exemptCertificate": {
              "city": "string",
              "country": "string",
              "county": "string",
              "district": "string",
              "geocode": "string",
              "postcode": "string",
              "province": "string",
              "state": "string"
            },
            "exemptReason": {
              "city": "string",
              "country": "string",
              "county": "string",
              "district": "string",
              "geocode": "string",
              "postcode": "string",
              "province": "string",
              "state": "string"
            },
            "freightOnBoard": "string",
            "grossAmount": "string",
            "grossPlusTax": "string",
            "id": "string",
            "inclusiveTaxIndicators": {
              "authorityType": [
                "string"
              ],
              "fullyInclusive": "string"
            },
            "inputRecoveryAmount": "string",
            "inputRecoveryPercent": "string",
            "inputRecoveryType": "string",
            "invoiceDate": "string",
            "isAllocatable": "string",
            "isBusinessSupply": "string",
            "isCredit": "string",
            "isExempt": {
              "all": "string",
              "city": "string",
              "country": "string",
              "county": "string",
              "district": "string",
              "geocode": "string",
              "postcode": "string",
              "province": "string",
              "state": "string"
            },
            "isManufacturing": "string",
            "isNoTax": {
              "all": "string",
              "city": "string",
              "country": "string",
              "county": "string",
              "district": "string",
              "geocode": "string",
              "postcode": "string",
              "province": "string",
              "state": "string"
            },
            "isSimplification": "string",
            "itemValue": "string",
            "licenses": {
              "customerLicense": [
                {
                  "number": "string",
                  "type": "string"
                }
              ]
            },
            "lineNumber": 0,
            "location": {
              "billTo": "string",
              "middleman": "string",
              "orderAcceptance": "string",
              "orderOrigin": "string",
              "shipFrom": "string",
              "shipTo": "string",
              "supply": "string"
            },
            "locationSet": "string",
            "mass": 0,
            "middleman": {
              "address1": "string",
              "address2": "string",
              "address3": "string",
              "addressValidationMode": "string",
              "city": "string",
              "companyBranchId": "string",
              "country": "string",
              "county": "string",
              "defaultAddressValidationMode": "string",
              "district": "string",
              "geocode": "string",
              "isBonded": "string",
              "locationTaxCategory": "string",
              "postcode": "string",
              "province": "string",
              "state": "string"
            },
            "modeOfTransport": "string",
            "movementDate": "string",
            "movementType": "string",
            "orderAcceptance": {
              "address1": "string",
              "address2": "string",
              "address3": "string",
              "addressValidationMode": "string",
              "city": "string",
              "companyBranchId": "string",
              "country": "string",
              "county": "string",
              "defaultAddressValidationMode": "string",
              "district": "string",
              "geocode": "string",
              "isBonded": "string",
              "locationTaxCategory": "string",
              "postcode": "string",
              "province": "string",
              "state": "string"
            },
            "orderOrigin": {
              "address1": "string",
              "address2": "string",
              "address3": "string",
              "addressValidationMode": "string",
              "city": "string",
              "companyBranchId": "string",
              "country": "string",
              "county": "string",
              "defaultAddressValidationMode": "string",
              "district": "string",
              "geocode": "string",
              "isBonded": "string",
              "locationTaxCategory": "string",
              "postcode": "string",
              "province": "string",
              "state": "string"
            },
            "originalDocumentId": "string",
            "originalDocumentItem": "string",
            "originalDocumentType": "string",
            "originalInvoiceDate": "string",
            "originalMovementDate": "string",
            "overrideAmount": {
              "city": "string",
              "country": "string",
              "county": "string",
              "district": "string",
              "geocode": "string",
              "postcode": "string",
              "province": "string",
              "state": "string"
            },
            "overrideRate": {
              "city": "string",
              "country": "string",
              "county": "string",
              "district": "string",
              "geocode": "string",
              "postcode": "string",
              "province": "string",
              "state": "string"
            },
            "partNumber": "string",
            "pointOfTitleTransfer": "string",
            "portOfEntry": "string",
            "portOfLoading": "string",
            "productCode": "string",
            "quantities": {
              "quantity": [
                {
                  "amount": "string",
                  "default": "string",
                  "uom": "string"
                }
              ]
            },
            "quantity": 0,
            "regime": "string",
            "registrations": {
              "buyerRole": [
                "string"
              ],
              "middlemanRole": [
                "string"
              ],
              "sellerRole": [
                "string"
              ]
            },
            "relatedLineNumber": 0,
            "sellerPrimary": {
              "address1": "string",
              "address2": "string",
              "address3": "string",
              "addressValidationMode": "string",
              "city": "string",
              "companyBranchId": "string",
              "country": "string",
              "county": "string",
              "defaultAddressValidationMode": "string",
              "district": "string",
              "geocode": "string",
              "isBonded": "string",
              "locationTaxCategory": "string",
              "postcode": "string",
              "province": "string",
              "state": "string"
            },
            "shipFrom": {
              "address1": "string",
              "address2": "string",
              "address3": "string",
              "addressValidationMode": "string",
              "city": "string",
              "companyBranchId": "string",
              "country": "string",
              "county": "string",
              "defaultAddressValidationMode": "string",
              "district": "string",
              "geocode": "string",
              "isBonded": "string",
              "locationTaxCategory": "string",
              "postcode": "string",
              "province": "string",
              "state": "string"
            },
            "shipTo": {
              "address1": "string",
              "address2": "string",
              "address3": "string",
              "addressValidationMode": "string",
              "city": "string",
              "companyBranchId": "string",
              "country": "string",
              "county": "string",
              "defaultAddressValidationMode": "string",
              "district": "string",
              "geocode": "string",
              "isBonded": "string",
              "locationTaxCategory": "string",
              "postcode": "string",
              "province": "string",
              "state": "string"
            },
            "supplementaryUnit": "string",
            "supply": {
              "address1": "string",
              "address2": "string",
              "address3": "string",
              "addressValidationMode": "string",
              "city": "string",
              "companyBranchId": "string",
              "country": "string",
              "county": "string",
              "defaultAddressValidationMode": "string",
              "district": "string",
              "geocode": "string",
              "isBonded": "string",
              "locationTaxCategory": "string",
              "postcode": "string",
              "province": "string",
              "state": "string"
            },
            "supplyExemptPercent": {
              "city": "string",
              "country": "string",
              "county": "string",
              "district": "string",
              "geocode": "string",
              "postcode": "string",
              "province": "string",
              "state": "string"
            },
            "taxAmount": "string",
            "taxCode": "string",
            "taxDeterminationDate": "string",
            "taxExchangeRateDate": "string",
            "taxPointDate": "string",
            "taxTreatment": "string",
            "taxType": {
              "all": "string",
              "city": "string",
              "country": "string",
              "county": "string",
              "district": "string",
              "geocode": "string",
              "postcode": "string",
              "province": "string",
              "state": "string"
            },
            "titleTransferLocation": "string",
            "transactionType": "string",
            "uniqueLineNumber": "string",
            "unitOfMeasure": "string",
            "uom": "string",
            "userElement": [
              {
                "name": "string",
                "value": "string"
              }
            ],
            "vatGroupRegistration": "string",
            "vendorName": "string",
            "vendorNumber": "string",
            "vendorTax": "string"
          }
        ],
        "location": {
          "billTo": "string",
          "middleman": "string",
          "orderAcceptance": "string",
          "orderOrigin": "string",
          "shipFrom": "string",
          "shipTo": "string",
          "supply": "string"
        },
        "locationSet": "string",
        "middleman": {
          "address1": "string",
          "address2": "string",
          "address3": "string",
          "addressValidationMode": "string",
          "city": "string",
          "companyBranchId": "string",
          "country": "string",
          "county": "string",
          "defaultAddressValidationMode": "string",
          "district": "string",
          "geocode": "string",
          "isBonded": "string",
          "locationTaxCategory": "string",
          "postcode": "string",
          "province": "string",
          "state": "string"
        },
        "modeOfTransport": "string",
        "movementDate": "string",
        "movementType": "string",
        "natureOfTransactionCode": "string",
        "orderAcceptance": {
          "address1": "string",
          "address2": "string",
          "address3": "string",
          "addressValidationMode": "string",
          "city": "string",
          "companyBranchId": "string",
          "country": "string",
          "county": "string",
          "defaultAddressValidationMode": "string",
          "district": "string",
          "geocode": "string",
          "isBonded": "string",
          "locationTaxCategory": "string",
          "postcode": "string",
          "province": "string",
          "state": "string"
        },
        "orderOrigin": {
          "address1": "string",
          "address2": "string",
          "address3": "string",
          "addressValidationMode": "string",
          "city": "string",
          "companyBranchId": "string",
          "country": "string",
          "county": "string",
          "defaultAddressValidationMode": "string",
          "district": "string",
          "geocode": "string",
          "isBonded": "string",
          "locationTaxCategory": "string",
          "postcode": "string",
          "province": "string",
          "state": "string"
        },
        "originalDocumentId": "string",
        "originalDocumentItem": "string",
        "originalDocumentType": "string",
        "originalInvoiceDate": "string",
        "originalInvoiceNumber": "string",
        "originalMovementDate": "string",
        "overrideAmount": {
          "city": "string",
          "country": "string",
          "county": "string",
          "district": "string",
          "geocode": "string",
          "postcode": "string",
          "province": "string",
          "state": "string"
        },
        "overrideRate": {
          "city": "string",
          "country": "string",
          "county": "string",
          "district": "string",
          "geocode": "string",
          "postcode": "string",
          "province": "string",
          "state": "string"
        },
        "pointOfTitleTransfer": "string",
        "portOfEntry": "string",
        "portOfLoading": "string",
        "productMappingGroupName": "string",
        "productMappingGroupOwner": "string",
        "regime": "string",
        "registrations": {
          "buyerRole": [
            "string"
          ],
          "middlemanRole": [
            "string"
          ],
          "sellerRole": [
            "string"
          ]
        },
        "sellerPrimary": {
          "address1": "string",
          "address2": "string",
          "address3": "string",
          "addressValidationMode": "string",
          "city": "string",
          "companyBranchId": "string",
          "country": "string",
          "county": "string",
          "defaultAddressValidationMode": "string",
          "district": "string",
          "geocode": "string",
          "isBonded": "string",
          "locationTaxCategory": "string",
          "postcode": "string",
          "province": "string",
          "state": "string"
        },
        "shipFrom": {
          "address1": "string",
          "address2": "string",
          "address3": "string",
          "addressValidationMode": "string",
          "city": "string",
          "companyBranchId": "string",
          "country": "string",
          "county": "string",
          "defaultAddressValidationMode": "string",
          "district": "string",
          "geocode": "string",
          "isBonded": "string",
          "locationTaxCategory": "string",
          "postcode": "string",
          "province": "string",
          "state": "string"
        },
        "shipTo": {
          "address1": "string",
          "address2": "string",
          "address3": "string",
          "addressValidationMode": "string",
          "city": "string",
          "companyBranchId": "string",
          "country": "string",
          "county": "string",
          "defaultAddressValidationMode": "string",
          "district": "string",
          "geocode": "string",
          "isBonded": "string",
          "locationTaxCategory": "string",
          "postcode": "string",
          "province": "string",
          "state": "string"
        },
        "statisticalProcedure": "string",
        "supply": {
          "address1": "string",
          "address2": "string",
          "address3": "string",
          "addressValidationMode": "string",
          "city": "string",
          "companyBranchId": "string",
          "country": "string",
          "county": "string",
          "defaultAddressValidationMode": "string",
          "district": "string",
          "geocode": "string",
          "isBonded": "string",
          "locationTaxCategory": "string",
          "postcode": "string",
          "province": "string",
          "state": "string"
        },
        "supplyExemptPercent": {
          "city": "string",
          "country": "string",
          "county": "string",
          "district": "string",
          "geocode": "string",
          "postcode": "string",
          "province": "string",
          "state": "string"
        },
        "taxCode": "string",
        "taxDeterminationDate": "string",
        "taxExchangeRateDate": "string",
        "taxPointDate": "string",
        "taxTreatment": "string",
        "taxType": {
          "all": "string",
          "city": "string",
          "country": "string",
          "county": "string",
          "district": "string",
          "geocode": "string",
          "postcode": "string",
          "province": "string",
          "state": "string"
        },
        "titleTransferLocation": "string",
        "transactionType": "string",
        "uniqueInvoiceNumber": "string",
        "userElement": [
          {
            "name": "string",
            "value": "string"
          }
        ],
        "vatGroupRegistration": "string",
        "vendorName": "string",
        "vendorNumber": "string",
        "vendorTax": "string"
      }
    ],
    "scenarioId": 0,
    "scenarioName": "string",
    "version": "G"
  }
}`

func TestRequestParsing(t *testing.T) {
	request := &TaxRequest{}
	err := json.Unmarshal([]byte(requestBody), request)
	if err != nil {
		t.Errorf("Can't parse request body: %v", err)
	}
}

package model

import (
	"encoding/json"
	"testing"
)

var responseBody = `{
  "outdata": {
    "companyId": 0,
    "companyName": "string",
    "companyRole": "string",
    "externalCompanyId": "string",
    "invoice": [
      {
        "basisPercent": "string",
        "calculationDirection": "string",
        "callingSystemNumber": "string",
        "companyId": 0,
        "companyName": "string",
        "companyRole": "string",
        "currencyCode": "string",
        "currencyName": "string",
        "customerGroupName": "string",
        "customerGroupOwner": "string",
        "customerName": "string",
        "customerNumber": "string",
        "endUserName": "string",
        "externalCompanyId": "string",
        "fiscalDate": "string",
        "functionalCurrencyCode": "string",
        "hostSystem": "string",
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
          "password": "string",
          "scenarioId": 0,
          "scenarioName": "string",
          "username": "string",
          "version": "G",
          "xmlGroupName": "string",
          "xmlGroupOwner": "string"
        },
        "invoiceDate": "string",
        "invoiceNumber": "string",
        "isAuditUpdate": "string",
        "isBusinessSupply": "string",
        "isCredit": "string",
        "isReported": "string",
        "isReversed": "string",
        "line": [
          {
            "accountingCode": "string",
            "allocationLine": [
              {}
            ],
            "allocationName": "string",
            "basisPercent": "string",
            "billToBranchId": "string",
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
            "endUserName": "string",
            "freightOnBoard": "string",
            "grossAmount": "string",
            "id": "string",
            "inputRecoveryAmount": "string",
            "inputRecoveryPercent": "string",
            "invoiceDate": "string",
            "isBusinessSupply": "string",
            "isCredit": "string",
            "itemValue": "string",
            "lineNumber": 0,
            "mass": 0,
            "message": [
              {
                "category": "string",
                "code": "string",
                "location": "string",
                "messageText": "string",
                "severity": 0
              }
            ],
            "middlemanBranchId": "string",
            "middlemanMarkupAmount": 0,
            "middlemanMarkupRate": 0,
            "modeOfTransport": "string",
            "movementDate": "string",
            "originalDocumentId": "string",
            "originalDocumentItem": "string",
            "originalDocumentType": "string",
            "originalInvoiceDate": "string",
            "originalInvoiceNumber": "string",
            "originalMovementDate": "string",
            "partNumber": "string",
            "pointOfTitleTransfer": "string",
            "portOfEntry": "string",
            "portOfLoading": "string",
            "quantities": {
              "quantity": [
                {
                  "amount": "string",
                  "default": "string",
                  "uom": "string"
                }
              ]
            },
            "regime": "string",
            "relatedLineNumber": 0,
            "shipFromBranchId": "string",
            "shipFromCountry": "string",
            "shipToBranchId": "string",
            "shipToCountry": "string",
            "supplementaryUnit": "string",
            "supplyBranchId": "string",
            "tax": [
              {
                "addressType": "string",
                "adminZoneLevel": "string",
                "authorityAttribute": [
                  {
                    "name": "string",
                    "value": "string"
                  }
                ],
                "authorityCategory": "string",
                "authorityCurrencyCode": "string",
                "authorityFips": "string",
                "authorityName": "string",
                "authorityOfficialName": "string",
                "authorityType": "string",
                "authorityTypeAlias": "string",
                "authorityUuid": "string",
                "basisPercent": "string",
                "buyerRegistration": "string",
                "calculationMethod": "string",
                "comment": "string",
                "currencyConversion": [
                  {
                    "conversionSteps": [
                      {
                        "conversionStep": 0,
                        "exchangeRate": 0,
                        "fromCurrencyCode": "string",
                        "toCurrencyCode": "string"
                      }
                    ],
                    "exchangeRateSource": "string",
                    "taxExchangeRateDate": "string"
                  }
                ],
                "effectiveZoneLevel": "string",
                "erpTaxCode": "string",
                "euTransaction": true,
                "exemptAmount": {
                  "authorityAmount": "string",
                  "documentAmount": "string",
                  "unroundedAuthorityAmount": "string",
                  "unroundedDocumentAmount": "string"
                },
                "exemptCertificate": "string",
                "exemptCertificateExpireDate": "string",
                "exemptReason": "string",
                "fee": [
                  {
                    "amount": 0,
                    "currencyCode": "string",
                    "unitOfMeasure": "string"
                  }
                ],
                "fiscalRepAddress1": "string",
                "fiscalRepAddress2": "string",
                "fiscalRepContact": "string",
                "fiscalRepName": "string",
                "grossAmount": {
                  "authorityAmount": "string",
                  "documentAmount": "string",
                  "unroundedAuthorityAmount": "string",
                  "unroundedDocumentAmount": "string"
                },
                "inclusiveTax": "string",
                "inputRecoveryAmount": "string",
                "inputRecoveryPercent": "string",
                "invoiceDescription": "string",
                "isExempt": true,
                "isIntrastatReported": "string",
                "isNotax": true,
                "isTriangulation": true,
                "isVatReported": true,
                "isViesReported": true,
                "jurisdictionText": "string",
                "licenses": {
                  "license": [
                    {
                      "licenseCategory": "string",
                      "licenseEndDate": "string",
                      "licenseExternalIdentifier": "string",
                      "licenseNumber": "string",
                      "licenseTypeName": "string"
                    }
                  ]
                },
                "locationCode": "string",
                "message": [
                  {
                    "category": "string",
                    "code": "string",
                    "location": "string",
                    "messageText": "string",
                    "severity": 0
                  }
                ],
                "middlemanRegistration": "string",
                "natureOfTax": "string",
                "nonTaxableBasis": {
                  "authorityAmount": "string",
                  "documentAmount": "string",
                  "unroundedAuthorityAmount": "string",
                  "unroundedDocumentAmount": "string"
                },
                "overrideAmount": 0,
                "overrideRate": 0,
                "registrationAttribute1": "string",
                "registrationAttribute10": "string",
                "registrationAttribute11": "string",
                "registrationAttribute12": "string",
                "registrationAttribute13": "string",
                "registrationAttribute14": "string",
                "registrationAttribute15": "string",
                "registrationAttribute16": "string",
                "registrationAttribute17": "string",
                "registrationAttribute18": "string",
                "registrationAttribute19": "string",
                "registrationAttribute2": "string",
                "registrationAttribute20": "string",
                "registrationAttribute21": "string",
                "registrationAttribute22": "string",
                "registrationAttribute23": "string",
                "registrationAttribute24": "string",
                "registrationAttribute25": "string",
                "registrationAttribute26": "string",
                "registrationAttribute27": "string",
                "registrationAttribute28": "string",
                "registrationAttribute29": "string",
                "registrationAttribute3": "string",
                "registrationAttribute30": "string",
                "registrationAttribute31": "string",
                "registrationAttribute32": "string",
                "registrationAttribute33": "string",
                "registrationAttribute34": "string",
                "registrationAttribute35": "string",
                "registrationAttribute36": "string",
                "registrationAttribute37": "string",
                "registrationAttribute38": "string",
                "registrationAttribute39": "string",
                "registrationAttribute4": "string",
                "registrationAttribute40": "string",
                "registrationAttribute41": "string",
                "registrationAttribute42": "string",
                "registrationAttribute43": "string",
                "registrationAttribute44": "string",
                "registrationAttribute45": "string",
                "registrationAttribute46": "string",
                "registrationAttribute47": "string",
                "registrationAttribute48": "string",
                "registrationAttribute49": "string",
                "registrationAttribute5": "string",
                "registrationAttribute50": "string",
                "registrationAttribute6": "string",
                "registrationAttribute7": "string",
                "registrationAttribute8": "string",
                "registrationAttribute9": "string",
                "relatedAllocationLineNumber": 0,
                "relatedLineNumber": 0,
                "revisedGrossAmount": 0,
                "ruleOrder": 0,
                "ruleReportingCategory": "string",
                "sellerRegistration": "string",
                "supplyExemptPercent": 0,
                "taxAmount": {
                  "authorityAmount": "string",
                  "documentAmount": "string",
                  "unroundedAuthorityAmount": "string",
                  "unroundedDocumentAmount": "string"
                },
                "taxDeterminationDate": "string",
                "taxDirection": "string",
                "taxPointDate": "string",
                "taxRate": "string",
                "taxRateCode": "string",
                "taxTreatment": "string",
                "taxType": "string",
                "taxableBasis": {
                  "authorityAmount": "string",
                  "documentAmount": "string",
                  "unroundedAuthorityAmount": "string",
                  "unroundedDocumentAmount": "string"
                },
                "taxableCity": "string",
                "taxableCountry": "string",
                "taxableCountryName": "string",
                "taxableCounty": "string",
                "taxableDistrict": "string",
                "taxableGeocode": "string",
                "taxablePostcode": "string",
                "taxableProvince": "string",
                "taxableState": "string",
                "uomConversion": {
                  "factor": "string",
                  "from": {
                    "amount": "string",
                    "default": "string",
                    "uom": "string"
                  },
                  "operator": "string",
                  "toRounded": {
                    "amount": "string",
                    "default": "string",
                    "uom": "string"
                  },
                  "toUnrounded": {
                    "amount": "string",
                    "default": "string",
                    "uom": "string"
                  }
                },
                "vatGroupRegistration": "string",
                "zoneLevel": "string",
                "zoneName": "string"
              }
            ],
            "taxCode": "string",
            "taxSummary": {
              "advisories": {
                "advisory": [
                  "string"
                ]
              },
              "effectiveTaxRate": "string",
              "exemptAmount": "string",
              "nonTaxableBasis": "string",
              "taxRate": "string",
              "taxableBasis": "string"
            },
            "titleTransferLocation": "string",
            "totalTaxAmount": "string",
            "transactionType": "string",
            "uniqueLineNumber": "string",
            "unitOfMeasure": "string",
            "userElement": [
              {
                "name": "string",
                "value": "string"
              }
            ],
            "vendorName": "string",
            "vendorNumber": "string",
            "vendorTax": "string"
          }
        ],
        "message": [
          {
            "category": "string",
            "code": "string",
            "location": "string",
            "messageText": "string",
            "severity": 0
          }
        ],
        "minAccountableUnit": "string",
        "natureOfTransactionCode": "string",
        "originalDocumentId": "string",
        "originalDocumentItem": "string",
        "originalDocumentType": "string",
        "originalInvoiceDate": "string",
        "originalInvoiceNumber": "string",
        "originalMovementDate": "string",
        "requestStatus": {
          "error": [
            {
              "code": "string",
              "description": "string",
              "errorLocationPath": "string"
            }
          ],
          "isPartialSuccess": true,
          "isSuccess": true
        },
        "roundingPrecision": 0,
        "roundingRule": "string",
        "statisticalProcedure": "string",
        "taxSummary": {
          "advisories": {
            "advisory": [
              "string"
            ]
          },
          "effectiveTaxRate": "string",
          "exemptAmount": "string",
          "nonTaxableBasis": "string",
          "taxRate": "string",
          "taxableBasis": "string"
        },
        "totalTaxAmount": "string",
        "transactionDate": "string",
        "uniqueInvoiceNumber": "string",
        "userElement": [
          {
            "name": "string",
            "value": "string"
          }
        ],
        "vendorName": "string",
        "vendorNumber": "string",
        "vendorTax": "string"
      }
    ],
    "requestStatus": {
      "error": [
        {
          "code": "string",
          "description": "string",
          "errorLocationPath": "string"
        }
      ],
      "isPartialSuccess": true,
      "isSuccess": true
    },
    "scenarioName": "string",
    "version": "G"
  }
}`

func TestResponseParsing(t *testing.T) {
	response := &TaxResponse{}
	err := json.Unmarshal([]byte(responseBody), response)
	if err != nil {
		t.Errorf("Can't parse response body: %v", err)
	}
}

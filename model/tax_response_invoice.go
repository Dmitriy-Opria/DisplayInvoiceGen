package model

type Response struct {
	Login struct {
		UserName  interface{} `json:"userName"`
		Password  interface{} `json:"password"`
		TrustedID interface{} `json:"trustedId"`
	} `json:"login"`
	AccrualRequest                   interface{} `json:"accrualRequest"`
	AccrualResponse                  interface{} `json:"accrualResponse"`
	AccrualSyncRequest               interface{} `json:"accrualSyncRequest"`
	AccrualSyncResponse              interface{} `json:"accrualSyncResponse"`
	AssetMovementRequest             interface{} `json:"assetMovementRequest"`
	AssetMovementResponse            interface{} `json:"assetMovementResponse"`
	BuyerInputTaxRequest             interface{} `json:"buyerInputTaxRequest"`
	BuyerInputTaxResponse            interface{} `json:"buyerInputTaxResponse"`
	DeleteRequest                    interface{} `json:"deleteRequest"`
	DeleteResponse                   interface{} `json:"deleteResponse"`
	DistributeTaxProcurementRequest  interface{} `json:"distributeTaxProcurementRequest"`
	DistributeTaxProcurementResponse interface{} `json:"distributeTaxProcurementResponse"`
	DistributeTaxRequest             interface{} `json:"distributeTaxRequest"`
	DistributeTaxResponse            interface{} `json:"distributeTaxResponse"`
	InventoryRemovalRequest          interface{} `json:"inventoryRemovalRequest"`
	InventoryRemovalResponse         interface{} `json:"inventoryRemovalResponse"`
	InvoiceRequest                   interface{} `json:"invoiceRequest"`
	InvoiceResponse                  struct {
		Currency            interface{} `json:"currency"`
		OriginalCurrency    interface{} `json:"originalCurrency"`
		CompanyCodeCurrency interface{} `json:"companyCodeCurrency"`
		Seller              struct {
			Company         string      `json:"company"`
			Division        string      `json:"division"`
			Department      interface{} `json:"department"`
			UtilityProvider interface{} `json:"utilityProvider"`
			Dispatcher      interface{} `json:"dispatcher"`
			PhysicalOrigin  struct {
				StreetAddress1           interface{} `json:"streetAddress1"`
				StreetAddress2           interface{} `json:"streetAddress2"`
				City                     interface{} `json:"city"`
				MainDivision             interface{} `json:"mainDivision"`
				SubDivision              interface{} `json:"subDivision"`
				PostalCode               interface{} `json:"postalCode"`
				Country                  string      `json:"country"`
				CurrencyConversion       interface{} `json:"currencyConversion"`
				TaxAreaID                int         `json:"taxAreaId"`
				Latitude                 interface{} `json:"latitude"`
				Longitude                interface{} `json:"longitude"`
				LocationCustomsStatus    interface{} `json:"locationCustomsStatus"`
				LocationCode             interface{} `json:"locationCode"`
				ExternalJurisdictionCode interface{} `json:"externalJurisdictionCode"`
			} `json:"physicalOrigin"`
			AdministrativeOrigin interface{}   `json:"administrativeOrigin"`
			TaxRegistration      []interface{} `json:"taxRegistration"`
			NexusIndicator       interface{}   `json:"nexusIndicator"`
			NexusReasonCode      interface{}   `json:"nexusReasonCode"`
		} `json:"seller"`
		Customer struct {
			CustomerCode struct {
				Value               string      `json:"value"`
				ClassCode           interface{} `json:"classCode"`
				IsBusinessIndicator bool        `json:"isBusinessIndicator"`
			} `json:"customerCode"`
			Destination struct {
				StreetAddress1           interface{} `json:"streetAddress1"`
				StreetAddress2           interface{} `json:"streetAddress2"`
				City                     interface{} `json:"city"`
				MainDivision             interface{} `json:"mainDivision"`
				SubDivision              interface{} `json:"subDivision"`
				PostalCode               interface{} `json:"postalCode"`
				Country                  string      `json:"country"`
				CurrencyConversion       interface{} `json:"currencyConversion"`
				TaxAreaID                int         `json:"taxAreaId"`
				Latitude                 interface{} `json:"latitude"`
				Longitude                interface{} `json:"longitude"`
				LocationCustomsStatus    interface{} `json:"locationCustomsStatus"`
				LocationCode             interface{} `json:"locationCode"`
				ExternalJurisdictionCode interface{} `json:"externalJurisdictionCode"`
			} `json:"destination"`
			AdministrativeDestination interface{}       `json:"administrativeDestination"`
			ExemptionCertificate      interface{}       `json:"exemptionCertificate"`
			TaxRegistration           []TaxRegistration `json:"taxRegistration"`
			IsTaxExempt               bool              `json:"isTaxExempt"`
			ExemptionReasonCode       interface{}       `json:"exemptionReasonCode"`
		} `json:"customer"`
		TaxOverride                       interface{}   `json:"taxOverride"`
		ImpositionToProcess               []interface{} `json:"impositionToProcess"`
		JurisdictionOverride              []interface{} `json:"jurisdictionOverride"`
		SitusOverride                     interface{}   `json:"situsOverride"`
		Discount                          interface{}   `json:"discount"`
		ProratePercentage                 interface{}   `json:"proratePercentage"`
		SubTotal                          float64       `json:"subTotal"`
		Total                             float64       `json:"total"`
		TotalTax                          float64       `json:"totalTax"`
		CurrencyConversionFactors         interface{}   `json:"currencyConversionFactors"`
		LineItem                          []LineItem    `json:"lineItem"`
		DocumentNumber                    interface{}   `json:"documentNumber"`
		AccumulationDocumentNumber        interface{}   `json:"accumulationDocumentNumber"`
		AccumulationCustomerNumber        interface{}   `json:"accumulationCustomerNumber"`
		DocumentType                      interface{}   `json:"documentType"`
		BillingType                       interface{}   `json:"billingType"`
		OrderType                         interface{}   `json:"orderType"`
		PostingDate                       interface{}   `json:"postingDate"`
		LocationCode                      interface{}   `json:"locationCode"`
		ReturnAssistedParametersIndicator interface{}   `json:"returnAssistedParametersIndicator"`
		ReturnGeneratedLineItemsIndicator interface{}   `json:"returnGeneratedLineItemsIndicator"`
		DeliveryTerm                      interface{}   `json:"deliveryTerm"`
		DocumentDate                      int64         `json:"documentDate"`
		TransactionID                     interface{}   `json:"transactionId"`
		TransactionType                   string        `json:"transactionType"`
		SimplificationCode                interface{}   `json:"simplificationCode"`
		RoundAtLineLevel                  interface{}   `json:"roundAtLineLevel"`
		PaymentDate                       interface{}   `json:"paymentDate"`
		DocumentSequenceID                interface{}   `json:"documentSequenceId"`
		TaxPointDate                      interface{}   `json:"taxPointDate"`
	} `json:"invoiceResponse"`
	InvoiceVerificationRequest    interface{} `json:"invoiceVerificationRequest"`
	InvoiceVerificationResponse   interface{} `json:"invoiceVerificationResponse"`
	PurchaseOrderRequest          interface{} `json:"purchaseOrderRequest"`
	PurchaseOrderResponse         interface{} `json:"purchaseOrderResponse"`
	QuotationRequest              interface{} `json:"quotationRequest"`
	QuotationResponse             interface{} `json:"quotationResponse"`
	ReversalRequest               interface{} `json:"reversalRequest"`
	ReversalResponse              interface{} `json:"reversalResponse"`
	RollbackRequest               interface{} `json:"rollbackRequest"`
	RollbackResponse              interface{} `json:"rollbackResponse"`
	TransactionExistsRequest      interface{} `json:"transactionExistsRequest"`
	TransactionExistsResponse     interface{} `json:"transactionExistsResponse"`
	FindChangedTaxAreaIdsRequest  interface{} `json:"findChangedTaxAreaIdsRequest"`
	FindChangedTaxAreaIdsResponse interface{} `json:"findChangedTaxAreaIdsResponse"`
	IsTaxAreaChangedRequest       interface{} `json:"isTaxAreaChangedRequest"`
	IsTaxAreaChangedResponse      interface{} `json:"isTaxAreaChangedResponse"`
	TaxAreaRequest                interface{} `json:"taxAreaRequest"`
	TaxAreaResponse               interface{} `json:"taxAreaResponse"`
	FindTaxAreasRequest           interface{} `json:"findTaxAreasRequest"`
	FindTaxAreasResponse          interface{} `json:"findTaxAreasResponse"`
	VersionRequest                interface{} `json:"versionRequest"`
	VersionResponse               interface{} `json:"versionResponse"`
	ApplicationData               struct {
		Sender              interface{}   `json:"sender"`
		ApplicationProperty []interface{} `json:"applicationProperty"`
		MessageLogging      interface{}   `json:"messageLogging"`
		LogEntry            []interface{} `json:"logEntry"`
		ResponseTimeMS      float64       `json:"responseTimeMS"`
	} `json:"applicationData"`
	Ersrequest            interface{} `json:"ersrequest"`
	Ersresponse           interface{} `json:"ersresponse"`
	ApinvoiceSyncRequest  interface{} `json:"apinvoiceSyncRequest"`
	ApinvoiceSyncResponse interface{} `json:"apinvoiceSyncResponse"`
	ArbillingSyncRequest  interface{} `json:"arbillingSyncRequest"`
	ArbillingSyncResponse interface{} `json:"arbillingSyncResponse"`
}

type TaxRegistration struct {
	TaxRegistrationNumber        string        `json:"taxRegistrationNumber"`
	NexusOverride                []interface{} `json:"nexusOverride"`
	PhysicalLocation             []interface{} `json:"physicalLocation"`
	ImpositionType               interface{}   `json:"impositionType"`
	IsoCountryCode               string        `json:"isoCountryCode"`
	MainDivision                 interface{}   `json:"mainDivision"`
	HasPhysicalPresenceIndicator bool          `json:"hasPhysicalPresenceIndicator"`
	JurisdictionID               interface{}   `json:"jurisdictionId"`
}

type LineItem struct {
	Seller                           interface{}   `json:"seller"`
	Customer                         interface{}   `json:"customer"`
	TaxOverride                      interface{}   `json:"taxOverride"`
	ImpositionToProcess              []interface{} `json:"impositionToProcess"`
	JurisdictionOverride             []interface{} `json:"jurisdictionOverride"`
	SitusOverride                    interface{}   `json:"situsOverride"`
	Product                          Product       `json:"product"`
	LineType                         interface{}   `json:"lineType"`
	CommodityCode                    interface{}   `json:"commodityCode"`
	Quantity                         `json:"quantity"`
	Weight                           interface{} `json:"weight"`
	Volume                           interface{} `json:"volume"`
	SupplementaryUnit                interface{} `json:"supplementaryUnit"`
	StatisticalValue                 interface{} `json:"statisticalValue"`
	Freight                          interface{} `json:"freight"`
	FairMarketValue                  float64     `json:"fairMarketValue"`
	UnitPrice                        float64     `json:"unitPrice"`
	ExtendedPrice                    float64     `json:"extendedPrice"`
	Discount                         interface{} `json:"discount"`
	AmountBilledToDate               interface{} `json:"amountBilledToDate"`
	CompanyCodeCurrencyTaxableAmount interface{} `json:"companyCodeCurrencyTaxableAmount"`
	CompanyCodeCurrencyTaxAmount     interface{} `json:"companyCodeCurrencyTaxAmount"`
	Taxes                            []Taxes     `json:"taxes"`
	TotalTax                         float64     `json:"totalTax"`
	FlexibleFields                   interface{} `json:"flexibleFields"`
	ReturnsFields                    interface{} `json:"returnsFields"`
	AssistedParameters               `json:"assistedParameters"`
	LineItem                         []interface{} `json:"lineItem"`
	LineItemNumber                   int           `json:"lineItemNumber"`
	TaxDate                          interface{}   `json:"taxDate"`
	IsMulticomponent                 interface{}   `json:"isMulticomponent"`
	LocationCode                     interface{}   `json:"locationCode"`
	DeliveryTerm                     interface{}   `json:"deliveryTerm"`
	PostingDate                      interface{}   `json:"postingDate"`
	CostCenter                       interface{}   `json:"costCenter"`
	DepartmentCode                   interface{}   `json:"departmentCode"`
	GeneralLedgerAccount             interface{}   `json:"generalLedgerAccount"`
	MaterialCode                     interface{}   `json:"materialCode"`
	ProjectNumber                    interface{}   `json:"projectNumber"`
	Usage                            interface{}   `json:"usage"`
	UsageClass                       interface{}   `json:"usageClass"`
	VendorSKU                        interface{}   `json:"vendorSKU"`
	CountryOfOriginISOCode           interface{}   `json:"countryOfOriginISOCode"`
	ModeOfTransport                  interface{}   `json:"modeOfTransport"`
	NatureOfTransaction              interface{}   `json:"natureOfTransaction"`
	IntrastatCommodityCode           interface{}   `json:"intrastatCommodityCode"`
	NetMassKilograms                 interface{}   `json:"netMassKilograms"`
	LineItemID                       interface{}   `json:"lineItemId"`
	TaxIncludedIndicator             interface{}   `json:"taxIncludedIndicator"`
	TransactionType                  interface{}   `json:"transactionType"`
	SimplificationCode               interface{}   `json:"simplificationCode"`
	TitleTransfer                    interface{}   `json:"titleTransfer"`
	ChainTransactionPhase            interface{}   `json:"chainTransactionPhase"`
	ExportProcedure                  interface{}   `json:"exportProcedure"`
	MaterialOrigin                   interface{}   `json:"materialOrigin"`
}

type Product struct {
	Value        string `json:"value"`
	ProductClass string `json:"productClass"`
}
type Quantity struct {
	Value         float64     `json:"value"`
	UnitOfMeasure interface{} `json:"unitOfMeasure"`
}

type Taxes struct {
	Jurisdiction                     `json:"jurisdiction"`
	AccumulateAsJurisdiction         interface{} `json:"accumulateAsJurisdiction"`
	CalculatedTax                    float64     `json:"calculatedTax"`
	EffectiveRate                    float64     `json:"effectiveRate"`
	TaxApportionmentRate             interface{} `json:"taxApportionmentRate"`
	BasisReductionPercentage         interface{} `json:"basisReductionPercentage"`
	Exempt                           interface{} `json:"exempt"`
	NonTaxable                       interface{} `json:"nonTaxable"`
	Taxable                          `json:"taxable"`
	ReportingBasis                   interface{} `json:"reportingBasis"`
	Imposition                       `json:"imposition"`
	ImpositionType                   `json:"impositionType"`
	AccumulateAsImposition           interface{} `json:"accumulateAsImposition"`
	AccumulateAsImpositionType       interface{} `json:"accumulateAsImpositionType"`
	TaxRuleID                        `json:"taxRuleId"`
	BasisRuleID                      interface{} `json:"basisRuleId"`
	InclusionRuleID                  interface{} `json:"inclusionRuleId"`
	MaxTaxRuleID                     interface{} `json:"maxTaxRuleId"`
	RecoverableRuleID                interface{} `json:"recoverableRuleId"`
	PostCalculationEvaluationRuleID  interface{} `json:"postCalculationEvaluationRuleId"`
	CreditRuleID                     interface{} `json:"creditRuleId"`
	BasisApportionmentRuleID         interface{} `json:"basisApportionmentRuleId"`
	TaxRateAdjustmentRuleID          interface{} `json:"taxRateAdjustmentRuleId"`
	TaxApportionmentRuleID           interface{} `json:"taxApportionmentRuleId"`
	AccumulationRuleID               interface{} `json:"accumulationRuleId"`
	TelecomUnitConversionRuleID      interface{} `json:"telecomUnitConversionRuleId"`
	ReportingBasisRuleID             interface{} `json:"reportingBasisRuleId"`
	CertificateNumber                interface{} `json:"certificateNumber"`
	RecoverableAmount                interface{} `json:"recoverableAmount"`
	RecoverablePercent               interface{} `json:"recoverablePercent"`
	BlockingRecoverablePercent       interface{} `json:"blockingRecoverablePercent"`
	PartialExemptRecoverablePercent  interface{} `json:"partialExemptRecoverablePercent"`
	UnrecoverableAmount              interface{} `json:"unrecoverableAmount"`
	FilingCurrencyAmounts            interface{} `json:"filingCurrencyAmounts"`
	SellerRegistrationID             string      `json:"sellerRegistrationId"`
	BuyerRegistrationID              string      `json:"buyerRegistrationId"`
	OwnerRegistrationID              interface{} `json:"ownerRegistrationId"`
	DispatcherRegistrationID         interface{} `json:"dispatcherRegistrationId"`
	RecipientRegistrationID          interface{} `json:"recipientRegistrationId"`
	InvoiceTextCode                  []int       `json:"invoiceTextCode"`
	SummaryInvoiceText               interface{} `json:"summaryInvoiceText"`
	InvoiceTexts                     interface{} `json:"invoiceTexts"`
	AssistedParameters               interface{} `json:"assistedParameters"`
	TaxRuleCurrencyConversionFactors interface{} `json:"taxRuleCurrencyConversionFactors"`
	OriginalTax                      interface{} `json:"originalTax"`
	IncludedTax                      interface{} `json:"includedTax"`
	NominalRate                      float64     `json:"nominalRate"`
	MarkUpRate                       interface{} `json:"markUpRate"`
	AccumulateAsLineType             interface{} `json:"accumulateAsLineType"`
	TaxResult                        string      `json:"taxResult"`
	TaxType                          string      `json:"taxType"`
	MaxTaxIndicator                  interface{} `json:"maxTaxIndicator"`
	Situs                            string      `json:"situs"`
	NotRegisteredIndicator           interface{} `json:"notRegisteredIndicator"`
	InputOutputType                  string      `json:"inputOutputType"`
	TaxCode                          interface{} `json:"taxCode"`
	VertexTaxCode                    interface{} `json:"vertexTaxCode"`
	ReasonCode                       interface{} `json:"reasonCode"`
	FilingCategoryCode               interface{} `json:"filingCategoryCode"`
	IsService                        interface{} `json:"isService"`
	RateClassification               string      `json:"rateClassification"`
	TaxCollectedFromParty            string      `json:"taxCollectedFromParty"`
	TaxStructure                     string      `json:"taxStructure"`
}

type Jurisdiction struct {
	Value                    string      `json:"value"`
	JurisdictionLevel        string      `json:"jurisdictionLevel"`
	JurisdictionID           int         `json:"jurisdictionId"`
	EffectiveDate            interface{} `json:"effectiveDate"`
	ExpirationDate           interface{} `json:"expirationDate"`
	ExternalJurisdictionCode interface{} `json:"externalJurisdictionCode"`
}

type Taxable struct {
	Value         float64     `json:"value"`
	UnitOfMeasure interface{} `json:"unitOfMeasure"`
}

type Imposition struct {
	Value        string      `json:"value"`
	UserDefined  interface{} `json:"userDefined"`
	ImpositionID int         `json:"impositionId"`
}

type ImpositionType struct {
	Value            string      `json:"value"`
	UserDefined      interface{} `json:"userDefined"`
	ImpositionTypeID int         `json:"impositionTypeId"`
	WithholdingType  interface{} `json:"withholdingType"`
}

type TaxRuleID struct {
	Value                    int         `json:"value"`
	UserDefined              interface{} `json:"userDefined"`
	SalesTaxHolidayIndicator interface{} `json:"salesTaxHolidayIndicator"`
	TaxRuleType              interface{} `json:"taxRuleType"`
}

type AssistedParameters struct {
	Param []AssistedParameter `json:"assistedParameter"`
}

type AssistedParameter struct {
	Value         string `json:"value"`
	ParamName     string `json:"paramName"`
	Phase         string `json:"phase"`
	RuleName      string `json:"ruleName"`
	OriginalValue string `json:"originalValue"`
}

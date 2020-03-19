package model

type TaxResponse struct {
	OutData OutData `json:"outdata"`
}

type OutData struct {
	CompanyId         int64               `json:"companyId"`
	CompanyName       string              `json:"companyName"`
	CompanyRole       string              `json:"companyRole"`
	Invoice           []ResponseInvoiceUP `json:"invoice"`
	RequestStatus     RequestStatus       `json:"requestStatus"`
	ExternalCompanyId string              `json:"externalCompanyId"`
	ScenarioName      string              `json:"scenarioName"`
	Version           string              `json:"version"` // need to verify

}
type ResponseInData struct {
	CallingSystemNumber string                `json:"callingSystemNumber"`
	CompanyId           int64                 `json:"companyId"`
	CompanyName         string                `json:"companyName"`
	CompanyRole         string                `json:"companyRole"`
	ExternalCompanyId   string                `json:"externalCompanyId"`
	HostInfo            HostRequestInfo       `json:"hostRequestInfo"`
	HostSystem          string                `json:"hostSystem"`
	Invoice             []ResponseInvoiceDown `json:"invoice"`
	Password            string                `json:"password,omitempty"`
	ScenarioId          int64                 `json:"scenarioId"`
	ScenarioName        string                `json:"scenarioName"`
	Username            string                `json:"username,omitempty"`
	Version             string                `json:"version"` // need to verify
	XmlGroupName        string                `json:"xmlGroupName,omitempty"`
	XmlGroupOwner       string                `json:"xmlGroupOwner,omitempty"`
}

type ResponseInvoiceUP struct {
	BasisPercent            string         `json:"basisPercent"`
	CalculationDirection    string         `json:"calculationDirection"`
	CallingSystemNumber     string         `json:"callingSystemNumber"`
	CompanyId               int64          `json:"companyId"`
	CompanyName             string         `json:"companyName"`
	CompanyRole             string         `json:"companyRole"`
	CurrencyCode            string         `json:"currencyCode"`
	CurrencyName            string         `json:"currencyName"`
	CustomerGroupName       string         `json:"customerGroupName"`
	CustomerGroupOwner      string         `json:"customerGroupOwner"`
	CustomerName            string         `json:"customerName"`
	CustomerNumber          string         `json:"customerNumber"`
	EndUserName             string         `json:"endUserName"`
	ExternalCompanyId       string         `json:"externalCompanyId"`
	FiscalDate              string         `json:"fiscalDate"`
	FunctionalCurrencyCode  string         `json:"functionalCurrencyCode"`
	HostSystem              string         `json:"hostSystem"`
	InData                  ResponseInData `json:"indata"`
	InvoiceDate             string         `json:"invoiceDate"`
	InvoiceNumber           string         `json:"invoiceNumber"`
	IsAuditUpdate           string         `json:"isAuditUpdate"`
	IsBusinessSupply        string         `json:"isBusinessSupply"`
	IsCredit                string         `json:"isCredit"`
	IsReported              string         `json:"isReported"`
	IsReversed              string         `json:"isReversed"`
	Line                    []LineResponse `json:"line"`
	Message                 []Message      `json:"message"`
	MinAccountableUnit      string         `json:"minAccountableUnit"`
	NatureOfTransactionCode string         `json:"natureOfTransactionCode"`
	OriginalDocumentId      string         `json:"originalDocumentId"`
	OriginalDocumentItem    string         `json:"originalDocumentItem"`
	OriginalDocumentType    string         `json:"originalDocumentType"`
	OriginalInvoiceDate     string         `json:"originalInvoiceDate"`
	OriginalInvoiceNumber   string         `json:"originalInvoiceNumber"`
	OriginalMovementDate    string         `json:"originalMovementDate"`
	RequestStatus           RequestStatus  `json:"requestStatus"`
	RoundingPrecision       int64          `json:"roundingPrecision"`
	RoundingRule            string         `json:"roundingRule"`
	StatisticalProcedure    string         `json:"statisticalProcedure"`
	TaxSummary              TaxSummary     `json:"taxSummary"`
	TotalTaxAmount          string         `json:"totalTaxAmount"`
	TransactionDate         string         `json:"transactionDate"`
	UniqueInvoiceNumber     string         `json:"uniqueInvoiceNumber"`
	UserElement             []UserElement  `json:"userElement"`
	VendorName              string         `json:"vendorName"`
	VendorNumber            string         `json:"vendorNumber"`
	VendorTax               string         `json:"vendorTax"`
}

type ResponseInvoiceDown struct {
	AllocationGroupName      string         `json:"allocationGroupName"`
	AllocationGroupOwner     string         `json:"allocationGroupOwner"`
	AllocationName           string         `json:"allocationName"`
	AutoCreateCertificates   string         `json:"autoCreateCertificates"`
	AutoCreateCustomers      string         `json:"autoCreateCustomers"`
	BillTo                   Address        `json:"billTo"`
	BuyerPrimary             Address        `json:"buyerPrimary"`
	CalculationDirection     string         `json:"calculationDirection"`
	CallingSystemNumber      string         `json:"callingSystemNumber"`
	CompanyId                int64          `json:"companyId"`
	CompanyName              string         `json:"companyName"`
	CompanyRole              string         `json:"companyRole"`
	CountryOfOrigin          string         `json:"countryOfOrigin"`
	CurrencyCode             string         `json:"currencyCode"`
	CustomerGroupName        string         `json:"customerGroupName"`
	CustomerGroupOwner       string         `json:"customerGroupOwner"`
	CustomerName             string         `json:"customerName"`
	CustomerNumber           string         `json:"customerNumber"`
	DeliveryTerms            string         `json:"deliveryTerms"`
	DeptOfConsign            string         `json:"deptOfConsign"`
	DocumentType             string         `json:"documentType"`
	EndUse                   []string       `json:"endUse"`
	EndUserName              string         `json:"endUserName"`
	Establishments           Establishment  `json:"establishments"`
	ExemptAmount             Amount         `json:"exemptAmount"`
	ExemptCertificate        Amount         `json:"exemptCertificate"`
	ExemptReason             Amount         `json:"exemptReason"`
	ExternalCompanyId        string         `json:"externalCompanyId"`
	FilterGroupName          string         `json:"filterGroupName"`
	FilterGroupOwner         string         `json:"filterGroupOwner"`
	FiscalDate               string         `json:"fiscalDate"`
	FunctionalCurrencyCode   string         `json:"functionalCurrencyCode"`
	HostSystem               string         `json:"hostSystem"`
	InclusiveTaxIndicators   TaxIndicator   `json:"inclusiveTaxIndicators"`
	InputRecoveryType        string         `json:"inputRecoveryType"`
	InvoiceDate              string         `json:"invoiceDate"`
	InvoiceNumber            string         `json:"invoiceNumber"`
	IsAuditUpdate            string         `json:"isAuditUpdate"`
	IsAudited                string         `json:"isAudited"`
	IsBusinessSupply         string         `json:"isBusinessSupply"`
	IsCredit                 string         `json:"isCredit"`
	IsExempt                 Example        `json:"isExempt"`
	IsNoTax                  Example        `json:"isNoTax"`
	IsReported               string         `json:"isReported"`
	IsReversed               string         `json:"isReversed"`
	IsRounding               string         `json:"isRounding"`
	IsSimplification         string         `json:"isSimplification"`
	Licenses                 License        `json:"licenses"`
	Line                     []LineResponse `json:"line"`
	Location                 BillingRole    `json:"location"`
	LocationSet              string         `json:"locationSet"`
	Middleman                Address        `json:"middleman"`
	ModeOfTransport          string         `json:"modeOfTransport"`
	MovementDate             string         `json:"movementDate"`
	MovementType             string         `json:"movementType"`
	NatureOfTransactionCode  string         `json:"natureOfTransactionCode"`
	OrderAcceptance          Address        `json:"orderAcceptance"`
	OrderOrigin              Address        `json:"orderOrigin"`
	OriginalDocumentId       string         `json:"originalDocumentId"`
	OriginalDocumentItem     string         `json:"originalDocumentItem"`
	OriginalDocumentType     string         `json:"originalDocumentType"`
	OriginalInvoiceDate      string         `json:"originalInvoiceDate"`
	OriginalInvoiceNumber    string         `json:"originalInvoiceNumber"`
	OriginalMovementDate     string         `json:"originalMovementDate"`
	OverrideAmount           Amount         `json:"overrideAmount"`
	OverrideRate             Amount         `json:"overrideRate"`
	PointOfTitleTransfer     string         `json:"pointOfTitleTransfer"`
	PortOfEntry              string         `json:"portOfEntry"`
	PortOfLoading            string         `json:"portOfLoading"`
	ProductMappingGroupName  string         `json:"productMappingGroupName"`
	ProductMappingGroupOwner string         `json:"productMappingGroupOwner"`
	Regime                   string         `json:"regime"`
	Registrations            Registrations  `json:"registrations"`
	SellerPrimary            Address        `json:"sellerPrimary"`
	ShipFrom                 Address        `json:"shipFrom"`
	ShipTo                   Address        `json:"shipTo"`
	StatisticalProcedure     string         `json:"statisticalProcedure"`
	Supply                   Address        `json:"supply"`
	SupplyExemptPercent      Amount         `json:"supplyExemptPercent"`
	TaxCode                  string         `json:"taxCode"`
	TaxDeterminationDate     string         `json:"taxDeterminationDate"`
	TaxExchangeRateDate      string         `json:"taxExchangeRateDate"`
	TaxPointDate             string         `json:"taxPointDate"`
	TaxTreatment             string         `json:"taxTreatment"`
	TaxType                  Example        `json:"taxType"`
	TitleTransferLocation    string         `json:"titleTransferLocation"`
	TransactionType          string         `json:"transactionType"`
	UniqueInvoiceNumber      string         `json:"uniqueInvoiceNumber"`
	UserElement              []UserElement  `json:"userElement"`
	VatGroupRegistration     string         `json:"vatGroupRegistration"`
	VendorName               string         `json:"vendorName"`
	VendorNumber             string         `json:"vendorNumber"`
	VendorTax                string         `json:"vendorTax"`
}

type RequestStatus struct {
	Error            []Err `json:"error"`
	IsPartialSuccess bool  `json:"isPartialSuccess"`
	IsSuccess        bool  `json:"isSuccess"`
}

type Err struct {
	Code              string `json:"code"`
	Description       string `json:"description"`
	ErrorLocationPath string `json:"errorLocationPath"`
}

type Message struct {
	Category    string `json:"category"`
	Code        string `json:"code"`
	Location    string `json:"location"`
	MessageText string `json:"messageText"`
	Severity    int64  `json:"severity"`
}

type TaxSummary struct {
	Advisories       Advisories `json:"advisories"`
	EffectiveTaxRate string     `json:"effectiveTaxRate"`
	ExemptAmount     string     `json:"exemptAmount"`
	NonTaxableBasis  string     `json:"nonTaxableBasis"`
	TaxRate          string     `json:"taxRate"`
	TaxableBasis     string     `json:"taxableBasis"`
}

type Advisories struct {
	Advisory []string `json:"advisory"`
}

type LineResponse struct {
	AccountingCode        string        `json:"accountingCode"`
	AllocationLine        []interface{} `json:"allocationLine"`
	AllocationName        string        `json:"allocationName"`
	BasisPercent          string        `json:"basisPercent"`
	BillToBranchId        string        `json:"billToBranchId"`
	CommodityCode         string        `json:"commodityCode"`
	CountryOfOrigin       string        `json:"countryOfOrigin"`
	CustomerGroupName     string        `json:"customerGroupName"`
	CustomerGroupOwner    string        `json:"customerGroupOwner"`
	CustomerName          string        `json:"customerName"`
	CustomerNumber        string        `json:"customerNumber"`
	DeliveryTerms         string        `json:"deliveryTerms"`
	DeptOfConsign         string        `json:"deptOfConsign"`
	Description           string        `json:"description"`
	DiscountAmount        string        `json:"discountAmount"`
	EndUserName           string        `json:"endUserName"`
	FreightOnBoard        string        `json:"freightOnBoard"`
	GrossAmount           string        `json:"grossAmount"`
	Id                    string        `json:"id"`
	InputRecoveryAmount   string        `json:"inputRecoveryAmount"`
	InputRecoveryPercent  string        `json:"inputRecoveryPercent"`
	InvoiceDate           string        `json:"invoiceDate"`
	IsBusinessSupply      string        `json:"isBusinessSupply"`
	IsCredit              string        `json:"isCredit"`
	ItemValue             string        `json:"itemValue"`
	LineNumber            int64         `json:"lineNumber"`
	Mass                  int64         `json:"mass"`
	Message               []Message     `json:"message"`
	MiddlemanBranchId     string        `json:"middlemanBranchId"`
	MiddlemanMarkupAmount int64         `json:"middlemanMarkupAmount"`
	MiddlemanMarkupRate   int64         `json:"middlemanMarkupRate"`
	ModeOfTransport       string        `json:"modeOfTransport"`
	MovementDate          string        `json:"movementDate"`
	OriginalDocumentId    string        `json:"originalDocumentId"`
	OriginalDocumentItem  string        `json:"originalDocumentItem"`
	OriginalDocumentType  string        `json:"originalDocumentType"`
	OriginalInvoiceDate   string        `json:"originalInvoiceDate"`
	OriginalMovementDate  string        `json:"originalMovementDate"`
	PartNumber            string        `json:"partNumber"`
	PointOfTitleTransfer  string        `json:"pointOfTitleTransfer"`
	PortOfEntry           string        `json:"portOfEntry"`
	PortOfLoading         string        `json:"portOfLoading"`
	Quantities            Quantities    `json:"quantities"`
	Regime                string        `json:"regime"`
	RelatedLineNumber     int64         `json:"relatedLineNumber"`
	ShipFromBranchId      string        `json:"shipFromBranchId"`
	ShipFromCountry       string        `json:"shipFromCountry"`
	ShipToBranchId        string        `json:"shipToBranchId"`
	ShipToCountry         string        `json:"shipToCountry"`
	SupplementaryUnit     string        `json:"supplementaryUnit"`
	SupplyBranchId        string        `json:"supplyBranchId"`
	Tax                   []Tax         `json:"tax"` //
	TaxCode               string        `json:"taxCode"`
	TaxSummary            TaxSummary    `json:"taxSummary"`
	TitleTransferLocation string        `json:"titleTransferLocation"`
	TotalTaxAmount        string        `json:"totalTaxAmount"`
	TransactionType       string        `json:"transactionType"`
	UniqueLineNumber      string        `json:"uniqueLineNumber"`
	UnitOfMeasure         string        `json:"unitOfMeasure"`
	UserElement           []UserElement `json:"userElement"`
	VendorName            string        `json:"vendorName"`
	VendorNumber          string        `json:"vendorNumber"`
	VendorTax             string        `json:"vendorTax"`
}

type Tax struct {
	AddressType                 string               `json:"addressType"`
	AdminZoneLevel              string               `json:"adminZoneLevel"`
	AuthorityAttribute          []UserElement        `json:"authorityAttribute"`
	AuthorityCategory           string               `json:"authorityCategory"`
	AuthorityCurrencyCode       string               `json:"authorityCurrencyCode"`
	AuthorityFips               string               `json:"authorityFips"`
	AuthorityName               string               `json:"authorityName"`
	AuthorityOfficialName       string               `json:"authorityOfficialName"`
	AuthorityType               string               `json:"authorityType"`
	AuthorityTypeAlias          string               `json:"authorityTypeAlias"`
	AuthorityUuid               string               `json:"authorityUuid"`
	BasisPercent                string               `json:"basisPercent"`
	BuyerRegistration           string               `json:"buyerRegistration"`
	CalculationMethod           string               `json:"calculationMethod"`
	Comment                     string               `json:"comment"`
	CurrencyConversion          []CurrencyConversion `json:"currencyConversion"`
	EffectiveZoneLevel          string               `json:"effectiveZoneLevel"`
	ErpTaxCode                  string               `json:"erpTaxCode"`
	EuTransaction               bool                 `json:"euTransaction"`
	ExemptAmount                ExemptAmount         `json:"exemptAmount"`
	ExemptCertificate           string               `json:"exemptCertificate"`
	ExemptCertificateExpireDate string               `json:"exemptCertificateExpireDate"`
	ExemptReason                string               `json:"exemptReason"`
	Fee                         []Fee                `json:"fee"`
	FiscalRepAddress1           string               `json:"fiscalRepAddress1"`
	FiscalRepAddress2           string               `json:"fiscalRepAddress2"`
	FiscalRepContact            string               `json:"fiscalRepContact"`
	FiscalRepName               string               `json:"fiscalRepName"`
	GrossAmount                 ExemptAmount         `json:"grossAmount"`
	InclusiveTax                string               `json:"inclusiveTax"`
	InputRecoveryAmount         string               `json:"inputRecoveryAmount"`
	InputRecoveryPercent        string               `json:"inputRecoveryPercent"`
	InvoiceDescription          string               `json:"invoiceDescription"`
	IsExempt                    bool                 `json:"isExempt"`
	IsIntrastatReported         string               `json:"isIntrastatReported"`
	IsNotax                     bool                 `json:"isNotax"`
	IsTriangulation             bool                 `json:"isTriangulation"`
	IsVatReported               bool                 `json:"isVatReported"`
	IsViesReported              bool                 `json:"isViesReported"`
	JurisdictionText            string               `json:"jurisdictionText"`
	Licenses                    License              `json:"licenses"`
	LocationCode                string               `json:"locationCode"`
	Message                     []Message            `json:"message"`
	MiddlemanRegistration       string               `json:"middlemanRegistration"`
	NatureOfTax                 string               `json:"natureOfTax"`
	NonTaxableBasis             ExemptAmount         `json:"nonTaxableBasis"`
	OverrideAmount              int64                `json:"overrideAmount"`
	OverrideRate                int64                `json:"overrideRate"`
	RegistrationAttribute1      string               `json:"registrationAttribute1"`
	RegistrationAttribute10     string               `json:"registrationAttribute10"`
	RegistrationAttribute11     string               `json:"registrationAttribute11"`
	RegistrationAttribute12     string               `json:"registrationAttribute12"`
	RegistrationAttribute13     string               `json:"registrationAttribute13"`
	RegistrationAttribute14     string               `json:"registrationAttribute14"`
	RegistrationAttribute15     string               `json:"registrationAttribute15"`
	RegistrationAttribute16     string               `json:"registrationAttribute16"`
	RegistrationAttribute17     string               `json:"registrationAttribute17"`
	RegistrationAttribute18     string               `json:"registrationAttribute18"`
	RegistrationAttribute19     string               `json:"registrationAttribute19"`
	RegistrationAttribute2      string               `json:"registrationAttribute2"`
	RegistrationAttribute20     string               `json:"registrationAttribute20"`
	RegistrationAttribute21     string               `json:"registrationAttribute21"`
	RegistrationAttribute22     string               `json:"registrationAttribute22"`
	RegistrationAttribute23     string               `json:"registrationAttribute23"`
	RegistrationAttribute24     string               `json:"registrationAttribute24"`
	RegistrationAttribute25     string               `json:"registrationAttribute25"`
	RegistrationAttribute26     string               `json:"registrationAttribute26"`
	RegistrationAttribute27     string               `json:"registrationAttribute27"`
	RegistrationAttribute28     string               `json:"registrationAttribute28"`
	RegistrationAttribute29     string               `json:"registrationAttribute29"`
	RegistrationAttribute3      string               `json:"registrationAttribute3"`
	RegistrationAttribute30     string               `json:"registrationAttribute30"`
	RegistrationAttribute31     string               `json:"registrationAttribute31"`
	RegistrationAttribute32     string               `json:"registrationAttribute32"`
	RegistrationAttribute33     string               `json:"registrationAttribute33"`
	RegistrationAttribute34     string               `json:"registrationAttribute34"`
	RegistrationAttribute35     string               `json:"registrationAttribute35"`
	RegistrationAttribute36     string               `json:"registrationAttribute36"`
	RegistrationAttribute37     string               `json:"registrationAttribute37"`
	RegistrationAttribute38     string               `json:"registrationAttribute38"`
	RegistrationAttribute39     string               `json:"registrationAttribute39"`
	RegistrationAttribute4      string               `json:"registrationAttribute4"`
	RegistrationAttribute40     string               `json:"registrationAttribute40"`
	RegistrationAttribute41     string               `json:"registrationAttribute41"`
	RegistrationAttribute42     string               `json:"registrationAttribute42"`
	RegistrationAttribute43     string               `json:"registrationAttribute43"`
	RegistrationAttribute44     string               `json:"registrationAttribute44"`
	RegistrationAttribute45     string               `json:"registrationAttribute45"`
	RegistrationAttribute46     string               `json:"registrationAttribute46"`
	RegistrationAttribute47     string               `json:"registrationAttribute47"`
	RegistrationAttribute48     string               `json:"registrationAttribute48"`
	RegistrationAttribute49     string               `json:"registrationAttribute49"`
	RegistrationAttribute5      string               `json:"registrationAttribute5"`
	RegistrationAttribute50     string               `json:"registrationAttribute50"`
	RegistrationAttribute6      string               `json:"registrationAttribute6"`
	RegistrationAttribute7      string               `json:"registrationAttribute7"`
	RegistrationAttribute8      string               `json:"registrationAttribute8"`
	RegistrationAttribute9      string               `json:"registrationAttribute9"`
	RelatedAllocationLineNumber int64                `json:"relatedAllocationLineNumber"`
	RelatedLineNumber           int64                `json:"relatedLineNumber"`
	RevisedGrossAmount          int64                `json:"revisedGrossAmount"`
	RuleOrder                   int64                `json:"ruleOrder"`
	RuleReportingCategory       string               `json:"ruleReportingCategory"`
	SellerRegistration          string               `json:"sellerRegistration"`
	SupplyExemptPercent         int64                `json:"supplyExemptPercent"`
	TaxAmount                   ExemptAmount         `json:"taxAmount"`
	TaxDeterminationDate        string               `json:"taxDeterminationDate"`
	TaxDirection                string               `json:"taxDirection"`
	TaxPointDate                string               `json:"taxPointDate"`
	TaxRate                     string               `json:"taxRate"`
	TaxRateCode                 string               `json:"taxRateCode"`
	TaxTreatment                string               `json:"taxTreatment"`
	TaxType                     string               `json:"taxType"`
	TaxableBasis                ExemptAmount         `json:"taxableBasis"`
	TaxableCity                 string               `json:"taxableCity"`
	TaxableCountry              string               `json:"taxableCountry"`
	TaxableCountryName          string               `json:"taxableCountryName"`
	TaxableCounty               string               `json:"taxableCounty"`
	TaxableDistrict             string               `json:"taxableDistrict"`
	TaxableGeocode              string               `json:"taxableGeocode"`
	TaxablePostcode             string               `json:"taxablePostcode"`
	TaxableProvince             string               `json:"taxableProvince"`
	TaxableState                string               `json:"taxableState"`
	UomConversion               UomConversion        `json:"uomConversion"`
	VatGroupRegistration        string               `json:"vatGroupRegistration"`
	ZoneLevel                   string               `json:"zoneLevel"`
	ZoneName                    string               `json:"zoneName"`
}
type CurrencyConversion struct {
	ConversionStep      []ConversionStep `json:"conversionSteps"`
	ExchangeRateSource  string           `json:"exchangeRateSource"`
	TaxExchangeRateDate string           `json:"taxExchangeRateDate"`
}

type ConversionStep struct {
	ConversionStep   int64  `json:"conversionStep"`
	ExchangeRate     int64  `json:"exchangeRate"`
	FromCurrencyCode string `json:"fromCurrencyCode"`
	ToCurrencyCode   string `json:"toCurrencyCode"`
}

type ExemptAmount struct {
	AuthorityAmount          string `json:"authorityAmount"`
	DocumentAmount           string `json:"documentAmount"`
	UnroundedAuthorityAmount string `json:"unroundedAuthorityAmount"`
	UnroundedDocumentAmount  string `json:"unroundedDocumentAmount"`
}

type Fee struct {
	Amount        int64  `json:"amount"`
	CurrencyCode  string `json:"currencyCode"`
	UnitOfMeasure string `json:"unitOfMeasure"`
}

type License struct {
	LicenseCategory           string `json:"licenseCategory"`
	LicenseEndDate            string `json:"licenseEndDate"`
	LicenseExternalIdentifier string `json:"licenseExternalIdentifier"`
	LicenseNumber             string `json:"licenseNumber"`
	LicenseTypeName           string `json:"licenseTypeName"`
}

type UomConversion struct {
	Factor      string `json:"factor"`
	From        UOM    `json:"from"`
	Operator    string `json:"operator"`
	ToRounded   UOM    `json:"toRounded"`
	ToUnrounded UOM    `json:"toUnrounded"`
}

type UOM struct {
	Amount  string `json:"amount"`
	Default string `json:"default"`
	Uom     string `json:"uom"`
}

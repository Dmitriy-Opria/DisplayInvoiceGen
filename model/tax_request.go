package model

type TaxRequest struct {
	In RequestInData `json:"indata"`
}
type RequestInData struct {
	CallingSystemNumber string           `json:"callingSystemNumber"`
	CompanyId           int64            `json:"companyId"`
	CompanyName         string           `json:"companyName"`
	CompanyRole         string           `json:"companyRole"`
	ExternalCompanyId   string           `json:"externalCompanyId"`
	HostInfo            HostRequestInfo  `json:"hostRequestInfo"`
	HostSystem          string           `json:"hostSystem"`
	Invoice             []RequestInvoice `json:"invoice"`
	ScenarioId          int64            `json:"scenarioId"`
	ScenarioName        string           `json:"scenarioName"`
	Version             string           `json:"version"` // need to verify
}

type HostRequestInfo struct {
	HostRequestId         string `json:"hostRequestId,omitempty"`
	HostRequestLogEntryId string `json:"hostRequestLogEntryId"`
}

type RequestInvoice struct {
	AllocationGroupName      string          `json:"allocationGroupName"`
	AllocationGroupOwner     string          `json:"allocationGroupOwner"`
	AllocationName           string          `json:"allocationName"`
	AutoCreateCertificates   string          `json:"autoCreateCertificates"`
	AutoCreateCustomers      string          `json:"autoCreateCustomers"`
	BasisPercent             string          `json:"basisPercent"`
	BillTo                   Address         `json:"billTo"`
	BuyerPrimary             Address         `json:"buyerPrimary"`
	CalculationDirection     string          `json:"calculationDirection"`
	CallingSystemNumber      string          `json:"callingSystemNumber"`
	CompanyId                int64           `json:"companyId"`
	CompanyName              string          `json:"companyName"`
	CompanyRole              string          `json:"companyRole"`
	CountryOfOrigin          string          `json:"countryOfOrigin"`
	CurrencyCode             string          `json:"currencyCode"`
	CustomerGroupName        string          `json:"customerGroupName"`
	CustomerGroupOwner       string          `json:"customerGroupOwner"`
	CustomerName             string          `json:"customerName"`
	CustomerNumber           string          `json:"customerNumber"`
	DeliveryTerms            string          `json:"deliveryTerms"`
	DeptOfConsign            string          `json:"deptOfConsign"`
	DocumentType             string          `json:"documentType"`
	EndUse                   []string        `json:"endUse"`
	EndUserName              string          `json:"endUserName"`
	Establishments           Establishment   `json:"establishments"`
	ExemptAmount             Amount          `json:"exemptAmount"`
	ExemptCertificate        Amount          `json:"exemptCertificate"`
	ExemptReason             Amount          `json:"exemptReason"`
	ExternalCompanyId        string          `json:"externalCompanyId"`
	FilterGroupName          string          `json:"filterGroupName"`
	FilterGroupOwner         string          `json:"filterGroupOwner"`
	FiscalDate               string          `json:"fiscalDate"`
	FunctionalCurrencyCode   string          `json:"functionalCurrencyCode"`
	HostSystem               string          `json:"hostSystem"`
	InclusiveTaxIndicators   TaxIndicator    `json:"inclusiveTaxIndicators"`
	InputRecoveryType        string          `json:"inputRecoveryType"`
	InvoiceDate              string          `json:"invoiceDate"`
	InvoiceNumber            string          `json:"invoiceNumber"`
	IsAuditUpdate            string          `json:"isAuditUpdate"`
	IsAudited                string          `json:"isAudited"`
	IsBusinessSupply         string          `json:"isBusinessSupply"`
	IsCredit                 string          `json:"isCredit"`
	IsExempt                 Example         `json:"isExempt"`
	IsNoTax                  Example         `json:"isNoTax"`
	IsReported               string          `json:"isReported"`
	IsReversed               string          `json:"isReversed"`
	IsRounding               string          `json:"isRounding"`
	IsSimplification         string          `json:"isSimplification"`
	Licenses                 CustomerLicense `json:"licenses"`
	Line                     []LineRequest   `json:"line"`
	Location                 BillingRole     `json:"location"`
	LocationSet              string          `json:"locationSet"`
	Middleman                Address         `json:"middleman"`
	ModeOfTransport          string          `json:"modeOfTransport"`
	MovementDate             string          `json:"movementDate"`
	MovementType             string          `json:"movementType"`
	NatureOfTransactionCode  string          `json:"natureOfTransactionCode"`
	OrderAcceptance          Address         `json:"orderAcceptance"`
	OrderOrigin              Address         `json:"orderOrigin"`
	OriginalDocumentId       string          `json:"originalDocumentId"`
	OriginalDocumentItem     string          `json:"originalDocumentItem"`
	OriginalDocumentType     string          `json:"originalDocumentType"`
	OriginalInvoiceDate      string          `json:"originalInvoiceDate"`
	OriginalInvoiceNumber    string          `json:"originalInvoiceNumber"`
	OriginalMovementDate     string          `json:"originalMovementDate"`
	OverrideAmount           Amount          `json:"overrideAmount"`
	OverrideRate             Amount          `json:"overrideRate"`
	PointOfTitleTransfer     string          `json:"pointOfTitleTransfer"`
	PortOfEntry              string          `json:"portOfEntry"`
	PortOfLoading            string          `json:"portOfLoading"`
	ProductMappingGroupName  string          `json:"productMappingGroupName"`
	ProductMappingGroupOwner string          `json:"productMappingGroupOwner"`
	Regime                   string          `json:"regime"`
	Registrations            Registrations   `json:"registrations"`
	SellerPrimary            Address         `json:"sellerPrimary"`
	ShipFrom                 Address         `json:"shipFrom"`
	ShipTo                   Address         `json:"shipTo"`
	StatisticalProcedure     string          `json:"statisticalProcedure"`
	Supply                   Address         `json:"supply"`
	SupplyExemptPercent      Amount          `json:"supplyExemptPercent"`
	TaxCode                  string          `json:"taxCode"`
	TaxDeterminationDate     string          `json:"taxDeterminationDate"`
	TaxExchangeRateDate      string          `json:"taxExchangeRateDate"`
	TaxPointDate             string          `json:"taxPointDate"`
	TaxTreatment             string          `json:"taxTreatment"`
	TaxType                  Example         `json:"taxType"`
	TitleTransferLocation    string          `json:"titleTransferLocation"`
	TransactionType          string          `json:"transactionType"`
	UniqueInvoiceNumber      string          `json:"uniqueInvoiceNumber"`
	UserElement              []UserElement   `json:"userElement"`
	VatGroupRegistration     string          `json:"vatGroupRegistration"`
	VendorName               string          `json:"vendorName"`
	VendorNumber             string          `json:"vendorNumber"`
	VendorTax                string          `json:"vendorTax"`
}

type Address struct {
	Address1                     string `json:"address1"`
	Address2                     string `json:"address2"`
	Address3                     string `json:"address3"`
	AddressValidationMode        string `json:"addressValidationMode"`
	City                         string `json:"city"`
	CompanyBranchId              string `json:"companyBranchId"`
	Country                      string `json:"country"`
	County                       string `json:"county"`
	DefaultAddressValidationMode string `json:"defaultAddressValidationMode"`
	District                     string `json:"district"`
	Geocode                      string `json:"geocode"`
	IsBonded                     string `json:"isBonded"`
	LocationTaxCategory          string `json:"locationTaxCategory"`
	Postcode                     string `json:"postcode"`
	Province                     string `json:"province"`
	State                        string `json:"state"`
}

type Amount struct {
	City     string `json:"city"`
	Country  string `json:"country"`
	County   string `json:"county"`
	District string `json:"district"`
	Geocode  string `json:"geocode"`
	Postcode string `json:"postcode"`
	Province string `json:"province"`
	State    string `json:"state"`
}

type BillingRole struct {
	BillTo          string `json:"billTo"`
	BuyerPrimary    string `json:"buyerPrimary"`
	Middleman       string `json:"middleman"`
	OrderAcceptance string `json:"orderAcceptance"`
	OrderOrigin     string `json:"orderOrigin"`
	SellerPrimary   string `json:"sellerPrimary"`
	ShipFrom        string `json:"shipFrom"`
	ShipTo          string `json:"shipTo"`
	Supply          string `json:"supply"`
}

type Establishment struct {
	BuyerRole     BillingRole `json:"buyerRole"`
	MiddlemanRole BillingRole `json:"middlemanRole"`
	SellerRole    BillingRole `json:"sellerRole"`
}

type Registrations struct {
	BuyerRole     []string `json:"buyerRole"`
	MiddlemanRole []string `json:"middlemanRole"`
	SellerRole    []string `json:"sellerRole"`
}

type UserElement struct {
	Name  string `json:"name"`
	Value string `json:"value"`
}

type CustomerLicense struct {
	CustomerLicense []LicenseObject `json:"customerLicense"`
}

type LicenseObject struct {
	Number      string `json:"number"`
	LicenseType string `json:"type"`
}

type Example struct {
	All      string `json:"all"`
	City     string `json:"city"`
	Country  string `json:"country"`
	County   string `json:"county"`
	District string `json:"district"`
	Geocode  string `json:"geocode"`
	Postcode string `json:"postcode"`
	Province string `json:"province"`
	State    string `json:"state"`
}

type TaxIndicator struct {
	AuthorityType  []string `json:"authorityType"`
	FullyInclusive string   `json:"fullyInclusive"`
}

type LineRequest struct {
	AccountingCode         string          `json:"accountingCode"`
	AllocationGroupName    string          `json:"allocationGroupName"`
	AllocationGroupOwner   string          `json:"allocationGroupOwner"`
	AllocationName         string          `json:"allocationName"`
	BasisPercent           string          `json:"basisPercent"`
	BillTo                 Address         `json:"billTo"`
	BuyerPrimary           Address         `json:"buyerPrimary"`
	CommodityCode          string          `json:"commodityCode"`
	CountryOfOrigin        string          `json:"countryOfOrigin"`
	CustomerGroupName      string          `json:"customerGroupName"`
	CustomerGroupOwner     string          `json:"customerGroupOwner"`
	CustomerName           string          `json:"customerName"`
	CustomerNumber         string          `json:"customerNumber"`
	DeliveryTerms          string          `json:"deliveryTerms"`
	DeptOfConsign          string          `json:"deptOfConsign"`
	Description            string          `json:"description"`
	DiscountAmount         string          `json:"discountAmount"`
	EndUse                 []string        `json:"endUse"`
	EndUserName            string          `json:"endUserName"`
	Establishments         Establishment   `json:"establishments"`
	ExemptAmount           Amount          `json:"exemptAmount"`
	ExemptCertificate      Amount          `json:"exemptCertificate"`
	ExemptReason           Amount          `json:"exemptReason"`
	FreightOnBoard         string          `json:"freightOnBoard"`
	GrossAmount            string          `json:"grossAmount"`
	GrossPlusTax           string          `json:"grossPlusTax"`
	Id                     string          `json:"id"`
	InclusiveTaxIndicators TaxIndicator    `json:"inclusiveTaxIndicators"`
	InputRecoveryAmount    string          `json:"inputRecoveryAmount"`
	InputRecoveryPercent   string          `json:"inputRecoveryPercent"`
	InputRecoveryType      string          `json:"inputRecoveryType"`
	InvoiceDate            string          `json:"invoiceDate"`
	IsAllocatable          string          `json:"isAllocatable"`
	IsBusinessSupply       string          `json:"isBusinessSupply"`
	IsCredit               string          `json:"isCredit"`
	IsExempt               Example         `json:"isExempt"`
	IsManufacturing        string          `json:"isManufacturing"`
	IsNoTax                Example         `json:"isNoTax"`
	IsSimplification       string          `json:"isSimplification"`
	ItemValue              string          `json:"itemValue"`
	Licenses               CustomerLicense `json:"licenses"`
	LineNumber             int64           `json:"lineNumber"`
	Location               BillingRole     `json:"location"`
	LocationSet            string          `json:"locationSet"`
	Mass                   int64           `json:"mass"`
	Middleman              Address         `json:"middleman"`
	ModeOfTransport        string          `json:"modeOfTransport"`
	MovementDate           string          `json:"movementDate"`
	MovementType           string          `json:"movementType"`
	OrderAcceptance        Address         `json:"orderAcceptance"`
	OrderOrigin            Address         `json:"orderOrigin"`
	OriginalDocumentId     string          `json:"originalDocumentId"`
	OriginalDocumentItem   string          `json:"originalDocumentItem"`
	OriginalDocumentType   string          `json:"originalDocumentType"`
	OriginalInvoiceDate    string          `json:"originalInvoiceDate"`
	OriginalMovementDate   string          `json:"originalMovementDate"`
	OverrideAmount         Amount          `json:"overrideAmount"`
	OverrideRate           Amount          `json:"overrideRate"`
	PartNumber             string          `json:"partNumber"`
	PointOfTitleTransfer   string          `json:"pointOfTitleTransfer"`
	PortOfEntry            string          `json:"portOfEntry"`
	PortOfLoading          string          `json:"portOfLoading"`
	ProductCode            string          `json:"productCode"`
	Quantities             Quantities      `json:"quantities"`
	Quantity               int64           `json:"quantity"`
	Regime                 string          `json:"regime"`
	Registrations          Registrations   `json:"registrations"`
	RelatedLineNumber      int64           `json:"relatedLineNumber"`
	SellerPrimary          Address         `json:"sellerPrimary"`
	ShipFrom               Address         `json:"shipFrom"`
	ShipTo                 Address         `json:"shipTo"`
	SupplementaryUnit      string          `json:"supplementaryUnit"`
	Supply                 Address         `json:"supply"`
	SupplyExemptPercent    Amount          `json:"supplyExemptPercent"`
	TaxAmount              string          `json:"taxAmount"`
	TaxCode                string          `json:"taxCode"`
	TaxDeterminationDate   string          `json:"taxDeterminationDate"`
	TaxExchangeRateDate    string          `json:"taxExchangeRateDate"`
	TaxPointDate           string          `json:"taxPointDate"`
	TaxTreatment           string          `json:"taxTreatment"`
	TaxType                Example         `json:"taxType"`
	TitleTransferLocation  string          `json:"titleTransferLocation"`
	TransactionType        string          `json:"transactionType"`
	UniqueLineNumber       string          `json:"uniqueLineNumber"`
	UnitOfMeasure          string          `json:"unitOfMeasure"`
	Uom                    string          `json:"uom"`
	UserElement            []UserElement   `json:"userElement"`
	VatGroupRegistration   string          `json:"vatGroupRegistration"`
	VendorName             string          `json:"vendorName"`
	VendorNumber           string          `json:"vendorNumber"`
	VendorTax              string          `json:"vendorTax"`
}

type Quantities struct {
	Quantity []QuantityItem `json:"quantity"`
}

type QuantityItem struct {
	Amount string `json:"amount"`
	Def    string `json:"default"`
	Uom    string `json:"uom"`
}

package model

type Invoice struct {
	InvoiceRequest InvoiceRequest `json:"invoiceRequest"`
}

type InvoiceRequest struct {
	Customer        Customer          `json:"customer"`
	Seller          Seller            `json:"seller"`
	Lines           []LineItemInvoice `json:"lineItem"`
	DocumentDate    string            `json:"documentDate"`
	TransactionType string            `json:"transactionType"`
}

type Customer struct {
	Destination Destination              `json:"destination"`
	Tax         []TaxRegistrationInvoice `json:"taxRegistration"`
}

type Seller struct {
	Company         string                   `json:"company"`
	Division        string                   `json:"division"`
	PhysicalOrigin  Destination              `json:"physicalOrigin"`
	TaxRegistration []TaxRegistrationInvoice `json:"taxRegistration"`
}

type LineItemInvoice struct {
	Product        ProductInvoice `json:"product"`
	ExtendedPrice  string         `json:"extendedPrice"`
	LineItemNumber string         `json:"lineItemNumber"`
}

type Destination struct {
	StreetAddress1 string `json:"streetAddress1"`
	StreetAddress2 string `json:"streetAddress2"`
	City           string `json:"city"`
	MainDivision   string `json:"mainDivision"`
	SubDivision    string `json:"subDivision"`
	PostalCode     string `json:"postalCode"`
	Country        string `json:"country"`
}

type TaxRegistrationInvoice struct {
	TaxRegistrationNumber        string `json:"taxRegistrationNumber"`
	HasPhysicalPresenceIndicator string `json:"hasPhysicalPresenceIndicator"`
	IsoCountryCode               string `json:"isoCountryCode"`
}

type ProductInvoice struct {
	ProductClass string `json:"productClass"`
}

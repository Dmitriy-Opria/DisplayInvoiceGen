package model

type TaxResponse struct {
	VertexEnvelope VertexEnvelope `json:"VertexEnvelope"`
}

type VertexEnvelope struct {
	Login `json:"Login"`
	InvoiceResponse`json:"InvoiceResponse"`
	ApplicationData `json:"ApplicationData"`
}

type Login struct {

}

type InvoiceResponse struct {

}

type ApplicationData struct {

}

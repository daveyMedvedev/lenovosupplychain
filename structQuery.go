//QueryOrder : definition of Query Order
//structure to retrieve order from ledger
//###########################################

package main

type QueryOrder struct {
	OrderNumber string `json:"orderNumber"`
	Requestor   string `json:"requestor"`
	Partner     string `json:"partner"`
	DocType     string `json:"docType"`
}

type QueryInvoice struct {
	InvoiceNumber string `json:"invoiceNumber"`
	Requestor     string `json:"requestor"`
	Partner       string `json:"partner"`
	DocType       string `json:"docType"`
}

type QueryShipment struct {
	ShipmentNumber string `json:"shipmentNumber"`
	Requestor      string `json:"requestor"`
	Partner        string `json:"partner"`
	DistributorID  string `json:"distributorId"`
}

type QueryPayment struct {
	PaymentNumber string `json:"paymentNumber"`
	Requestor     string `json:"requestor"`
	Partner       string `json:"partner"`
	DistributorID string `json:"distributorId"`
}

type QueryField struct {
	FieldName  string `json:"fieldName"`
	FieldValue string `json:"fieldValue"`
}

type RichQuery struct {
	QueryName   string       `json:"queryName"`
	QueryFields []QueryField `json:"queryFields"`
}

package main

type Invoice struct {
	ObjectType          string        `json:"doctype"`
	InvoiceNumber       string        `json:"invoiceNumber"`
	From                string        `json:"from"`
	To                  string        `json:"to"`
	OriginalOrderNumber string        `json:"originalOrderNumber"`
	OriginalOrderType   string        `json:"orginalOrderType"`
	OriginalShipNotice  string        `json:"originalShipNotice"`
	Items               []ItemDetails `json:"items"`
	InvoiceAmount       string        `json:"invoiceAmount"`
	Status              string        `json:"status"`
}

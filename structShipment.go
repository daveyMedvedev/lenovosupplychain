package main

type ShippedItem struct {
	CommodityCode   string `json:"commoditycode"`
	SerialNumber    string `json:"serialNumber"`
	OrderedQuantity string `json:"orderedQuantity"`
	UOM             string `json:"uom"`
}

type Shipment struct {
	ObjectType     string        `json:"doctype"`
	ShipmentNumber string        `json:"shipmentNumber"`
	TrackingNumber string        `json:"trackingNumber"`
	DistributorID  string        `json:"distributorId"`
	ShippedItems   []ShippedItem `json:"shippedItems"`
	OrderNumber    string        `json:"OrderNumber"`
	From           string        `json:"from"`
	To             string        `json:"to"`
}

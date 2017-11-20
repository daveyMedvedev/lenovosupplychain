package main

type ReceivedItem struct {
	CommodityCode     string `json:"commoditycode"`
	SerialNumber      string `json:"serialNumber"`
	UOM               string `json:"uom"`
	DeliveredQuantity string `json:"deliveredQuantity"`
}

type GoodsReceivedNotice struct {
	ObjectType          string         `json:"doctype"`
	GoodsReceivedNumber string         `json:"goodsReceivedNumber"`
	GoodsReceivedDate   string         `json:"goodsReceivedDate"`
	ShipmentNumber      string         `json:"shipmentNumber"`
	DistributorID       string         `json:"distributorId"`
	ReceivedItems       []ReceivedItem `json:"receivedItems"`
	OrderNumber         string         `json:"orderNumber"`
	From                string         `json:"from"`
	To                  string         `json:"to"`
}

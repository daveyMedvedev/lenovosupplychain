/*Copyright IT People Corp. 2017 All Rights Reserved.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

                 http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.

******************************************************************/

// Date Created:
// Author: Dinesh
// Organization: IT People Corporation
// Last Update: Aug 08 2017
// this file will contain all  go structure used on this project
package main

type Part struct {
	PartNumber string    `json:"partnumber"`
	ManufID    string    `json:"manufid"`
	RCType     string    `json:"rctype"`
	AuditInfo  AuditInfo `json:"auditInfo"`
}

//TODO
type OrderItem struct {
	PartNumber    string `json:"partNumber"`
	ItemCondition string `json:"itemCondition"` //New or Refurbished
	Quantity      int    `json:"quantity"`      //add unit of measurement
}

type UnitOfMeasure struct {
}

type Order struct {
	OrderNumber string      `json:"poNumber"`
	SupplierID  string      `json:"supplerId"`
	Items       []OrderItem `json:"items"`
	AuditInfo   AuditInfo   `json:"auditInfo"`
	From        string      `json:"from"`
	To          string      `json:"to"`
}


type ShippedItem struct {
	PartNumber   string `json:"partNumber"`
	SerialNumber string `json:"serialNumber"`
}

type Shipment struct {
	ShipmentNumber string        `json:"shipmentNumber"`
	TrackingNumber string        `json:"trackingNumber"`
	SupplierID     string        `json:"supplerId"`
	ShippedItems   []ShippedItem `json:"partSerialNumber"`
	PONumber       string        `json:"poNumber"`
}

type AuditInfo struct {
	CreatedBy string `json:"createdBy"`
	UpdatedBy string `json:"updatedBy"`
	CreatedTS string `json:"createdTS"`
	UpdatedTS string `json:"updatedTS"`
}

type OrderStatus struct {
	PONumber string `json:"poNumber"`
	Status   string `json:"status"`
}

type RecievedItem struct {
	PartNumber      string `json:"partNumber"`
	ShipmentNumber  string `json:"shipmentNumber"`
	ReceivingStatus string `json:"receivingStatus"`
}

type Acknowledgement struct {
	PONumber   string `json:"poNumber"`
	SupplierID string `json:"supplierId"`
}

type POInvoice struct {
	PONumber string      `json:"poNumber"`
	Items    []OrderItem `json:"items"`
}

type ReturnNotice struct {
	PONumber string      `json:"poNumber"`
	Items    []OrderItem `json:"items"`
}

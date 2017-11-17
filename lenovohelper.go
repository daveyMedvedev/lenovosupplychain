package main

import (
	"dbapi"
	"encoding/json"
	"fmt"
	"runtime"
	"strconv"
	"strings"

	"github.com/hyperledger/fabric/core/chaincode/shim"
)

func getFunctionName() string {
	pc, _, _, _ := runtime.Caller(1)
	splits := strings.Split(runtime.FuncForPC(pc).Name(), "/")
	name := strings.Split(splits[len(splits)-1], ".")
	return name[len(name)-1]
}

func retrieveAndMarshalPOObject(stub shim.ChaincodeStubInterface, args []string) (PurchaseOrder, error) {
	order := PurchaseOrder{}
	var err error
	var objectBytes []byte

	keys := args

	if objectBytes, err = dbapi.QueryObject(stub, "PO", keys); err != nil || objectBytes == nil {
		err = fmt.Errorf("retrieveAndMarshal() : Failed to query PurchaseOrder object")
		return order, err
	}

	if err = json.Unmarshal(objectBytes, &order); err != nil {
		return order, fmt.Errorf("retrieveAndMarshal()  : marshalling PO failed")
	}

	if err = json.Unmarshal(objectBytes, &order); err != nil {
		return order, fmt.Errorf("retrieveAndMarshal()  : marshalling PO failed")
	}
	return order, nil

}

func retrieveAndMarshalSOObject(stub shim.ChaincodeStubInterface, args []string) (SalesOrder, error) {
	order := SalesOrder{}
	var err error
	var objectBytes []byte

	keys := args

	if objectBytes, err = dbapi.QueryObject(stub, "SO", keys); err != nil || objectBytes == nil {
		err = fmt.Errorf("retrieveAndMarshal() : Failed to query SalesOrder object")
		return order, err
	}

	if err = json.Unmarshal(objectBytes, &order); err != nil {
		return order, fmt.Errorf("retrieveAndMarshal()  : marshalling SO failed")
	}
	return order, nil

}

func checkItemDetailsTotal(items []ItemDetails, orderTotal string) bool {
	var total float64
	var tempFloat float64
	var orderTotalFloat float64
	var err error

	for _, i := range items {
		if tempFloat, err = strconv.ParseFloat(i.OrderedValue, 64); err != nil {
			return false
		}
		total = total + tempFloat
	}
	if orderTotalFloat, err = strconv.ParseFloat(orderTotal, 64); err != nil {
		return false
	}
	if total != orderTotalFloat {
		return false
	}
	return true
}

func getFormattedPurchaseOrderQuery(PurchaseOrderNumber string) string {
	return fmt.Sprintf("{\\\"selector\\\": { \\\"PurchaseOrderNumber\\\": \\\"%s\\\"}}", PurchaseOrderNumber)
}

func checkItemDetails(poOrderItems []ItemDetails, soOrderItems []ItemDetails) (bool, error) {
	orderedquantity := make(map[string][]string)
	for _, i := range poOrderItems {
		orderedquantity[i.CommodityCode] = append(orderedquantity[i.CommodityCode], i.OrderedQuantity)
		orderedquantity[i.CommodityCode] = append(orderedquantity[i.CommodityCode], i.UOM)
		fmt.Println(orderedquantity[i.CommodityCode])
	}
	for _, j := range soOrderItems {
		if quantity := orderedquantity[j.CommodityCode][0]; quantity != j.OrderedQuantity {
			return false, fmt.Errorf("Part number : " + j.CommodityCode + " invalid quantity " + quantity)
		}
		if uom := orderedquantity[j.CommodityCode][1]; uom != j.UOM {
			return false, fmt.Errorf("Part number : " + j.CommodityCode + " invalid uom " + uom)
		}
	}

	return true, nil
}

//TODO FAILS AT STRING ARRAY
func checkShipDetails(soOrderItems []ItemDetails, Shipment []ShippedItem) (bool, error) {
	orderedquantity := make(map[string]string)

	for _, i := range soOrderItems {
		orderedquantity[i.CommodityCode] = i.OrderedQuantity
		//orderedquantity[i.CommodityCode] = append(orderedquantity[i.CommodityCode], i.UOM)
		fmt.Println(i.OrderedQuantity)
	}
	for _, j := range Shipment {
		fmt.Println(orderedquantity[j.CommodityCode])
		if quantity := orderedquantity[j.CommodityCode]; quantity != j.OrderedQuantity {

			return false, fmt.Errorf("Part number : " + j.CommodityCode + " invalid quantity " + quantity)
		}
		//	if uom := orderedquantity[j.CommodityCode][1]; uom != j.UOM {
		//	return false, fmt.Errorf("Part number : " + j.CommodityCode + " invalid uom " + uom)
		//}
	}

	return true, nil
}

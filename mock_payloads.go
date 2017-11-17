package main

/*=======================================
  PO Payloads (create/query)
 =======================================*/
var purchaseOrderPayload = []byte("{\"doctype\": \"PO\",\"ponumber\": \"1234\",\"supplierid\": \"supplier\",\"vendordescription\": \"Lenovo Laptop Builder\",\"from\": \"Manu1\", \"to\": \"Lenovo\", \"items\": [{\"commoditycode\": \"1234\",\"unitprice\": \"12.50\",\"uom\": \"EA\",\"shorttext\": \"widget1\",\"orderedquantity\": \"100\",\"orderedvalue\": \"1250.00\"},{\"commoditycode\": \"1235\",\"unitprice\": \"15.50\",\"uom\": \"EA\",\"shorttext\": \"widget2\",\"orderedquantity\": \"100\",\"orderedvalue\": \"1500.00\"}],\"orderTotal\": \"2750.00\"}")

var purchaseOrderQueryPayload = "{\"orderNumber\": \"1234\", \"requestor\": \"Manu1\", \"partner\": \"Lenovo\"}"

var purchaseQueryResponse1 = "{\"doctype\":\"PO\",\"ponumber\":\"1234\",\"supplierid\":\"supplier\",\"vendordescription\":\"Lenovo Laptop Builder\",\"poissuedate\":\"\",\"postartdate\":\"\",\"poenddate\":\"\",\"paymentterm\":\"\",\"paymentdays\":\"\",\"sownum\":\"\",\"currency\":\"\",\"from\":\"Manu1\",\"to\":\"Lenovo\",\"items\":[{\"commoditycode\":\"1234\",\"effectivedate\":\"\",\"unitprice\":\"12.50\",\"oldprice\":\"\",\"uom\":\"EA\",\"shorttext\":\"widget1\",\"orderedquantity\":\"100\",\"orderedvalue\":\"1250.00\",\"invoicedquantity\":\"\",\"invoicedvalue\":\"\",\"tobedeliveredvalue\":\"\",\"tobedeliveredquantity\":\"\"},{\"commoditycode\":\"1235\",\"effectivedate\":\"\",\"unitprice\":\"15.50\",\"oldprice\":\"\",\"uom\":\"EA\",\"shorttext\":\"widget2\",\"orderedquantity\":\"100\",\"orderedvalue\":\"1500.00\",\"invoicedquantity\":\"\",\"invoicedvalue\":\"\",\"tobedeliveredvalue\":\"\",\"tobedeliveredquantity\":\"\"}],\"auditInfo\":{\"createdBy\":\"\",\"updatedBy\":\"\",\"createdTS\":\"\",\"updatedTS\":\"\"},\"orderTotal\":\"2750.00\",\"status\":\"open\"}"

var purchaseOrderByNumberQuery = "{\"orderNumber\":\"1234\", \"requestor\":\"Manu1\"}"

var purchaseOrderByNumberResponse = "[{\"doctype\":\"PO\",\"ponumber\":\"1234\",\"supplierid\":\"supplier\",\"vendordescription\":\"Lenovo Laptop Builder\",\"poissuedate\":\"\",\"postartdate\":\"\",\"poenddate\":\"\",\"paymentterm\":\"\",\"paymentdays\":\"\",\"sownum\":\"\",\"currency\":\"\",\"from\":\"Manu1\",\"to\":\"Lenovo\",\"items\":[{\"commoditycode\":\"1234\",\"effectivedate\":\"\",\"unitprice\":\"12.50\",\"oldprice\":\"\",\"uom\":\"EA\",\"shorttext\":\"widget1\",\"orderedquantity\":\"100\",\"orderedvalue\":\"1250.00\",\"invoicedquantity\":\"\",\"invoicedvalue\":\"\",\"tobedeliveredvalue\":\"\",\"tobedeliveredquantity\":\"\"},{\"commoditycode\":\"1235\",\"effectivedate\":\"\",\"unitprice\":\"15.50\",\"oldprice\":\"\",\"uom\":\"EA\",\"shorttext\":\"widget2\",\"orderedquantity\":\"100\",\"orderedvalue\":\"1500.00\",\"invoicedquantity\":\"\",\"invoicedvalue\":\"\",\"tobedeliveredvalue\":\"\",\"tobedeliveredquantity\":\"\"}],\"auditInfo\":{\"createdBy\":\"\",\"updatedBy\":\"\",\"createdTS\":\"\",\"updatedTS\":\"\"},\"orderTotal\":\"2750.00\",\"status\":\"open\"}]"

var purchaseOrderAckQuery = "{\"orderNumber\": \"1234\", \"requestor\": \"Manu1\", \"partner\": \"Lenovo\"}"

var purchaseOrderAckResponse = "{\"doctype\":\"PO\",\"ponumber\":\"1234\",\"supplierid\":\"supplier\",\"vendordescription\":\"Lenovo Laptop Builder\",\"poissuedate\":\"\",\"postartdate\":\"\",\"poenddate\":\"\",\"paymentterm\":\"\",\"paymentdays\":\"\",\"sownum\":\"\",\"currency\":\"\",\"from\":\"Manu1\",\"to\":\"Lenovo\",\"items\":[{\"commoditycode\":\"1234\",\"effectivedate\":\"\",\"unitprice\":\"12.50\",\"oldprice\":\"\",\"uom\":\"EA\",\"shorttext\":\"widget1\",\"orderedquantity\":\"100\",\"orderedvalue\":\"1250.00\",\"invoicedquantity\":\"\",\"invoicedvalue\":\"\",\"tobedeliveredvalue\":\"\",\"tobedeliveredquantity\":\"\"},{\"commoditycode\":\"1235\",\"effectivedate\":\"\",\"unitprice\":\"15.50\",\"oldprice\":\"\",\"uom\":\"EA\",\"shorttext\":\"widget2\",\"orderedquantity\":\"100\",\"orderedvalue\":\"1500.00\",\"invoicedquantity\":\"\",\"invoicedvalue\":\"\",\"tobedeliveredvalue\":\"\",\"tobedeliveredquantity\":\"\"}],\"auditInfo\":{\"createdBy\":\"\",\"updatedBy\":\"\",\"createdTS\":\"\",\"updatedTS\":\"\"},\"orderTotal\":\"2750.00\",\"status\":\"acknowledged\"}"

/*=======================================
  SO Payloads (create/query)
 =======================================*/

var salesOrderPayload = []byte("{\"doctype\": \"SO\",\"ponumber\": \"1234\",\"supplierid\": \"supplier\",\"vendordescription\": \"Lenovo Laptop Builder\",\"from\": \"Lenovo\", \"to\": \"supid1\", \"originalOrderer\": \"Manu1\", \"originalPONumber\": \"1234\", \"shipTo\": \"supid1\", \"items\": [{\"commoditycode\": \"1234\",\"unitprice\": \"12.50\",\"uom\": \"EA\",\"shorttext\": \"widget1\",\"orderedquantity\": \"100\",\"orderedvalue\": \"1250.00\"},{\"commoditycode\": \"1235\",\"unitprice\": \"15.50\",\"uom\": \"EA\",\"shorttext\": \"widget2\",\"orderedquantity\": \"100\",\"orderedvalue\": \"1500.00\"}],\"orderTotal\": \"2750.00\"}")

var salesOrderQuery = "{\"orderNumber\": \"1234\", \"requestor\": \"Lenovo\", \"partner\": \"supid1\"}"

var salesOrderQueryResponse = "{\"doctype\":\"SO\",\"ponumber\":\"1234\",\"supplierid\":\"supplier\",\"vendordescription\":\"Lenovo Laptop Builder\",\"poissuedate\":\"\",\"postartdate\":\"\",\"poenddate\":\"\",\"paymentterm\":\"\",\"paymentdays\":\"\",\"sownum\":\"\",\"currency\":\"\",\"from\":\"Lenovo\",\"to\":\"supid1\",\"originalOrderer\":\"Manu1\",\"originalPONumber\":\"1234\",\"shipTo\":\"supid1\",\"items\":[{\"commoditycode\":\"1234\",\"effectivedate\":\"\",\"unitprice\":\"12.50\",\"oldprice\":\"\",\"uom\":\"EA\",\"shorttext\":\"widget1\",\"orderedquantity\":\"100\",\"orderedvalue\":\"1250.00\",\"invoicedquantity\":\"\",\"invoicedvalue\":\"\",\"tobedeliveredvalue\":\"\",\"tobedeliveredquantity\":\"\"},{\"commoditycode\":\"1235\",\"effectivedate\":\"\",\"unitprice\":\"15.50\",\"oldprice\":\"\",\"uom\":\"EA\",\"shorttext\":\"widget2\",\"orderedquantity\":\"100\",\"orderedvalue\":\"1500.00\",\"invoicedquantity\":\"\",\"invoicedvalue\":\"\",\"tobedeliveredvalue\":\"\",\"tobedeliveredquantity\":\"\"}],\"auditInfo\":{\"createdBy\":\"\",\"updatedBy\":\"\",\"createdTS\":\"\",\"updatedTS\":\"\"},\"orderTotal\":\"2750.00\",\"status\":\"open\"}"

var salesOrderAckQuery = "{\"orderNumber\": \"1234\", \"requestor\": \"Lenovo\", \"partner\": \"supid1\"}"

var salesOrderAckResponse = "{\"doctype\":\"SO\",\"ponumber\":\"1234\",\"supplierid\":\"supplier\",\"vendordescription\":\"Lenovo Laptop Builder\",\"poissuedate\":\"\",\"postartdate\":\"\",\"poenddate\":\"\",\"paymentterm\":\"\",\"paymentdays\":\"\",\"sownum\":\"\",\"currency\":\"\",\"from\":\"Lenovo\",\"to\":\"supid1\",\"originalOrderer\":\"Manu1\",\"originalPONumber\":\"1234\",\"shipTo\":\"supid1\",\"items\":[{\"commoditycode\":\"1234\",\"effectivedate\":\"\",\"unitprice\":\"12.50\",\"oldprice\":\"\",\"uom\":\"EA\",\"shorttext\":\"widget1\",\"orderedquantity\":\"100\",\"orderedvalue\":\"1250.00\",\"invoicedquantity\":\"\",\"invoicedvalue\":\"\",\"tobedeliveredvalue\":\"\",\"tobedeliveredquantity\":\"\"},{\"commoditycode\":\"1235\",\"effectivedate\":\"\",\"unitprice\":\"15.50\",\"oldprice\":\"\",\"uom\":\"EA\",\"shorttext\":\"widget2\",\"orderedquantity\":\"100\",\"orderedvalue\":\"1500.00\",\"invoicedquantity\":\"\",\"invoicedvalue\":\"\",\"tobedeliveredvalue\":\"\",\"tobedeliveredquantity\":\"\"}],\"auditInfo\":{\"createdBy\":\"\",\"updatedBy\":\"\",\"createdTS\":\"\",\"updatedTS\":\"\"},\"orderTotal\":\"2750.00\",\"status\":\"acknowledged\"}"

/*=======================================
  ACK Payloads (create/query)
 =======================================*/
var ackPayload = []byte("{\"doctype\":\"Acknowledgement\",\"ackNumber\":\"4321\",\"documentType\":\"PO\",\"documentNumber\":\"1234\",\"from\":\"Lenovo\",\"to\":\"Manu1\"}")

var ackSOPayload = []byte("{\"doctype\":\"Acknowledgement\",\"ackNumber\":\"4322\",\"documentType\":\"SO\",\"documentNumber\":\"1234\",\"from\":\"supid1\",\"to\":\"Lenovo\"}")

/*=======================================
  SHipment Payloads (create/query)
 =======================================*/

var shipmentPayload = []byte("{\"shipmentNumber\": \"1234\",\"trackingnumber\": \"4567\", \"shippedItems\": [{\"partNumber\":\"1111\",\"shippedQuantity\":100},{\"partNumber\": \"2222\", \"shippedQuantity\": 200}], \"supplierId\": \"supid1\", \"ponumber\": \"0001\", \"distributorId\": \"Lenovo\", \"to\": \"Manu1\"}")

var shipmentQuery = "{\"shipmentNumber\": \"1234\", \"requestor\": \"Manu1\", \"partner\": \"Lenovo\"}"
var shipmentQueryResponse = "{\"shipmentNumber\": \"1234\",\"trackingnumber\": \"4567\", \"shippedItems\": [{\"partNumber\":\"1111\",\"shippedQuantity\":100},{\"partNumber\": \"2222\", \"shippedQuantity\": 200}], \"supplierId\": \"supid1\", \"ponumber\": \"0001\", \"distributorId\": \"Lenovo\", \"to\": \"Manu1\"}"

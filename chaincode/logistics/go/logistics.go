
package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"strconv"
	"strings"
	// "time"

	"github.com/hyperledger/fabric/core/chaincode/shim"
	sc "github.com/hyperledger/fabric/protos/peer"
)

// Define the Smart Contract structure
type SmartContract struct {
}
// For the simplification taking all fields as string data type.
type Shipment struct {
	Id string `json:"id"`
	Shipment_Name string `json:"shipment_name"`
	SellerId string `json:"seller_id"`
	SellerName string `json:"seller_name"`
	SellerLocation string `json:"seller_location"`
	BuyerId string `json:"buyer_id"`
	BuyerName string `json:"buyer_name"`
	BuyerLocation string `json:"buyer_location"`
	LogisticProviderId string `json:"logistic_provider_id"`
	Status string `json:"status"`
}
// ToDo Services & Calculating temprature
type TimeRaster struct {
	ShipmentId string `json:"shipment_id"`
	TimeRaster string  `json:"time_raster"`
	Temprature float32 `json:"temprature"`
} 

func (s *SmartContract) Init(APIstub shim.ChaincodeStubInterface) sc.Response {
	return shim.Success(nil)
}

func (s *SmartContract) Invoke(APIstub shim.ChaincodeStubInterface) sc.Response {

	// Retrieve the requested Smart Contract function and arguments
	function, args := APIstub.GetFunctionAndParameters()
	// Route to the appropriate handler function to interact with the ledger appropriately
	if function == "initLedger" {
		return s.initLedger(APIstub)
	}else if function == "logTimeRaster" {
		return s.logTimeRaster(APIstub, args)
	}else if function == "addShipment" {
		return s.addShipment(APIstub, args)
	}else if function == "changeShipmentStatus" {
		return s.changeShipmentStatus(APIstub, args)
	}else if function == "queryShipmentsForSeller" {
		return s.queryShipmentsForSeller(APIstub, args)
	}else if function == "queryShipmentsForLogisticprovider" {
		return s.queryShipmentsForLogisticprovider(APIstub, args)
	}else if function == "queryShipmentsForBuyer" {
		return s.queryShipmentsForBuyer(APIstub, args)
	}else if function == "queryTimeRasterForShipment" {
		return s.queryTimeRasterForShipment(APIstub, args)
	}else if function == "queryShipment" {
		return s.queryShipment(APIstub, args)
	}

	return shim.Error("Invalid Smart Contract function name.")
}

func (s *SmartContract) initLedger(APIstub shim.ChaincodeStubInterface) sc.Response {
	
	shimpments := []Shipment{
		Shipment{Id: "001", Shipment_Name: "Shipment1", SellerId: "111", SellerName: "Pavan", SellerLocation: "Banglore", BuyerId: "222", BuyerName: "Hemanth", BuyerLocation: "Hyderbad", LogisticProviderId: "333", Status : "accepted"},
		Shipment{Id: "002", Shipment_Name: "Shipment2", SellerId: "111", SellerName: "Pavan", SellerLocation: "Banglore", BuyerId: "2222", BuyerName: "Ananth", BuyerLocation: "Pune", LogisticProviderId: "3333", Status : "rejected"},
		Shipment{Id: "003", Shipment_Name: "Shipment3", SellerId: "111", SellerName: "Pavan", SellerLocation: "Banglore", BuyerId: "22222", BuyerName: "Akhil", BuyerLocation: "Mumbai", LogisticProviderId: "333", Status : "accepted"},
	}

	i := 0
	for i < len(shimpments) {
		fmt.Println("i is ", i)
		shimpmentAsBytes, _ := json.Marshal(shimpments[i])
		APIstub.PutState("00"+strconv.Itoa(i), shimpmentAsBytes)
		fmt.Println("Added", shimpments[i])
		i = i + 1
	}

	return shim.Success(nil)
}



//////////////////////////////////////////////////////////
// Converts an Object to a JSON String
//////////////////////////////////////////////////////////
func JSONtoShipmentReq(areq []byte) (Shipment, error) {
	shipment := Shipment{}
	err := json.Unmarshal(areq, &shipment)
	if err != nil {
		fmt.Println("JSONtoShipmentReq error: ", err)
		return shipment, err
	}
	return shipment, err
}



//ToDO
func (s *SmartContract) logTimeRaster(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	/*if len(args) != 6 {
		return shim.Error("Incorrect number of arguments. Expecting 6")
	}*/

	// Check status of the shimpment is intransit or not.
	response := s.queryShipment(APIstub, args)
	
	shipmentR, err := JSONtoShipmentReq(response.Payload)

	if shipmentR.Status != "intransit" {
		fmt.Println("logTimeRaster() : Cannot accept log timeRasterr for this shimpment : ", args[0])
		return shim.Error("logTimeRaster() : Cannot accept log timeRasterr for this shimpment : " + args[0])
	}

	// var temp_in_float float32
	value, err := strconv.ParseFloat(args[2], 32)
	if err != nil {
	    return shim.Error("Invalid data type for temprature..!")
	}
	temp_in_float := float32(value)
	
	var timeRaster = TimeRaster{ShipmentId: args[0], TimeRaster: args[1], Temprature: temp_in_float}

	timeRaserAsBytes, _ := json.Marshal(timeRaster)
	APIstub.PutState(args[3], timeRaserAsBytes)

	return shim.Success(timeRaserAsBytes)
}


func (s *SmartContract) addShipment(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	/*if len(args) != 5 {
		return shim.Error("Incorrect number of arguments. Expecting 5")
	}*/
	var shipment = Shipment{Id: args[0], Shipment_Name: args[1], SellerId: args[2],SellerName: args[3], SellerLocation: args[4], BuyerId: args[5], BuyerName: args[6], BuyerLocation: args[7], LogisticProviderId: args[8], Status : args[9]}

	shipmentAsBytes, _ := json.Marshal(shipment)
	APIstub.PutState(args[0], shipmentAsBytes)

	return shim.Success(shipmentAsBytes)
}

func (s *SmartContract) changeShipmentStatus(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	/*if len(args) != 2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}*/
 	shipmentAsBytes, _ := APIstub.GetState(args[0])
	shipment := Shipment{}
	statusChanger := args[2]  // id of the user trying to change the asset status

	json.Unmarshal(shipmentAsBytes, &shipment)               // ToDo change according to status codes
	if shipment.Status == "store" && shipment.SellerId == statusChanger && args[1] == "intransit" {                // Only shaller can change status from store to Inransit
		shipment.Status = args[1] //  "intransit"
	}else if shipment.Status == "intransit" && shipment.LogisticProviderId == statusChanger && args[1] == "delivered" {
		shipment.Status = args[1] // "delivered"
	}else if shipment.Status == "delivered" && shipment.BuyerId == statusChanger && args[1] == "accepted" {
		shipment.Status = args[1] // "accepted"  #TODO validation of temprature 
	}else if shipment.Status == "delivered" && shipment.BuyerId == statusChanger && args[1] == "rejected" {
		shipment.Status = args[1] // "rejected" #TODO validation of temprature 
	}else{
		return shim.Error("Invalid access to change this status")
	}

	shipmentAsBytes, _ = json.Marshal(shipment)
	APIstub.PutState(args[0], shipmentAsBytes)

	return shim.Success(nil)
}


func (s *SmartContract) queryTimeRasterForShipment(APIstub shim.ChaincodeStubInterface,args []string) sc.Response {
	requesterID := strings.ToLower(args[0])
	queryStringForShipment := fmt.Sprintf("{\"selector\":{\"shipment_id\":\"%s\"}}", requesterID)

	queryResults, err := getQueryResultForQueryString(APIstub, queryStringForShipment)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(queryResults)
}



func (s *SmartContract) queryShipmentsForSeller(APIstub shim.ChaincodeStubInterface,args []string) sc.Response {
	requesterID := strings.ToLower(args[0])
	queryStringForSeller := fmt.Sprintf("{\"selector\":{\"seller_id\":\"%s\"}}", requesterID)

	queryResults, err := getQueryResultForQueryString(APIstub, queryStringForSeller)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(queryResults)
}


func (s *SmartContract) queryShipmentsForLogisticprovider(APIstub shim.ChaincodeStubInterface,args []string) sc.Response {
	 
	requesterID := strings.ToLower(args[0])
	queryStringForLogisticProvider := fmt.Sprintf("{\"selector\":{\"logistic_provider_id\":\"%s\"}}", requesterID)

	queryResults, err := getQueryResultForQueryString(APIstub, queryStringForLogisticProvider)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(queryResults)
}


func (s *SmartContract) queryShipmentsForBuyer(APIstub shim.ChaincodeStubInterface,args []string) sc.Response {
	
	requesterID := strings.ToLower(args[0])
	queryStringForBuyer := fmt.Sprintf("{\"selector\":{\"buyer_id\":\"%s\"}}", requesterID)

	queryResults, err := getQueryResultForQueryString(APIstub, queryStringForBuyer)
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(queryResults)
}

func getQueryResultForQueryString(stub shim.ChaincodeStubInterface, queryString string) ([]byte, error) {

	fmt.Printf("- getQueryResultForQueryString queryString:\n%s\n", queryString)

	resultsIterator, err := stub.GetQueryResult(queryString)
	if err != nil {
		return nil, err
	}
	defer resultsIterator.Close()

	// buffer is a JSON array containing QueryRecords
	var buffer bytes.Buffer
	buffer.WriteString("[")

	bArrayMemberAlreadyWritten := false
	for resultsIterator.HasNext() {
		queryResponse, err := resultsIterator.Next()
		if err != nil {
			return nil, err
		}
		// Add a comma before array members, suppress it for the first array member
		if bArrayMemberAlreadyWritten == true {
			buffer.WriteString(",")
		}
		buffer.WriteString("{\"Key\":")
		buffer.WriteString("\"")
		buffer.WriteString(queryResponse.Key)
		buffer.WriteString("\"")

		buffer.WriteString(", \"Record\":")
		// Record is a JSON object, so we write as-is
		buffer.WriteString(string(queryResponse.Value))
		buffer.WriteString("}")
		bArrayMemberAlreadyWritten = true
	}
	buffer.WriteString("]")

	fmt.Printf("- getQueryResultForQueryString queryResult:\n%s\n", buffer.String())

	return buffer.Bytes(), nil
}

func (s *SmartContract) queryShipment(APIstub shim.ChaincodeStubInterface, args []string) sc.Response {

	/*if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}*/

	shipmentAsBytes, _ := APIstub.GetState(args[0])
	return shim.Success(shipmentAsBytes)
}

// The main function is only relevant in unit test mode. Only included here for completeness.
func main() {

	// Create a new Smart Contract
	err := shim.Start(new(SmartContract))
	if err != nil {
		fmt.Printf("Error creating new Smart Contract: %s", err)
	}
}

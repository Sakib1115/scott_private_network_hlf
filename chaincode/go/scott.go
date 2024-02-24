/* Licensed to the Apache Software Foundation (ASF) under one
 * or more contributor license agreements.  See the NOTICE file
 * distributed with this work for additional information
 * regarding copyright ownership.  The ASF licenses this file
 * to you under the Apache License, Version 2.0 (the
 * "License"); you may not use this file except in compliance
 * with the License.  You may obtain a copy of the License at
 *
 * http://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing,
 * software distributed under the License is distributed on an
 * "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY
 * KIND, either express or implied.  See the License for the
 * specific language governing permissions and limitations
 * under the License.
 */

/*
 * The sample smart contract for documentation topic:
 * Writing Your First Blockchain Application
 */

 package main

 /* Imports
  * 4 utility libraries for formatting, handling bytes, reading and writing JSON, and string manipulation
  * 2 specific Hyperledger Fabric specific libraries for Smart Contracts
  */
 import (
	 "encoding/json"
	 "fmt"
	//  "reflect"
    //  "bytes"
    //  "strconv"
    //  "time"
	 "github.com/hyperledger/fabric-chaincode-go/shim"
    "github.com/hyperledger/fabric-protos-go/peer"
	//  id "github.com/hyperledger/fabric/core/chaincode/shim/ext/cid"
 )
// SmartContract  structure
 type Scott struct {
 }

type Refund struct {
	 DocType                string  `json:"docType"`
	 RID                    string  `json:"rid"`
	 ContractNo             string  `json:"contractno"`
	 CName                  string  `json:"cname"`
	 EffectiveDate          string  `json:"effectivedate"`
	 StartTime              string  `json:"starttime "`
	 EndTime                string  `json:"endtime "`
	 AgreementPurchasePrice float64 `json:"purchaseprice"`
	 CancelDate             string  `json: "canceldate"`
	 ApprovalFromAdmin      string  `json: "approvalfromadmin"`
	 ID                     string  `json: "id"`
	 Years                  int     `json: "years"`
	 PresentMonth           float64 `json: "presentmonth"`
}


type User struct {
	DocType			string `json:"docType"`
	// UserID          string `json:"id"`
	FirstName       string `json:"fName"`
	LastName        string `json:"lName"`	
	EmailId         string `json:"EmailId"`
	MobileNo        string `json:"mobile"`
	Password		string `json:"password"`
	Country			string `json:"country"`
	VINnumber		string `json:"vinnumber"`
}


type Admin struct {
	DocType			string `json:"docType"`
	ID          	string `json:"adminId"`
	FirstName   	string `json:"adminFirstName"`
	LastName    	string `json:"adminLastName"`
	MobileNo    	string `json:"adminPhone"`
	EmailID     	string `json:"adminEmail"`
	Cts         	string `json:"cts"`
	DateOfBirth		string `json:"dob"`
	Location       	string `json:"location"`
	Password		string `json:"password"`
	
}

type PrePaidAgreement struct {
	DocType				string `json:"docType"`
	PPId				string `json:"ppid"`
	CustomerName        string `json:"customername"`
	CustomerAddress    	string `json:"customeraddress"`
	VehicleYear    		string `json:"vehicleyear"`
	VehicleName     	string `json:"vehiclename"`
	VehicleModel		string `json:"vehiclemodel"`
	OdoMeter			string `json:"odometer"`
	VINNumber			string `json:"vinnumber"`
	CustomerPhoneNo     string `json:"customerphone"`
	DealersName			string `json:"dealersname"`
	DealersAddress		string `json:"dealersaddress"`
	DealersPhonenumber	string `json:"dealersphone"`
	LienHolder			string `json:"lienholder"`
	OriginalInServiceDate string `json:"oisd"`
	EffectiveDate		string `json:"effectivedate"`
	Deductible			string `json:"deductible"`
	Term				string `json:"term"`
	StartTime 			string `json:"starttime "`
	EndTime 			string `json:"endtime "`
	AgreementPurchasePrice	string `json:"purchaseprice"`
	CoverageType		string `json:"coveragetype"`
	
}
type PaintlessRepair struct {
	DocType				string `json:"docType"`
	PRId				string `json:"prid"`
	CustomerName        string `json:"customername"`
	CustomerAddress    	string `json:"customeraddress"`
	CustomerPhoneNo     string `json:"customerphone"`
	VehicleYear    		string `json:"vehicleyear"`
	VehicleName     	string `json:"vehiclename"`
	VehicleModel		string `json:"vehiclemodel"`
	VehicleMileage		string `json:"mileage"`
	VINNumber			string `json:"vinnumber"`	
	DealersName			string `json:"dealersname"`
	DealersAddress		string `json:"dealersaddress"`
	DealersPhonenumber	string `json:"dealersphone"`
	VehiclePrice		string `json:"vehicleprice"`
	Term				string `json:"term"`
	Deductible			string `json:"deductible"`	
	
}

type BatteryAgreement struct {
	DocType				string `json:"docType"`
	BId					string `json:"bid"`
	CustomerName        string `json:"customername"`
	CustomerAddress    	string `json:"customeraddress"`
	CustomerPhoneNo     string `json:"customerphone"`
	VehicleYear    		string `json:"vehicleyear"`
	VehicleName     	string `json:"vehiclename"`
	VehicleModel		string `json:"vehiclemodel"`
	VINNumber			string `json:"vinnumber"`
	VehicleMileage		string `json:"vehiclemil"`	
	IssuingLocation		string `json:"issuingloc"`
	IssuingId			string `json:"issuingid"`
	IAddress			string `json:"iaddress"`
	IPhonenumber		string `json:"iphone"`
	SellingAssociate	string `json:"sassociate"`
	LienHolder			string `json:"lienholder"`
	LAddress 			string `json:"laddress"`
	Deductible			string `json:"deductible"`
	Term				string `json:"term"`
	Cost 	 			string `json:"cost"`
	ADate				string 	`json:"adate"`
}

type ServiceAgreement struct {
	DocType				string `json:"docType"`
	SId					string `json:"sid"`
	VehicleYear    		string `json:"vehicleyear"`
	VehicleName     	string `json:"vehiclename"`
	VehicleModel		string `json:"vehiclemodel"`
	VINNumber			string `json:"vinnumber"`
	VehicleMileage		string `json:"vehiclemileage"`	
	VehiclePrice		string `json:"vehicleprice"`
	CustomerName        string `json:"customername"`
	CustomerAddress    	string `json:"customeraddress"`
	CustomerPhoneNo     string `json:"customerphone"`
	DealersName			string `json:"dealersname"`
	DealersAddress		string `json:"dealersaddress"`
	DealersPhonenumber	string `json:"dealersphone"`
	LienHolder			string `json:"lienholder"`
	LAddress 			string `json:"laddress"`
	AgreementPrice		string `json:"agreementprice"`
	AgreementDate		string `json:"agreementdate"`
	Term				string `json:"term"`
	Coverage			string `json:"coverage"`		
	Deductible			string `json:"deductible"`
	Surcharges			string `json:"surcharges"`
	OptionalCoverage	string `json:"oc"`
}
type Claim struct {
	DocType			string `json:"docType"`
	ContractNo     string `json:"contractno"`
	FName			string `json:"fname"`
	LName			string `json:"lname"`
	MobileNo       string `json:"mobile"`
	Dealer       	string `json:"dealer"`	
	Address         string `json:"address"`
	EmailId         string `json:"emailid"`
	DateOfLoss		string `json:"dateofloss"`
	Mileage			string `json:"mileage"`
	Insurance		string `json:"insurance"`
}	

// *************ADMINISTRATION*****************//

func (s *SmartContract) createAdministration(APIstub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 1 {
		return shim.Error("Incorrect number of argument, expected 1")
    }
	var  admin Admin
	err := json.Unmarshal([]byte(args[0]),&admin)
	fmt.Println("admin", admin)
	if err != nil {
		return shim.Error("Input argument unmarshling error")
	}
	if len(admin.ID) == 0 {
		return shim.Error("Admin Id is mandatory")
	}
	if len(admin.EmailID) == 0 {
		return shim.Error("Emailid is mandatory")
	}
	if len(admin.MobileNo) == 0 {
		return shim.Error("phone is mandatory")
	}
	if len(admin.Password) == 0 {
		return shim.Error("Password is mandatory")
	}
	if recordBytes, _ := APIstub.GetState(admin.ID); len(recordBytes) > 0 {
		return shim.Error("Admin already registered. Provide an unique id")
	}
	var Admin = Admin{
		DocType:		"admin",
        ID:             admin.ID,
		FirstName:      admin.FirstName,
		LastName:   	admin.LastName,
		EmailID:    	admin.EmailID,
		MobileNo:    	admin.MobileNo,
		DateOfBirth:   	admin.DateOfBirth,
	    Cts:      	    admin.Cts,
		Location:        admin.Location,
		Password:   	admin.Password,
	}

	
	AsByte, _ := json.Marshal(Admin)
	APIstub.PutState(admin.ID, AsByte)
    Data := map[string]interface{}{
		"trxnID":  APIstub.GetTxID(),
		"status":  true,
		"message": "success",
		"data":    admin,
	}
	userJSON, _ := json.Marshal(Data)
	
    return shim.Success(userJSON)
}
//update Admin
func (s *SmartContract) updateAdministration(APIstub shim.ChaincodeStubInterface, args []string) peer.Response {

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}
	var admin Admin
	err := json.Unmarshal([]byte(args[0]), &admin)
	if err != nil {
		return shim.Error("Input arguments unmarshling error")
	}
	fmt.Println("unmarshal admin details", admin)
	AsByte, err := APIstub.GetState(admin.ID)

	if err != nil {
		return shim.Error("error while getting the user")
	} else if len(AsByte) == 0 {
		return shim.Error(" id : " + admin.ID + " does not exist")
	}
	var update Admin
	err = json.Unmarshal(AsByte, &update)
	if err != nil {
		return shim.Error("Existing admin details unmarshaling error : " + string(err.Error()))
	}

	update.FirstName = admin.FirstName
	update.EmailID = admin.EmailID
	update.LastName = admin.LastName
	update.MobileNo = admin.MobileNo
	update.Location = admin.Location
	update.DateOfBirth = admin.DateOfBirth
	update.Password = admin.Password

	updateAsByte, err := json.Marshal(update)
	err = APIstub.PutState(update.ID, updateAsByte)
	if err != nil {
		return shim.Error("Error while inserting the data into Ledger : " + err.Error())
	}
	finalData := map[string]interface{}{
		"status":  "true",
		"message": "successfully changed details",
		"trxnID":  APIstub.GetTxID(),
		"data":    update,
	}
	DataAsByte, err := json.Marshal(finalData)
	return shim.Success(DataAsByte)
}
func (s *SmartContract) queryAdministration(APIstub shim.ChaincodeStubInterface, args []string) peer.Response {

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}

	AsBytes, err := APIstub.GetState(args[0])
	if err !=nil{
		return shim.Error("Failed to read from world state  ")
	}
	if AsBytes == nil {
		return shim.Error("id does not exist")
	}
	return shim.Success(AsBytes)
}

// ****************************USER*****************//
func (s *SmartContract) createUser(APIstub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 1 {
		return shim.Error("Incorrect number of argument, expected 1")

	}
	var user User

	err := json.Unmarshal([]byte(args[0]), &user)
	fmt.Println("userr", user)
	if err != nil {
		return shim.Error("Input argument unmarshling error")
	}
	if len(user.FirstName) == 0 {
		return shim.Error(" Name is mandatory")
	}
	if len(user.EmailId) == 0 {
		return shim.Error("Emailid is mandatory")
	}
	if len(user.VINnumber) == 0 {
		return shim.Error("VIN number is mandatory")
	}
	if len(user.MobileNo) == 0 {
		return shim.Error("Mobile number is mandatory")
	}
	// if reflect.TypeOf(user.MobileNo).Kind() != reflect.Int {
	// 	return shim.Error("Mobile number is mandatory")
	// }	
	
	if recordBytes, _ := APIstub.GetState(user.MobileNo); len(recordBytes) > 0 {
		return shim.Error("user already registered. Provide an unique mobile number")
	}

	
	var usr = User{
		DocType:"user",
		FirstName:   user.FirstName,
		LastName:   user.LastName,
		EmailId:     user.EmailId,
		MobileNo:    user.MobileNo,
		Country: 	user.Country,
		Password:	user.Password,
		VINnumber:	user.VINnumber,		
	}
	userAsByte, _ := json.Marshal(usr)
	APIstub.PutState(user.MobileNo, userAsByte)
    userData := map[string]interface{}{
		"trxnID":  APIstub.GetTxID(),
		"status":  true,
		"message": "success",
		"data":    usr,
	}
    userJSON, _ := json.Marshal(userData)

	return shim.Success(userJSON)

}
//update user
func (s *SmartContract) updateUser(APIstub shim.ChaincodeStubInterface, args []string) peer.Response {

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}
	var user User
	err := json.Unmarshal([]byte(args[0]), &user)
	if err != nil {
		return shim.Error("Input arguments unmarshling error")
	}
	fmt.Println("unmarshal user details", user)
	AsByte, err := APIstub.GetState(user.MobileNo)

	if err != nil {
		return shim.Error("error while getting the user")
	} else if len(AsByte) == 0 {
		return shim.Error(" id : " + user.MobileNo + " does not exist")
	}
	var update User
	err = json.Unmarshal(AsByte, &update)
	if err != nil {
		return shim.Error("Existing user details unmarshaling error : " + string(err.Error()))
	}

	update.FirstName = user.FirstName
	update.LastName = user.LastName
	update.EmailId = user.EmailId
	update.Password = user.Password
	// update.VINnumber = user.VINnumber
	update.Country = user.Country

	updateAsByte, err := json.Marshal(update)
	err = APIstub.PutState(update.MobileNo, updateAsByte)
	if err != nil {
		return shim.Error("Error while inserting the data into Ledger : " + err.Error())
	}
	finalData := map[string]interface{}{
		"status":  "true",
		"message": "successfully changed details",
		"trxnID":  APIstub.GetTxID(),
		"data":    update,
	}
	DataAsByte, err := json.Marshal(finalData)
	return shim.Success(DataAsByte)
}


func (s *SmartContract) queryUser(APIstub shim.ChaincodeStubInterface, args []string) peer.Response {

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments")
	}

	AsBytes, err := APIstub.GetState(args[0])
	if err !=nil{
		return shim.Error("Failed to read from world state  ")
	}
	if AsBytes == nil {
		return shim.Error("id does not exist")
	}
	return shim.Success(AsBytes)
}
//userLogin-passing mobile.VIN number and password,checks if mob no exist and 
// checks if password & VIN matches.if so it return success Login

func (s *SmartContract) userLogin(APIstub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}
	var user User
	err := json.Unmarshal([]byte(args[0]), &user)
	if err != nil {
		return shim.Error("Input arguments unmarshling error")
	}
	if len(user.VINnumber) == 0 {
		return shim.Error("VIN number is mandatory")
	}
	if len(user.Password) == 0 {
		return shim.Error("Password is mandatory")
	}
	fmt.Println("unmarshal user details", user)
	AsByte, err := APIstub.GetState(user.MobileNo)

	if err != nil {
		return shim.Error("error while getting the user")
	} else if len(AsByte) == 0 {
		return shim.Error(" id : " + user.MobileNo + " does not exist")
	}
	var update User
	err = json.Unmarshal(AsByte, &update)
	if err != nil {
		return shim.Error("Existing user details unmarshaling error : " + string(err.Error()))
	}
	fmt.Println("update.password", update.Password)
	fmt.Println("user.Password", user.Password)
	fmt.Println("update.vinnumber", update.VINnumber)
	fmt.Println("user.vinnumber", user.VINnumber)
	
	if (update.Password != user.Password ){
		return shim.Error("password not matching")
	}
	if (update.VINnumber == user.VINnumber){
		return shim.Success([]byte("Login is successful "))	
	}
	
		return shim.Error("VIN number not matching")
	

}
//**************Find by mobile no and last name*******//
//*************Search by user mobile number and last name and returns the record if exists******//

func (s *SmartContract) findByNumberLastName(APIstub shim.ChaincodeStubInterface, args []string) peer.Response {

	if len(args) <2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}
	fmt.Println("- start findByNumberLastName",args[0],args[1])
		
	mNumber := args[0]
	lName := args[1]

	fmt.Println("- start findByNumberLastName",mNumber,lName)
	AsByte, err := APIstub.GetState(mNumber)

	if err != nil {
		return shim.Error("error while getting the user")
	} else if len(AsByte) == 0 {
		return shim.Error(" id : " + mNumber + " does not exist")
	}
	var update User
	err = json.Unmarshal(AsByte, &update)
	if err != nil {
		return shim.Error("Existing user details unmarshaling error : " + string(err.Error()))
	}
	fmt.Println("Db mob no", update.MobileNo)
	fmt.Println("user.Password", mNumber)
	fmt.Println("db last name", update.LastName)
	fmt.Println("user.lastname", lName)
	
	if (update.LastName == lName){
		return shim.Success(AsByte)
	}
	
		return shim.Error("last name not matching")
	

}

//**************Find by contract no and last name*******//
//**********Search by claim contract number and last name and returns the record if exists***/

func (s *SmartContract) findByContractLastName(APIstub shim.ChaincodeStubInterface, args []string) peer.Response {

	if len(args) <2 {
		return shim.Error("Incorrect number of arguments. Expecting 2")
	}
	fmt.Println("- start findByContractLastName",args[0],args[1])
		
	contract := args[0]
	lName := args[1]

	fmt.Println("- start findBycontractLastName",contract,lName)
	AsByte, err := APIstub.GetState(contract)

	if err != nil {
		return shim.Error("error while getting the user")
	} else if len(AsByte) == 0 {
		return shim.Error(" id : " + contract + " does not exist")
	}
	var update Claim
	err = json.Unmarshal(AsByte, &update)
	if err != nil {
		return shim.Error("Existing user details unmarshaling error : " + string(err.Error()))
	}
	fmt.Println("Db contract no", update.ContractNo)
	fmt.Println("contract", contract)
	fmt.Println("db name", update.LName)
	fmt.Println("name", lName)
		
	if (update.LName == lName){
		return shim.Success(AsByte)
	}
	
		return shim.Error("Error name not matching")
	

}


// ****************Prepaid Agreement*****************//
func (s *SmartContract) createPrePaidAgreement(APIstub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 1 {
		return shim.Error("Incorrect number of argument, expected 1")
	}
	var ppa PrePaidAgreement

	err := json.Unmarshal([]byte(args[0]), &ppa)
	fmt.Println("ppa", ppa)
	if err != nil {
		return shim.Error("Input argument unmarshling error")
	}
	if len(ppa.PPId) == 0 {
		return shim.Error(" Id is mandatory")
	}
	if len(ppa.CustomerPhoneNo) == 0 {
		return shim.Error("CustomerPhoneNo is mandatory")
	}
	if len(ppa.VINNumber) == 0 {
		return shim.Error("VINNumber is mandatory")
	}
	// if reflect.TypeOf(user.MobileNo).Kind() != reflect.Int {
	// 	return shim.Error("Mobile number is mandatory")
	// }	
	
	if recordBytes, _ := APIstub.GetState(ppa.PPId); len(recordBytes) > 0 {
		return shim.Error("PrePaidAgreement id exists,Please provide unique id")
	}
	recordBytesb, _ := APIstub.GetState(ppa.CustomerPhoneNo)
	if len(recordBytesb) == 0 {
		return shim.Error("Customer is not registered,Please give a registered id")
	}
	var ppaData = PrePaidAgreement{
		DocType:"pre paid",
		PPId:				ppa.PPId,
		CustomerName:  		 ppa.CustomerName,
		CustomerAddress:  	 ppa.CustomerAddress,
		VehicleYear:    	 ppa.VehicleYear,
		VehicleName:   		 ppa.VehicleName,
		VehicleModel:		 ppa.VehicleModel,
		OdoMeter:			 ppa.OdoMeter,	
		VINNumber:   		 ppa.VINNumber,
		CustomerPhoneNo:   	 ppa.CustomerPhoneNo,
		DealersName:    	 ppa.DealersName,
		DealersAddress:		 ppa.DealersAddress,
		DealersPhonenumber:	 ppa.DealersPhonenumber,	
		LienHolder:   	 	 ppa.LienHolder,
		OriginalInServiceDate: ppa.OriginalInServiceDate,
		EffectiveDate:		 ppa.EffectiveDate,
		Deductible:			 ppa.Deductible,
		Term:				 ppa.Term,	
		StartTime:   	 	 ppa.StartTime,
		EndTime: 			 ppa.EndTime,
		AgreementPurchasePrice:	 ppa.AgreementPurchasePrice,
		CoverageType:		 ppa.CoverageType,			
	}
	AsByte, _ := json.Marshal(ppaData)
	APIstub.PutState(ppa.PPId, AsByte)
    dataByte := map[string]interface{}{
		"trxnID":  APIstub.GetTxID(),
		"status":  true,
		"message": "success",
		"data":    ppa,
	}
    userJSON, _ := json.Marshal(dataByte)

	return shim.Success(userJSON)

}
func (s *SmartContract) queryPrepaidAgreement(APIstub shim.ChaincodeStubInterface, args []string) peer.Response {

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}
	AsBytes, err := APIstub.GetState(args[0])
	if err !=nil{
		return shim.Error("Failed to read from world state  ")
	}
	if AsBytes == nil {
		return shim.Error("id does not exist")
	}
	return shim.Success(AsBytes)
}

// ****************Battery Agreement*****************//
func (s *SmartContract) createBatteryAgreement(APIstub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 1 {
		return shim.Error("Incorrect number of argument, expected 1")
	}
	var bAgmt BatteryAgreement

	err := json.Unmarshal([]byte(args[0]), &bAgmt)
	fmt.Println("battery", bAgmt)
	if err != nil {
		return shim.Error("Input argument unmarshling error")
	}
	if len(bAgmt.BId) == 0 {
		return shim.Error(" Id is mandatory")
	}
	if len(bAgmt.CustomerPhoneNo) == 0 {
		return shim.Error("CustomerPhoneNo is mandatory")
	}
	if len(bAgmt.VINNumber) == 0 {
		return shim.Error("VINNumber is mandatory")
	}
	// if reflect.TypeOf(user.MobileNo).Kind() != reflect.Int {
	// 	return shim.Error("Mobile number is mandatory")
	// }	
	
	if recordBytes, _ := APIstub.GetState(bAgmt.BId); len(recordBytes) > 0 {
		return shim.Error("BatteryAgreement id exists,Please provide unique id")
	}
	recordBytesb, _ := APIstub.GetState(bAgmt.CustomerPhoneNo)
	if len(recordBytesb) == 0 {
		return shim.Error("Customer is not registered,Please give a registered id")
	}
	var bData = BatteryAgreement{
		DocType:"battery",
		BId:	bAgmt.BId,
		CustomerName:  		 bAgmt.CustomerName,
		CustomerAddress:  	 bAgmt.CustomerAddress,
		CustomerPhoneNo:   	 bAgmt.CustomerPhoneNo,
		VehicleYear:    	 bAgmt.VehicleYear,
		VehicleName:   		 bAgmt.VehicleName,
		VehicleModel:		 bAgmt.VehicleModel,
		VINNumber:   		 bAgmt.VINNumber,
		VehicleMileage:		 bAgmt.VehicleMileage,
		IssuingLocation:     bAgmt.IssuingLocation,
		IssuingId:			 bAgmt.IssuingId,
		IAddress:			 bAgmt.IAddress,	
		LienHolder:   	 	 bAgmt.LienHolder,
		LAddress:   	 	 bAgmt.LAddress,
		IPhonenumber: 		 bAgmt.IPhonenumber,
		SellingAssociate:	 bAgmt.SellingAssociate,
		Deductible:			 bAgmt.Deductible,
		Term:				 bAgmt.Term,	
		Cost:   	 		 bAgmt.Cost,
		ADate: 				 bAgmt.ADate,
					
	}
	AsByte, _ := json.Marshal(bData)
	APIstub.PutState(bAgmt.BId, AsByte)
    dataByte := map[string]interface{}{
		"trxnID":  APIstub.GetTxID(),
		"status":  true,
		"message": "success",
		"data":    bAgmt,
	}
    userJSON, _ := json.Marshal(dataByte)

	return shim.Success(userJSON)

}
func (s *SmartContract) queryBatteryAgreement(APIstub shim.ChaincodeStubInterface, args []string) peer.Response {

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}
	AsBytes, err := APIstub.GetState(args[0])
	if err !=nil{
		return shim.Error("Failed to read from world state  ")
	}
	if AsBytes == nil {
		return shim.Error("id does not exist")
	}
	return shim.Success(AsBytes)
}
// ****************Paintless Dent Repair*****************//
func (s *SmartContract) createPaintlessRepair(APIstub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 1 {
		return shim.Error("Incorrect number of argument, expected 1")
	}
	var pr PaintlessRepair

	err := json.Unmarshal([]byte(args[0]), &pr)
	fmt.Println("pr", pr)
	if err != nil {
		return shim.Error("Input argument unmarshling error")
	}
	if len(pr.PRId) == 0 {
		return shim.Error(" Id is mandatory")
	}
	if len(pr.CustomerPhoneNo) == 0 {
		return shim.Error("CustomerPhoneNo is mandatory")
	}
	if len(pr.VINNumber) == 0 {
		return shim.Error("VINNumber is mandatory")
	}
	// if reflect.TypeOf(user.MobileNo).Kind() != reflect.Int {
	// 	return shim.Error("Mobile number is mandatory")
	// }	
	
	if recordBytes, _ := APIstub.GetState(pr.PRId); len(recordBytes) > 0 {
		return shim.Error("PrePaidAgreement id exists,Please provide unique id")
	}
	recordBytesb, _ := APIstub.GetState(pr.CustomerPhoneNo)
	if len(recordBytesb) == 0 {
		return shim.Error("Customer is not registered,Please give a registered id")
	}
	var prData = PaintlessRepair{
		DocType:"paintless repair",
		PRId:				 pr.PRId,
		CustomerName:  		 pr.CustomerName,
		CustomerAddress:  	 pr.CustomerAddress,
		CustomerPhoneNo:   	 pr.CustomerPhoneNo,
		VehicleYear:    	 pr.VehicleYear,
		VehicleName:   		 pr.VehicleName,
		VehicleModel:		 pr.VehicleModel,
		VehicleMileage:		 pr.VehicleMileage,	
		VINNumber:   		 pr.VINNumber,		
		DealersName:    	 pr.DealersName,
		DealersAddress:		 pr.DealersAddress,
		DealersPhonenumber:	 pr.DealersPhonenumber,			
		VehiclePrice:		 pr.VehiclePrice,
		Term:				 pr.Term,	
		Deductible:   	 	 pr.Deductible,
				
	}
	AsByte, _ := json.Marshal(prData)
	APIstub.PutState(pr.PRId, AsByte)
    dataByte := map[string]interface{}{
		"trxnID":  APIstub.GetTxID(),
		"status":  true,
		"message": "success",
		"data":    pr,
	}
    userJSON, _ := json.Marshal(dataByte)

	return shim.Success(userJSON)

}
func (s *SmartContract) queryPaintlessRepair(APIstub shim.ChaincodeStubInterface, args []string) peer.Response {

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}
	AsBytes, err := APIstub.GetState(args[0])
	if err !=nil{
		return shim.Error("Failed to read from world state  ")
	}
	if AsBytes == nil {
		return shim.Error("id does not exist")
	}
	return shim.Success(AsBytes)
}
// ****************Service Agreement*****************//
func (s *SmartContract) createServiceAgreement(APIstub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 1 {
		return shim.Error("Incorrect number of argument, expected 1")
	}
	var sa ServiceAgreement

	err := json.Unmarshal([]byte(args[0]), &sa)
	fmt.Println("sa", sa)
	if err != nil {
		return shim.Error("Input argument unmarshling error")
	}
	if len(sa.SId) == 0 {
		return shim.Error(" Id is mandatory")
	}
	if len(sa.CustomerPhoneNo) == 0 {
		return shim.Error("CustomerPhoneNo is mandatory")
	}
	if len(sa.VINNumber) == 0 {
		return shim.Error("VINNumber is mandatory")
	}
	// if reflect.TypeOf(user.MobileNo).Kind() != reflect.Int {
	// 	return shim.Error("Mobile number is mandatory")
	// }	
	
	if recordBytes, _ := APIstub.GetState(sa.SId); len(recordBytes) > 0 {
		return shim.Error("PrePaidAgreement id exists,Please provide unique id")
	}
	recordBytesb, _ := APIstub.GetState(sa.CustomerPhoneNo)
	if len(recordBytesb) == 0 {
		return shim.Error("Customer is not registered,Please give a registered id")
	}
	var saData = ServiceAgreement{
		DocType:"Service",
		SId: sa.SId,
		VehicleYear:  		 sa.VehicleYear,
		VehicleName:  	 sa.VehicleName,
		VehicleModel:   	 sa.VehicleModel,
		VINNumber:    	 sa.VINNumber,
		VehicleMileage:   		 sa.VehicleMileage,
		VehiclePrice:		 sa.VehiclePrice,
		CustomerName:		 sa.CustomerName,	
		CustomerAddress:   		 sa.CustomerAddress,		
		CustomerPhoneNo:    	sa.CustomerPhoneNo,
		DealersName:		 sa.DealersName,
		DealersAddress:	 sa.DealersAddress,		
		DealersPhonenumber:		  sa.DealersPhonenumber,
		LienHolder:				sa.LienHolder,	
		LAddress:   	 	sa.LAddress,
		AgreementPrice:   		 sa.AgreementPrice,		
		AgreementDate:    	sa.AgreementDate,
		Term:		 sa.Term,
		Coverage:	 sa.Coverage,		
		Deductible:		  sa.Deductible,
		Surcharges:				sa.Surcharges,	
		OptionalCoverage:   	 	sa.OptionalCoverage,
				
	}
	AsByte, _ := json.Marshal(saData)
	APIstub.PutState(sa.SId, AsByte)
    dataByte := map[string]interface{}{
		"trxnID":  APIstub.GetTxID(),
		"status":  true,
		"message": "success",
		"data":    sa,
	}
    userJSON, _ := json.Marshal(dataByte)

	return shim.Success(userJSON)

}
func (s *SmartContract) queryServiceAgreement(APIstub shim.ChaincodeStubInterface, args []string) peer.Response {

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}
	AsBytes, err := APIstub.GetState(args[0])
	if err !=nil{
		return shim.Error("Failed to read from world state  ")
	}
	if AsBytes == nil {
		return shim.Error("id does not exist")
	}
	return shim.Success(AsBytes)
}
// ****************Claim handling*****************//
func (s *SmartContract) createClaim(APIstub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 1 {
		return shim.Error("Incorrect number of argument, expected 1")
	}
	var claim Claim

	err := json.Unmarshal([]byte(args[0]), &claim)
	fmt.Println("claim", claim)
	if err != nil {
		return shim.Error("Input argument unmarshling error")
	}
	if len(claim.ContractNo) == 0 {
		return shim.Error(" Contract No is mandatory")
	}
	if len(claim.MobileNo) == 0 {
		return shim.Error("Customer Phone No is mandatory")
	}
	if len(claim.FName) == 0 {
		return shim.Error("First name is mandatory")
	}
	if len(claim.LName) == 0 {
		return shim.Error("Last name is mandatory")
	}
	// if reflect.TypeOf(user.MobileNo).Kind() != reflect.Int {
	// 	return shim.Error("Mobile number is mandatory")
	// }	
	
	if recordBytes, _ := APIstub.GetState(claim.ContractNo); len(recordBytes) > 0 {
		return shim.Error("Claim id exists,Please provide unique id")
	}
	recordBytesb, _ := APIstub.GetState(claim.MobileNo)
	if len(recordBytesb) == 0 {
		return shim.Error("Customer is not registered,Please give a registered id")
	}
	var cData = Claim{
		DocType:"Claim",
		ContractNo:		claim.ContractNo,
		FName:  		claim.FName,
		LName:			claim.LName,
		MobileNo:  		claim.MobileNo,
		Dealer:  	 	claim.Dealer,
		Address:   	 	claim.Address,
		EmailId:    	claim.EmailId,
		DateOfLoss:   	claim.DateOfLoss,
		Mileage:		claim.Mileage,
		Insurance:		claim.Insurance,	
							
	}
	AsByte, _ := json.Marshal(cData)
	APIstub.PutState(claim.ContractNo, AsByte)
    dataByte := map[string]interface{}{
		"trxnID":  APIstub.GetTxID(),
		"status":  true,
		"message": "success",
		"data":    cData,
	}
    userJSON, _ := json.Marshal(dataByte)

	return shim.Success(userJSON)

}
func (s *SmartContract) queryClaim(APIstub shim.ChaincodeStubInterface, args []string) peer.Response {

	if len(args) != 1 {
		return shim.Error("Incorrect number of arguments. Expecting 1")
	}
	AsBytes, err := APIstub.GetState(args[0])
	if err !=nil{
		return shim.Error("Failed to read from world state  ")
	}
	if AsBytes == nil {
		return shim.Error("id does not exist")
	}
	return shim.Success(AsBytes)
}

 func (s *SmartContract) createRefundRequest(APIstub shim.ChaincodeStubInterface, args []string) peer.Response {
	 if len(args) != 1 {
		 return shim.Error("Incorrect number of argument, expected 1")
	 }
	 var refund Refund
	 err := json.Unmarshal([]byte(args[0]), &refund)
	 // fmt.Println("rid", refund)
	 fmt.Println(args[0])

	 if err != nil {
		 return shim.Error("Input argument unmarshling error")
	 }
	 if len(refund.RID) == 0 {
		 return shim.Error("refund Id is mandatory")
	 }
	 if len(refund.ContractNo) == 0 {
		 return shim.Error("Contractno  is mandatory")
	 }
	 if len(refund.CName) == 0 {
		 return shim.Error("customer name is mandatory")
	 }
	 if len(refund.CancelDate) == 0 {
		 return shim.Error("cancel date is mandatory")
	 }
	 if recordBytes, _ := APIstub.GetState(refund.RID); len(recordBytes) > 0 {
		 return shim.Error("refund already registered. Provide an unique id")
	 }

	 var Years int = refund.Years
	 var nYears int = 12 * Years
	 var presentMonth float64 = refund.PresentMonth
	 diff := float64(1 - nYears)
	 var delta float64 = (diff / float64(nYears)) * 100
	 var percent float64 = 100 + delta
	 var contractPrice float64 = refund.AgreementPurchasePrice
	 var minusAmount float64 = contractPrice * presentMonth * percent / 100
	 fmt.Println(minusAmount)
	 var returnAmount float64 = (contractPrice - minusAmount - 50)
	 fmt.Println("percentage to be removed per month", percent)
	 fmt.Println("return Amount is ", returnAmount)
	 var rData = Refund{

		 DocType:                "refund",
		 RID:                    refund.RID,
		 ContractNo:             refund.ContractNo,
		 CName:                  refund.CName,
		 EffectiveDate:          refund.EffectiveDate,
		 StartTime:              refund.StartTime,
		 EndTime:                refund.EndTime,
		 AgreementPurchasePrice: refund.AgreementPurchasePrice,
		 CancelDate:             refund.CancelDate,
		 ApprovalFromAdmin:      refund.ApprovalFromAdmin,
		 ID:                     refund.ID,
		 Years:                  refund.Years,
		 PresentMonth:           refund.PresentMonth,
	 }

	 AsByte, _ := json.Marshal(rData)
	 APIstub.PutState(refund.RID, AsByte)
	 Data := map[string]interface{}{
		 "trxnID":  APIstub.GetTxID(),
		 "status":  true,
		 "message": "success",
		 "data":    returnAmount,
	 }
	 userJSON, _ := json.Marshal(Data)

	 return shim.Success(userJSON)
 }
 func (s *SmartContract) updateRefundRequest(stub shim.ChaincodeStubInterface, args []string) peer.Response {

	 if len(args) != 1 {
		 return shim.Error("Incorrect number of arguments. Expecting 1")
	 }
	 fmt.Println(args[0])
	 var refund Refund
	 err := json.Unmarshal([]byte(args[0]), &refund)
	 if err != nil {
		 return shim.Error("Input arguments unmarshling error")
	 }
	 fmt.Println("unmarshal refund details", refund)
	 AsByte, err := stub.GetState(refund.RID)

	 if err != nil {
		 return shim.Error("error while getting the user")
	 } else if len(AsByte) == 0 {
		 return shim.Error(" id : " + refund.RID + " does not exist")
	 }

	 AByte, err := stub.GetState(refund.ID)
	 if err != nil {
		 return shim.Error("error while getting the user")
	 } else if len(AsByte) == 0 {
		 return shim.Error(" id : " + refund.ID + " does not exist")
	 }
	 fmt.Println(AByte)

	 var update Refund
	 err = json.Unmarshal(AsByte, &update)
	 if err != nil {
		 return shim.Error("Existing user details unmarshaling error : " + string(err.Error()))
	 }
	 update.ApprovalFromAdmin = refund.ApprovalFromAdmin
	 update.ID = refund.ID

	 refundAsBytes, _ := json.Marshal(update)

	 err = stub.PutState(refund.RID, refundAsBytes)
	 if err != nil {
		 return shim.Error("Error while inserting the data into Ledger : " + err.Error())
	 }
	 finalData := map[string]interface{}{
		 "status":  "true",
		 "message": "successfully changed details",
		 "trxnID":  stub.GetTxID(),
		 "data":    refundAsBytes,
	 }
	 DataAsByte, err := json.Marshal(finalData)
	 return shim.Success(DataAsByte)
 }
 func (s *SmartContract) queryRefundRequest(APIstub shim.ChaincodeStubInterface, args []string) peer.Response {

	 if len(args) != 1 {
		 return shim.Error("Incorrect number of arguments. Expecting 1")
	 }

	 AsBytes, err := APIstub.GetState(args[0])
	 if err != nil {
		 return shim.Error("Failed to read from world state  ")
	 }
	 if AsBytes == nil {
		 return shim.Error("id does not exist")
	 }
	 return shim.Success(AsBytes)
 }

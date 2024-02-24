package main

import (
	"fmt"
	"github.com/hyperledger/fabric-chaincode-go/shim"
    "github.com/hyperledger/fabric-protos-go/peer"
)

// SmartContract is a structure
type SmartContract struct {
	scott *Scott
}

// var _mainLogger = shim.NewLogger("SupplychainSmartContract-Updated")

func main() {
	err := shim.Start(new(SmartContract))
	if err != nil {
		// _mainLogger.Criticalf("Error starting  chaincode: %v", err)
		fmt.Printf("Error starting Simple chaincode: %s", err)
	}
}
// ==========================================================
// Init initializes chaincode
// ==========================================================
func (s *SmartContract) Init(APIstub shim.ChaincodeStubInterface) peer.Response {

	return shim.Success(nil)
}
//Invoke function
//==============================
func (s *SmartContract) Invoke(APIstub shim.ChaincodeStubInterface) peer.Response {

	// Retrieve the requested Smart Contract function and arguments
	function, args := APIstub.GetFunctionAndParameters()
	// Route to the appropriate handler function to interact with the ledger appropriately
	if function == "createAdministration" {
		return s.createAdministration(APIstub, args)
	}else if function == "updateAdministration" {
			return s.updateAdministration(APIstub, args)
	}else if function == "queryAdministration" {
				return s.queryAdministration(APIstub, args)
	}else if function == "createUser" {
		return s.createUser(APIstub, args)
	}else if function == "queryUser" {
		return s.queryUser(APIstub, args)
	}else if function == "updateUser" {
		return s.updateUser(APIstub, args)
	}else if function == "userLogin" {
		return s.userLogin(APIstub, args)
	}else if function == "createClaim" {
		return s.createClaim(APIstub, args)
		}else if function == "queryClaim" {
			return s.queryClaim(APIstub, args)
	}else if function == "findByNumberLastName" {
		return s.findByNumberLastName(APIstub, args)
	}else if function == "findByContractLastName" {
		return s.findByContractLastName(APIstub, args)
	}else if function == "createPrePaidAgreement" {
		return s.createPrePaidAgreement(APIstub, args)
	}else if function == "queryPrepaidAgreement" {
		return s.queryPrepaidAgreement(APIstub, args)
	}else if function == "createBatteryAgreement" {
		return s.createBatteryAgreement(APIstub, args)
	}else if function == "queryBatteryAgreement" {
		return s.queryBatteryAgreement(APIstub, args)
	}else if function == "createPaintlessRepair" {
		return s.createPaintlessRepair(APIstub, args)
	}else if function == "queryPaintlessRepair" {
		return s.queryPaintlessRepair(APIstub, args)
	}else if function == "queryServiceAgreement" {
		return s.queryServiceAgreement(APIstub, args)
	}else if function == "createServiceAgreement" {
		return s.createServiceAgreement(APIstub, args)
	}else if function == "createRefundRequest" {
		return s.createRefundRequest(APIstub, args)
	}else if function == "updateRefundRequest" {
		return s.updateRefundRequest(APIstub, args)
	}else if function == "queryRefundRequest" {
		return s.queryRefundRequest(APIstub, args)
	}else if function == "queryClaim" {
		return s.queryClaim(APIstub, args)
	}

    return shim.Error("Invalid chaincode function name: " + function)
}


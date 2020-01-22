package main

import "github.com/hyperledger/fabric-chaincode-go/shim"

import "github.com/hyperledger/fabric-protos-go/peer"

import "fmt"

type KarnaChaincode struct {
	// Test bool
}

func (c *KarnaChaincode) Init(stub shim.ChaincodeStubInterface) peer.Response {
	// c.Test = true
	return peer.Response{
		Status:  200,
		Message: "successfully initiated",
		Payload: nil,
	}
}
func (c *KarnaChaincode) Invoke(stub shim.ChaincodeStubInterface) peer.Response {
	
	funName, args := stub.GetFunctionAndParameters()
	if funName == "setSecret" {
		return putSecret(stub, args)
	}
	if funName == "userRegistration" {
		return userRegistration(stub, args)
	}
	if funName == "NGORegistration" {
		return NGORegistration(stub, args)
	}
	if funName == "getUser" {
		return getUser(stub, args)
	}
	if funName == "getNgo" {
		return getNgo(stub, args)
	}
	if funName == "userGateway" {
		return userGateway(stub, args)
	}
	if funName == "ngoGateway" {
		return ngoGateway(stub, args)
	}
	if funName == "getAllMission" {
		return getAllMission(stub)
	}
	return peer.Response{
		Status:  200,
		Message: "successfully initiated",
		Payload: nil,
	}
}
func main() {
	err := shim.Start(new(KarnaChaincode))
	if err != nil {
		fmt.Println(err.Error())
	}
}

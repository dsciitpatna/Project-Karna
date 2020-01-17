package main

import (
	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/hyperledger/fabric-chaincode-go/shimtest"
	"testing"
)

func Test(t *testing.T) {
	stub := shimtest.NewMockStub("cc", new(KarnaChaincode))
	response := stub.MockInit("init", nil)
	t.Logf("Init status code %d", response.Status)
	if response.Status != shim.OK {
		t.FailNow()
	}
	ccargs := SetupArgsArray("userRegistration","zzocker","Pritam","pw")
	stub.MockInvoke("in",ccargs)
	ccargs =  SetupArgsArray("getUser","zzocker")
	response =  stub.MockInvoke("c",ccargs)
	t.Log(string(response.Payload))
	ccargs = SetupArgsArray("userGateway","userLogin","zzocker","pw")
	response =  stub.MockInvoke("c",ccargs)
	t.Log(string(response.Payload))
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJsb2dpbl9pZCI6Inp6b2NrZXIiLCJleHAiOjE1NzkyODI3NDJ9.Pj2SWrQAnRQoXRFUOAef8bpYz-AjwyV-KS5HjiuSKSU"
	ccargs = SetupArgsArray("userGateway",token)
	response =  stub.MockInvoke("f",ccargs)
	t.Log(string(response.Payload))
}

func SetupArgsArray(funcName string, args ...string) [][]byte {
	ccArgs := make([][]byte, 1+len(args))
	ccArgs[0] = []byte(funcName)
	for i, arg := range args {
		ccArgs[i+1] = []byte(arg)
	}

	return ccArgs
}
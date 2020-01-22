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
	ccargs := SetupArgsArray("userRegistration", "zzocker", "Pritam", "pw")
	stub.MockInvoke("in", ccargs)

	ccargs = SetupArgsArray("userGateway", "userLogin", "zzocker", "pw")
	response = stub.MockInvoke("c", ccargs)
	t.Log(string(response.Payload))
	token := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJsb2dpbl9pZCI6Inp6b2NrZXIiLCJleHAiOjE1Nzk3MDExMzl9.91dK8XsFiEN_Pv8jhRPwTMtJNIFRhl6Hmp5PgAMZDas"
	ccargs = SetupArgsArray("userGateway", token)
	response = stub.MockInvoke("f", ccargs)
	t.Log(string(response.Payload))
	ccargs = SetupArgsArray("NGORegistration","ngo1","NGO1","IITP","Nss","pw")
	stub.MockInvoke("f", ccargs)
	ccargs = SetupArgsArray("ngoGateway","userLogin","ngo1","pw")
	response = stub.MockInvoke("f", ccargs)
	t.Log(string(response.Payload))
	token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJsb2dpbl9pZCI6Im5nbzEiLCJleHAiOjE1Nzk3MDExMzl9.w9KCx_7KL3-pqoMXdyrJUMjvzfAHFHy9_KAb0BFBzlQ"
	ccargs = SetupArgsArray("ngoGateway",token)
	response = stub.MockInvoke("f", ccargs)
	t.Log(string(response.Payload))
	ccargs = SetupArgsArray("ngoGateway",token,"createMission","first","FirstMission","hello no dec","5000")
	response = stub.MockInvoke("f", ccargs)
	t.Log(string(response.Payload))
	token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJsb2dpbl9pZCI6Inp6b2NrZXIiLCJleHAiOjE1Nzk3MDExMzl9.91dK8XsFiEN_Pv8jhRPwTMtJNIFRhl6Hmp5PgAMZDas"
	ccargs = SetupArgsArray("userGateway",token,"donate","ngo1","first","50")
	response = stub.MockInvoke("f", ccargs)
	t.Log(string(response.Payload))
	ccargs = SetupArgsArray("getAllMission")
	response = stub.MockInvoke("f", ccargs)
	t.Log(string(response.Payload))
	token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJsb2dpbl9pZCI6Im5nbzEiLCJleHAiOjE1Nzk3MDExMzl9.w9KCx_7KL3-pqoMXdyrJUMjvzfAHFHy9_KAb0BFBzlQ"
	ccargs = SetupArgsArray("ngoGateway",token,"createMission","second","FirstMission","hello no dec","5000")
	response = stub.MockInvoke("f", ccargs)
	t.Log(string(response.Payload))
	ccargs = SetupArgsArray("getAllMission")
	response = stub.MockInvoke("f", ccargs)
	t.Log(string(response.Payload))

	token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJsb2dpbl9pZCI6Inp6b2NrZXIiLCJleHAiOjE1Nzk3MDExMzl9.91dK8XsFiEN_Pv8jhRPwTMtJNIFRhl6Hmp5PgAMZDas"
	ccargs = SetupArgsArray("userGateway",token,"getDonatedMission")
	response = stub.MockInvoke("f", ccargs)
	t.Log(string(response.Payload))
	ccargs = SetupArgsArray("getNgo","ngo1")
	response = stub.MockInvoke("f", ccargs)
	t.Log(string(response.Payload))

	token = "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJsb2dpbl9pZCI6Im5nbzEiLCJleHAiOjE1Nzk3MDExMzl9.w9KCx_7KL3-pqoMXdyrJUMjvzfAHFHy9_KAb0BFBzlQ"
	ccargs = SetupArgsArray("ngoGateway",token,"getNgoMission")
	response = stub.MockInvoke("f", ccargs)
	t.Log(string(response.Payload))
	// ccargs = SetupArgsArray("ngoGateway","userLogin","ngo1","pw")
	// response = stub.MockInvoke("f", ccargs)
	// t.Log(string(response.Payload))
	// token ="eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJsb2dpbl9pZCI6Im5nbzEiLCJleHAiOjE1NzkzODY4NjR9.lV7kaiE8JuTrI45VbFylgkh7AxFjfEO6FDA0S5aaroQ"
	// ccargs = SetupArgsArray("userGateway", token)
	// response = stub.MockInvoke("f", ccargs)
	// t.Log(string(response.Message))
}

func SetupArgsArray(funcName string, args ...string) [][]byte {
	ccArgs := make([][]byte, 1+len(args))
	ccArgs[0] = []byte(funcName)
	for i, arg := range args {
		ccArgs[i+1] = []byte(arg)
	}

	return ccArgs
}

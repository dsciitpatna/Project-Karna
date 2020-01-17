package main

import (
	"github.com/hyperledger/fabric-chaincode-go/shim"
	"encoding/json"
	"github.com/hyperledger/fabric-protos-go/peer"
)

func userGateway(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args)<1{
		return shim.Error("Require more than 1 args")
	}
	adminkey := getSecretKey(stub)
	action:= args[0]
	newargs:= args[1:]
	if action=="userLogin" {
		return userLogin(stub,newargs,adminkey)
	} else{
		is := authToken(action,adminkey)
		if is{
			return shim.Success([]byte("Welcome!!"))
		}
	}	
	return shim.Error("Wrong Password!!")
}

func userRegistration(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	// args[0] : loginid , args[1]: Name , args[2]: password
	if len(args) != 3 {
		return shim.Error("Require 3 arguments")
	}
	userkey := getUserKey(stub, args[0])
	_, is, err := getState(stub, userkey)
	if is {
		return shim.Error("User already exists")
	}
	user := &User{
		Name:    args[1],
		Balance: 0,
	}
	err = stub.PutState(userkey, getMarshaled(user))
	if err != nil {
		return shim.Error(err.Error())
	}
	login := &Login{
		LoginId:  args[0],
		Password: args[2],
	}
	loginKey := getLoginKey(stub, args[0])
	err = stub.PutState(loginKey, getMarshaled(login))
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(nil)
}
func userLogin(stub shim.ChaincodeStubInterface, args []string,secret []byte) peer.Response {
	if len(args) != 2 {
		return shim.Error("require 2 args to login")
	}
	key := getLoginKey(stub, args[0])
	userByte, is, err := getState(stub, key)
	if !is {
		shim.Error(err.Error())
	}
	var user Login
	_= json.Unmarshal(userByte,&user)
	if user.Password!=args[1]{
		return shim.Error("Wrong Password!!")
	}
	token,err:= createJWT(args[0],secret)
	if err!=nil {
		shim.Error(err.Error())
	}
	return shim.Success([]byte(token))
}
func putSecret(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 1 {
		return shim.Error("Require 1 args to put admin secret")
	}
	err := stub.PutState("adminkey", []byte(args[0]))
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(nil)
}
func getSecretKey(stub shim.ChaincodeStubInterface) []byte  {
	key, _ := stub.GetState("adminkey")
	return key
}
func getUserTest(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) != 1 {
		return shim.Error("Need 1 arg")
	}
	key := getUserKey(stub, args[0])
	user, is, err := getState(stub, key)
	if !is {
		return shim.Error(err.Error())
	}
	return shim.Success(user)
}

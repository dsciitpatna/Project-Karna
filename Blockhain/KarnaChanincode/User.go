package main

import (
	"encoding/json"
	"fmt"
	"github.com/hyperledger/fabric-chaincode-go/shim"
	"github.com/hyperledger/fabric-protos-go/peer"
	"strconv"
)

func userGateway(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) < 1 {
		return shim.Error("Require more than 1 args")
	}
	adminkey := getSecretKey(stub)
	action := args[0]
	newargs := args[1:]
	if action == "userLogin" {
		return userLogin(stub, newargs, adminkey)
	} else {
		is, loginid := authToken(action, adminkey)
		if !is {
			return shim.Error("Wrong Token")
		}
		key := getUserKey(stub, loginid)
		byteUser, is, err := getState(stub, key)
		if !is {
			return shim.Error(err.Error())
		}
		var user User
		json.Unmarshal(byteUser, &user)
		if len(newargs) > 0 {
			/*1*/ if newargs[0] == "donate" {
				err := donate(stub, &user, newargs[1:])
				if err != nil {
					return shim.Error(err.Error())
				}
			}

			err = stub.PutState(key, getMarshaled(user))
			if err != nil {
				return shim.Error(err.Error())
			}
		}
		return shim.Success(getMarshaled(user))
	}
}
func donate(stub shim.ChaincodeStubInterface, user *User, args []string) error {
	// args[0]=ngo_id,args[1]=mission_name,args[2]=donation
	if len(args) != 3 {
		return fmt.Errorf("require 3 args for donation")
	}
	key := getMissionKey(stub, args[0], args[1])
	byteM, is, err := getState(stub, key)
	if !is {
		return err
	}
	var mission Mission
	json.Unmarshal(byteM, &mission)
	money, err := strconv.ParseInt(args[2], 10, 64)
	if err!=nil{
		return nil
	}
	mission.Total += money
	mission.Donation[user.Username] += money
	err = stub.PutState(key,getMarshaled(mission))
	if err!=nil{
		return err
	}
	user.Donation[key] += money
	return nil
}
func userRegistration(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	// args[0] : loginid , args[1]: Name , args[2]: password
	if len(args) != 3 {
		return shim.Error("Require 3 arguments")
	}
	userkey := getUserKey(stub, args[0])
	_, is, _ := getState(stub, userkey)
	if is {
		return shim.Error("User already exists")
	}
	user := &User{
		Username: args[0],
		Name:     args[1],
		Donation: map[string]int64{},
	}
	err := stub.PutState(userkey, getMarshaled(user))
	if err != nil {
		return shim.Error(err.Error())
	}
	return addLogin(stub, args[0], args[2])
}
func NGORegistration(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	// args[0] = loginid,args[1] : name , args[2]=Address , args[3] = Description ,args[4]= password
	if len(args) != 5 {
		return shim.Error("Require 4 arguments")
	}
	ngoKey := getNGOKey(stub, args[0])
	_, is, _ := getState(stub, ngoKey)
	if is {
		return shim.Error("NGO Already exists")
	}
	ngo := &NGO{
		Username:    args[0],
		Name:        args[1],
		Address:     args[2],
		Description: args[3],
		Missions:    make(map[string]int64),
	}
	err := stub.PutState(ngoKey, getMarshaled(ngo))
	if err != nil {
		return shim.Error(err.Error())
	}
	return addLogin(stub, args[0], args[4])
}
func userLogin(stub shim.ChaincodeStubInterface, args []string, secret []byte) peer.Response {
	if len(args) != 2 {
		return shim.Error("require 2 args to login")
	}
	key := getLoginKey(stub, args[0])
	userByte, is, err := getState(stub, key)
	if !is {
		shim.Error(err.Error())
	}
	var user Login
	_ = json.Unmarshal(userByte, &user)
	if user.Password != args[1] {
		return shim.Error("Wrong Password!!")
	}
	token, err := createJWT(args[0], secret)
	if err != nil {
		shim.Error(err.Error())
	}
	return shim.Success([]byte(token))
}
func addLogin(stub shim.ChaincodeStubInterface, loginid, password string) peer.Response {
	login := &Login{
		LoginId:  loginid,
		Password: password,
	}
	loginKey := getLoginKey(stub, loginid)
	err := stub.PutState(loginKey, getMarshaled(login))
	if err != nil {
		return shim.Error(err.Error())
	}
	return shim.Success(nil)
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
func getSecretKey(stub shim.ChaincodeStubInterface) []byte {
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

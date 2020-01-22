package main

import (
	"github.com/hyperledger/fabric-chaincode-go/shim"
	"fmt"
	"encoding/json"
	"strconv"
	"github.com/hyperledger/fabric-protos-go/peer"
)

func ngoGateway(stub shim.ChaincodeStubInterface, args []string) peer.Response {
	if len(args) < 1 {
		return shim.Error("Require more than 1 args")
	}
	adminkey := getSecretKey(stub)
	action := args[0]
	newargs := args[1:]
	if action == "userLogin" {
		return userLogin(stub, newargs, adminkey)
	} else {
		is ,loginid:= authToken(action, adminkey)
		if !is {
			return shim.Error("Wrong Token")
		}
		key:= getNGOKey(stub,loginid)
		bytengo,is,err:= getState(stub,key)
		if !is{
			return shim.Error(err.Error())
		}
		//////////////////////////////////////////
		var ngo NGO
		json.Unmarshal(bytengo,&ngo)
		if len(newargs)>0{
		/*1*/if newargs[0]=="createMission"{
			err=createMission(stub,&ngo,newargs[1:])
			if err!=nil{
				return shim.Error(err.Error())
			}
			err = stub.PutState(key,getMarshaled(ngo))
			if err!=nil{
			return shim.Error(err.Error())
			}}
			/*2*/ if newargs[0]=="getNgoMission"{
				return getNgoMission(stub,ngo.Missions,ngo.Username)
			}
		}
		
		// return shim.Success(nil)
		return shim.Success(getMarshaled(ngo))
	}
}
func createMission(stub shim.ChaincodeStubInterface,ngo *NGO,args []string) (error) {
	// args[0]=mission_id,args[1]=name,args[2]=description,args[3]=target
	if len(args)!=4{
		return fmt.Errorf("require 4 args for creating the mission")
	}
	key := getMissionKey(stub,ngo.Username,args[0])
	_,is,_:= getState(stub,key)
	if is{
		return fmt.Errorf("Mission already exists")
	}
	target,err:= strconv.ParseInt(args[3],10,64)
	if err!=nil{
		return err
	}
	mission := &Mission{
		MissionID: args[0],
		Name: args[1],
		Description: args[2],
		Target: target,
		Total: 0,
		Donation: make(map[string]int64),
	}
	err= stub.PutState(key,getMarshaled(mission))
	if err!=nil{
		return err
	}
	ngo.Missions[args[0]]=target
	return nil
}
func getNgo(stub shim.ChaincodeStubInterface,args []string) peer.Response{
	if len(args)!=1{
		return shim.Error("NGO user ID required")
	}
	key := getNGOKey(stub,args[0])
	ngoByte,err:= stub.GetState(key)
	if err!=nil{
		return shim.Error(err.Error())
	}
	return shim.Success(ngoByte)
}
func getNgoMission(stub shim.ChaincodeStubInterface,mission map[string]int64,username string) peer.Response{
	result := []struct {
		MDetails Mission `json:"mission_details"`
		Target int64 `json:"target"`
	}{}
	for k,v:= range mission{
		key := getMissionKey(stub,username,k)
		temp,err:= getMission(stub,key)
		if err!=nil{
			return shim.Error(err.Error())
		}
		singleResult := struct{
			MDetails Mission `json:"mission_details"`
			Target int64 `json:"target"`
		}{temp,v}
		result = append(result,singleResult)
	}
	output,_:= json.Marshal(result)
	return shim.Success(output)
}
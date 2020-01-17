package main

import (
	"github.com/hyperledger/fabric-protos-go/msp"
	"github.com/golang/protobuf/proto"
)
const ADMINMSP="GovernmentMSP"
func isAdmin(creator []byte) (bool,error){
	creatorSerializedID := &msp.SerializedIdentity{}
	err := proto.Unmarshal(creator,creatorSerializedID)
	if err!=nil{
		return false, err
	}
	return creatorSerializedID.GetMspid()==ADMINMSP,nil
}

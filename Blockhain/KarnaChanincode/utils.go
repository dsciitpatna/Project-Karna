package main

import (
	"encoding/json"
	"fmt"
	"time"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/hyperledger/fabric-chaincode-go/shim"
)

func getState(stub shim.ChaincodeStubInterface, key string) ([]byte, bool, error) {
	byteState, err := stub.GetState(key)
	if err != nil {
		return nil, false, err
	}
	if len(byteState) == 0 {
		return nil, false, fmt.Errorf("State doesn't exists")
	}
	return byteState, true, nil
}
func getUserKey(stub shim.ChaincodeStubInterface, loginid string) string {
	key, err := stub.CreateCompositeKey(USERKEY, []string{loginid})
	if err != nil {
		return ""
	}
	return key
}
func getLoginKey(stub shim.ChaincodeStubInterface, loginid string) string {
	key, err := stub.CreateCompositeKey(LOGINKEY, []string{loginid})
	if err != nil {
		return ""
	}
	return key
}
func getMarshaled(object interface{}) []byte {
	byteO, _ := json.Marshal(object)
	return byteO
}
func createJWT(loginid string, secret []byte) (string, error) {
	claim := &Claims{
		LoginID: loginid,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(2*time.Minute).Unix(),
		},
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256,claim)
	tokenString,err:=token.SignedString(secret)
	if err!=nil{
		return "",err
	}
	return tokenString,nil
}
func authToken(tokenString string, secret []byte) bool {
	c:=&Claims{}
	tkn,_:= jwt.ParseWithClaims(tokenString,c,func(token *jwt.Token) (interface{}, error) {
		return secret, nil
	})
	return tkn.Valid
}

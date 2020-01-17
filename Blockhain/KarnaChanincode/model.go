package main

import "github.com/dgrijalva/jwt-go"

const USERKEY = "USER"
const LOGINKEY = "LOGIN"

type Login struct {
	LoginId  string `json:"login_id"`
	Password string `json:"password"`
}
type User struct {
	Name    string `json:"name"`
	Balance int64  `json:"balance"`
}
type Claims struct {
	LoginID string `json:"login_id"`
	jwt.StandardClaims
}

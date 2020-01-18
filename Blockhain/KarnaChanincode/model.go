package main

import "github.com/dgrijalva/jwt-go"

const USERKEY = "USER"
const NGOKEY = "NGO"
const MISSIONSKEY = "MISSION"   // MISSION~NGOID~MISSIONID (Mission id will be given by NGO it should be unique)
const LOGINKEY = "LOGIN"

type Login struct {
	LoginId  string `json:"login_id"`
	Password string `json:"password"`
}
type User struct {
	Username string	`json:"login_id"`
	Name    string `json:"name"`
	Donation map[string]int64   `json:"donation"`
}
type Claims struct {
	LoginID string `json:"login_id"`
	jwt.StandardClaims
}
type NGO struct{
	Username string `json:"login_id"`
	Name string `json:"name"`
	Address string `json:"address"`
	Missions map[string]int64 `json:"missions"` // Missions map[key of project]Money Target
	Description string `json:"description"`
}
type Mission struct{
	MissionID string `json:"mission_id"`
	Name string `json:"mission_name"`
	Description string `json:"description"`
	Target int64 `json:"target"`
	Total int64 `json:"collected"`
	Donation map[string]int64  `json:"donors"`
}

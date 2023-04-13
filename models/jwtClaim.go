package models

import jwt "github.com/dgrijalva/jwt-go"


type JWTClaims struct {
	User 	string 	`json:"user"`
	ID 		string 	`json:"ID"`
	jwt.StandardClaims 
}
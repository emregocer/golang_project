package model

import "github.com/golang-jwt/jwt"

type Claims struct {
	jwt.StandardClaims
	UserID int `json:"id"`
}

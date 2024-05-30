package auth

import "github.com/golang-jwt/jwt/v5"

type Claims struct {
	ID int `json:"id"`
	jwt.RegisteredClaims
}

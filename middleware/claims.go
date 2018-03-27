package middleware

import jwt "github.com/dgrijalva/jwt-go"

type claims struct {
	Access []string `json:"access"`
	jwt.StandardClaims
}

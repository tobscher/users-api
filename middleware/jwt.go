package middleware

import (
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"strings"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/tobscher/users-api/core"
)

var (
	errMissingAuthorizationHeader = errors.New("Missing 'Authorization' header")
	errInvalidAuthorizationHeader = errors.New("'Authorization' header is invalid")
)

// JWT is a middleware to authenticate a user using a JSON Web Token
func JWT(h http.Handler, verifyKey []byte, requiredClaim string) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		var token string
		tokens, ok := r.Header["Authorization"]
		if ok && len(tokens) >= 1 {
			token = tokens[0]
			token = strings.TrimPrefix(token, "Bearer ")
		}

		if token == "" {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(core.NewErrorResponse(errMissingAuthorizationHeader))
			return
		}

		var claims claims
		_, err := jwt.ParseWithClaims(token, &claims, func(token *jwt.Token) (interface{}, error) {
			if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
				msg := fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
				return nil, msg
			}
			return verifyKey, nil
		})

		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(core.NewErrorResponse(errInvalidAuthorizationHeader))
			return
		}

		if !hasAccess(claims, requiredClaim) {
			w.WriteHeader(http.StatusUnauthorized)
			json.NewEncoder(w).Encode(core.NewErrorResponse(errInvalidAuthorizationHeader))
			return
		}

		h.ServeHTTP(w, r)
	})
}

func hasAccess(claims claims, key string) bool {
	for _, a := range claims.Access {
		if a == key {
			return true
		}
	}
	return false
}

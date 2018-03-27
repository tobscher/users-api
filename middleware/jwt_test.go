package middleware

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httptest"
	"testing"

	jwt "github.com/dgrijalva/jwt-go"
	"github.com/stretchr/testify/assert"
)

func createJwtToken(access []string, key string) string {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"access": access,
	})
	tokenString, err := token.SignedString([]byte(key))
	if err != nil {
		log.Fatal(err)
	}
	return tokenString
}

func TestJWTNoAuthorizationHeader(t *testing.T) {
	req := httptest.NewRequest("GET", "http://example.com/foo", nil)
	w := httptest.NewRecorder()

	stubHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	})

	JWT(stubHandler, []byte("key"), "user").ServeHTTP(w, req)

	resp := w.Result()
	assert.Equal(t, http.StatusUnauthorized, resp.StatusCode, "expected status code to be forbidden")
}

func TestJWTMalformedAuthorizationHeader(t *testing.T) {
	req := httptest.NewRequest("GET", "http://example.com/foo", nil)
	req.Header.Set("Authorization", "foo")
	w := httptest.NewRecorder()

	stubHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	})

	JWT(stubHandler, []byte("key"), "user").ServeHTTP(w, req)

	resp := w.Result()
	assert.Equal(t, http.StatusUnauthorized, resp.StatusCode, "expected status code to be forbidden")
}

func TestJWTInvalidAuthorizationHeader(t *testing.T) {
	req := httptest.NewRequest("GET", "http://example.com/foo", nil)
	invalidToken := createJwtToken([]string{"user"}, "invalid")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", invalidToken))
	w := httptest.NewRecorder()

	stubHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	})

	JWT(stubHandler, []byte("valid"), "user").ServeHTTP(w, req)

	resp := w.Result()
	assert.Equal(t, http.StatusUnauthorized, resp.StatusCode, "expected status code to be forbidden")
}

func TestJWTValidAuthorizationHeaderWithMissingClaim(t *testing.T) {
	req := httptest.NewRequest("GET", "http://example.com/foo", nil)
	validToken := createJwtToken([]string{"foo", "bar"}, "valid")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", validToken))
	w := httptest.NewRecorder()

	stubHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	})

	JWT(stubHandler, []byte("valid"), "user").ServeHTTP(w, req)

	resp := w.Result()
	assert.Equal(t, http.StatusUnauthorized, resp.StatusCode, "expected status code to be ok")
}

func TestJWTValidAuthorizationHeader(t *testing.T) {
	req := httptest.NewRequest("GET", "http://example.com/foo", nil)
	validToken := createJwtToken([]string{"user"}, "valid")
	req.Header.Set("Authorization", fmt.Sprintf("Bearer %s", validToken))
	w := httptest.NewRecorder()

	stubHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	})

	JWT(stubHandler, []byte("valid"), "user").ServeHTTP(w, req)

	resp := w.Result()
	assert.Equal(t, http.StatusOK, resp.StatusCode, "expected status code to be ok")
}

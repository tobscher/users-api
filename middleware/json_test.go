package middleware

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestContentTypeJSON(t *testing.T) {
	req := httptest.NewRequest("GET", "http://example.com/foo", nil)
	w := httptest.NewRecorder()

	stubHandler := http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {

	})

	ContentType(stubHandler, "application/json; charset=UTF-8").ServeHTTP(w, req)

	resp := w.Result()
	assert.Equal(t, "application/json; charset=UTF-8", resp.Header.Get("Content-Type"), "expected content type to be json")
}

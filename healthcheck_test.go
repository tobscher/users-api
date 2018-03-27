package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestHealthcheck(t *testing.T) {
	req := httptest.NewRequest("GET", "http://example.com/foo", nil)
	w := httptest.NewRecorder()
	healthcheckHandler(w, req)

	resp := w.Result()
	assert.Equal(t, resp.StatusCode, http.StatusOK, "expected response status code to be 200")
}

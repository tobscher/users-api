package handlers

import "net/http"

// healthcheckHandler always returns 200 and is used to indicate to the load balancer that
// the service is up and running.
//
// GET /health
func healthcheckHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}

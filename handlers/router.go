package handlers

import (
	"github.com/gorilla/mux"
)

// NewRouter returns a configured router.
func NewRouter(users *users) *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/health", healthcheckHandler).Methods("GET")
	router.HandleFunc("/users", users.index).Methods("GET")
	router.HandleFunc("/users/{id}", users.show).Methods("GET")
	return router
}

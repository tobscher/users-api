package handlers

import (
	"github.com/gorilla/mux"
	"github.com/tobscher/users-api/middleware"
)

const defaultContentType = "application/json; charset=UTF-8"

// NewRouter returns a configured router.
func NewRouter(users *users) *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/health", healthcheckHandler).Methods("GET")
	router.Handle("/users", middleware.ContentType(users.index(), defaultContentType)).Methods("GET")
	router.Handle("/users/{id}", middleware.ContentType(users.show(), defaultContentType)).Methods("GET")
	return router
}

package handlers

import (
	"github.com/gorilla/mux"
	"github.com/tobscher/users-api/middleware"
)

const defaultContentType = "application/json; charset=UTF-8"

// NewRouter returns a configured router.
func NewRouter(users *users, key []byte) *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/health", healthcheckHandler).Methods("GET")
	router.Handle("/users", middleware.ContentType(middleware.JWT(users.index(), key, "user"), defaultContentType)).Methods("GET")
	router.Handle("/users/{id}", middleware.ContentType(middleware.JWT(users.show(), key, "user"), defaultContentType)).Methods("GET")
	return router
}

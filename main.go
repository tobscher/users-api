package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/tobscher/users-api/handlers"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/health", handlers.HealthcheckHandler).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", router))
}

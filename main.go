package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/health", healthcheckHandler).Methods("GET")

	log.Fatal(http.ListenAndServe(":8080", router))
}

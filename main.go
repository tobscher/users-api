package main

import (
	"log"
	"net/http"

	"github.com/tobscher/users-api/handlers"
)

func main() {
	router := handlers.NewRouter(nil, nil)

	log.Fatal(http.ListenAndServe(":8080", router))
}

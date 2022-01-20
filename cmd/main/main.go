package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/manpreet1130/bookstore/pkg/routes"
)

func main() {
	// Creating a new router using the gorilla/mux package
	router := mux.NewRouter()
	
	// Calling CreateRoutes function to instantiate all required routes
	routes.CreateRoutes(router)
	http.Handle("/", router)
	
	// Creating server instance and listening on port 8080
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal(err)
	}
}

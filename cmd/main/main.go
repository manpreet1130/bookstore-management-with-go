package main

import (
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/manpreet1130/bookstore/pkg/routes"
)

func main() {
	router := mux.NewRouter()
	routes.CreateRoutes(router)
	http.Handle("/", router)
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		log.Fatal(err)
	}
}

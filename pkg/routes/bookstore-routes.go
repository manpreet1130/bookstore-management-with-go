// routes.go used to instantiate all the routes required

package routes

import (
	"github.com/gorilla/mux"
	"github.com/manpreet1130/bookstore/pkg/controllers"
)

// CreateRoutes takes in a reference to router and creates all the APIs with their corresponding methods
func CreateRoutes(router *mux.Router) {
	router.HandleFunc("/book/", controllers.CreateBook).Methods("POST")
	router.HandleFunc("/book/", controllers.GetAllBooks).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.GetBookById).Methods("GET")
	router.HandleFunc("/book/{bookId}", controllers.DeleteBook).Methods("DELETE")
	router.HandleFunc("/book/{bookId}", controllers.UpdateBook).Methods("PUT")
}

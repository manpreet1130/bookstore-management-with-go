// controller.go responds to the user input and performs interactions on the data model objects.
// it receives the input, optionally validates it and then passes the input to the model.

package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/manpreet1130/bookstore/pkg/models"
	"github.com/manpreet1130/bookstore/pkg/utils"
)

// GetAllBooks takes in an io.Writer and reference to an io.Reader
// and returns all books present in the database in JSON
func GetAllBooks(w http.ResponseWriter, r *http.Request) {
	books := models.GetAllBooks()
	res, _ := json.Marshal(books)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// GetBookById retrieves the id from the request and returns the
// book details for that particular id as response
func GetBookById(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["bookId"]
	bookId, err := strconv.ParseInt(id, 0, 0)
	if err != nil {
		fmt.Println("Cannot parse int")
	}
	bookDetails, _ := models.GetBookById(bookId)
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// DeleteBook deletes book with a given id from the database
func DeleteBook(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["bookId"]
	bookId, err := strconv.ParseInt(id, 0, 0)
	if err != nil {
		fmt.Println("Cannot parse int")
	}
	bookDetails := models.DeleteBook(bookId)
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// UpdateBook gets the book to be updated from the database by
// referring to its id, performs the updates on the data as required,
// saves the book back in the database and gives the updated book details as response
func UpdateBook(w http.ResponseWriter, r *http.Request) {
	var newBook = &models.Book{}
	utils.ParseBody(newBook, r)
	vars := mux.Vars(r)
	id := vars["bookId"]
	bookId, err := strconv.ParseInt(id, 0, 0)
	if err != nil {
		fmt.Println("Cannot parse int")
	}

	oldBook, db := models.GetBookById(bookId)

	if newBook.Name != "" {
		oldBook.Name = newBook.Name
	}

	if newBook.Author != "" {
		oldBook.Author = newBook.Author
	}

	if newBook.Publication != "" {
		oldBook.Publication = newBook.Publication
	}

	db.Save(&oldBook)
	res, _ := json.Marshal(oldBook)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// CreateBook is used to add a new book into the database and 
// returns the created book details as response
func CreateBook(w http.ResponseWriter, r *http.Request) {
	var book = &models.Book{}
	utils.ParseBody(book, r)
	b := book.CreateBook()
	res, _ := json.Marshal(b)
	w.Header().Set("Content-Type", "pkglication/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

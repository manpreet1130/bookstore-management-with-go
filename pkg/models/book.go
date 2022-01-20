// models.go is responsible for managing the data, logic and rules for the application

package models

import (
	"github.com/jinzhu/gorm"
	"github.com/manpreet1130/bookstore/pkg/config"
)

var db *gorm.DB

// Book struct comprises of the name, author and publication of any book
type Book struct {
	gorm.Model
	Name        string `json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}


// Used to create/open the database and create a table for Book if one doesn't already exist
func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

// CreateBook takes as input a Book, adds it to the database and returns pointer to Book
func (b *Book) CreateBook() *Book {
	db.Create(b)
	return b
}

// GetAllBooks returns a list of all books present in the database
func GetAllBooks() []Book {
	var books []Book
	db.Find(&books)
	return books
}

// GetBookById takes as input the id of a particular book and returns a reference to that book and the database
func GetBookById(id int64) (*Book, *gorm.DB) {
	var book Book
	db.Where("ID=?", id).Find(&book)
	return &book, db
}

// DeleteBook takes as input the id and deletes that particular book from the database
func DeleteBook(id int64) Book {
	var book Book
	db.Where("ID=?", id).Delete(&book)
	return book
}

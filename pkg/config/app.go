// config.go is responsible of instantiating a database connection by calling the Connect()
// and sending a reference to that database using the GetDB()

package config

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/sqlite"
)

var db *gorm.DB


// Connect is used to connect with a database using the gorm package
func Connect() {
	d, err := gorm.Open("sqlite3", "bookstore.db")
	if err != nil {
		log.Fatal(err)
	}

	db = d
}

// GetDB returns a reference to the database created using Connect
func GetDB() *gorm.DB {
	return db
}

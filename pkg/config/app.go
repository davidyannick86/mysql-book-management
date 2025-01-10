package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// * Global variable to hold the database connection
var (
	db *gorm.DB
)

func Connect() {
	// * Connect to the database
	d, err := gorm.Open("mysql", "root:root@/simplerest?charset=utf8&parseTime=True&loc=Local")

	if err != nil {
		panic("Failed to connect to database!")
	}

	// * Assign the connection to the global variable
	db = d
}

// * Get the database connection
func GetDB() *gorm.DB {
	return db
}

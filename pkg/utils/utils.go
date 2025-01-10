package utils

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"

	"github.com/jinzhu/gorm"
)

// * Parse the request body
func ParseBody(r *http.Request, model interface{}) error {
	// * Read the request body
	body, err := io.ReadAll(r.Body)
	if err != nil {
		return err
	}

	// * Unmarshal the request body
	if err := json.Unmarshal([]byte(body), model); err != nil {
		return err
	}
	return nil
}

func TestDBConnection() {
	// * Define the connection string
	dsn := "root:password@tcp(localhost:3306)/mysqldb?charset=utf8mb4&parseTime=True&loc=Local"

	// * Open the connection
	db, err := gorm.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("Could not connect to the database: %v", err)
	}
	defer db.Close()

	// * Test the connection
	err = db.DB().Ping()
	if err != nil {
		log.Fatalf("Could not ping the database: %v", err)
	}

	fmt.Println("Successfully connected to the database!")
}

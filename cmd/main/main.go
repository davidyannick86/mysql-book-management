package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/davidyannick86/mysql-book-management/pkg/routes"
	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

// * Test the database connection
func testDBConnection() {
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

func main() {
	// * Create a new router
	r := mux.NewRouter()

	// * Register the routes
	routes.RegisterBookStoreRoutes(r)

	// * Start the server
	http.Handle("/", r)

	// * Listen and serve
	log.Fatal(http.ListenAndServe("0.0.0.0:9010", nil))
	//testDBConnection()
}

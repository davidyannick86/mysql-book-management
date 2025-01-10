package main

import (
	"log"
	"net/http"

	"github.com/davidyannick86/mysql-book-management/pkg/routes"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	// * Create a new router
	r := mux.NewRouter()

	// * Register the routes
	routes.RegisterBookStoreRoutes(r)

	// * Start the server
	http.Handle("/", r)

	// * Listen and serve
	log.Fatal(http.ListenAndServe("0.0.0.0:9010", nil))
	//utils.TestDBConnection()
}

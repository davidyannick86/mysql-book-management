package main

import (
	"log"
	"net/http"

	"github.com/davidyannick86/mysql-book-management/pkg/routes"
	"github.com/gorilla/mux"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

func main() {
	r := mux.NewRouter()
	routes.RegisterBookStoreRoutes(r)
	http.Handle("/", r)
	log.Fatal(http.ListenAndServe("0.0.0.0:9010", r))
}

package controllers

import (
	"encoding/json"
	"net/http"

	"github.com/davidyannick86/mysql-book-management/pkg/models"
)

var NewBook models.Book

func GetBooks(w http.ResponseWriter, r *http.Request) {
	newBooks := models.GetAllBooks()
	res, _ := json.Marshal(newBooks)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

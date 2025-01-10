package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/davidyannick86/mysql-book-management/pkg/models"
	"github.com/davidyannick86/mysql-book-management/pkg/utils"
	"github.com/gorilla/mux"
)

// * Global variable to hold the new book
var NewBook models.Book

// * Get all books
func GetBooks(w http.ResponseWriter, r *http.Request) {
	// * Get all books
	newBooks := models.GetAllBooks()

	// * Convert the books to json
	res, _ := json.Marshal(newBooks)

	// * Set the header and write the response
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// * Get a book by its ID
func GetBookById(w http.ResponseWriter, r *http.Request) {
	//* Get the book ID from the request
	vars := mux.Vars(r)

	// * Convert the book ID to an integer
	bookId := vars["bookId"]
	Id, err := strconv.ParseInt(bookId, 0, 0)
	if err != nil {
		fmt.Println()
		return
	}

	// * Get the book by its ID
	bookDetails, _ := models.GetBookById(Id)

	// * Convert the book to json and write the response
	res, _ := json.Marshal(bookDetails)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(res)
}

// * Create a new book
func CreateBook(w http.ResponseWriter, r *http.Request) {
	// * Create a new book
	CreateBook := &models.Book{}

	// * Parse the request body
	utils.ParseBody(r, CreateBook)

	// * Create the book
	b := CreateBook.CreateBook()

	// * Convert the book to json and write the response
	res, _ := json.Marshal(b)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusCreated)
	w.Write(res)
}

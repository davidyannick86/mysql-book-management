package models

import (
	"github.com/davidyannick86/mysql-book-management/pkg/config"
	"github.com/jinzhu/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `gorm:""json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

// * Initialize the database connection
func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

// * Create a new book
func (b *Book) CreateBook() *Book {

	// * Create a new record
	db.NewRecord(b)
	db.Create(&b)
	return b
}

// * Get all books
func GetAllBooks() []Book {

	// * Get all books
	var Books []Book

	// * Find all books
	db.Find(&Books)

	// * Return all books
	return Books
}

// * Get a book by its ID
func GetBookById(Id int64) (*Book, *gorm.DB) {

	// * Get a book by its ID
	var getBook Book

	// * Find the book by its ID
	db := db.Where("ID = ?", Id).Find(&getBook)

	// * Return the book
	return &getBook, db
}

// * Delete a book
func DeleteBook(Id int64) Book {
	var book Book
	db.Where("ID = ?", Id).Delete(book)
	return book
}

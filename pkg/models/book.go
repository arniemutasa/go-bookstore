package models

import (
	"github.com/arniemutasa/bookstore/pkg/config"
	"gorm.io/gorm"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `gorm:""json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

// connect to database and perform migration
func init() {
	config.Connect()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

// Model function to Create New Book
func (b *Book) CreateBook() *Book {
	// db.NewRecord(b)
	db.Create(&b)
	return b
}

// Finds all books in the database
func GetAllBooks() []Book {
	var books []Book
	db.Find(&books)
	return books
}

// Finds book that matches the supplied ID
func GetBookById(id int64) (*Book, *gorm.DB) {
	var getBook Book
	db := db.Where("ID=?", id).Find(getBook)
	return &getBook, db
}

// Deletes book that matches the supplied ID
func DeleteBook(ID int64) Book {
	var book Book
	db.Where("ID=?", ID).Delete(book)
	return book
}

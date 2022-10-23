package models

import (
	"github.com/jinzhu/gorm"
	"tutorial.com.3.bookstore-management/pkg/config"
)

var db *gorm.DB

type Book struct {
	gorm.Model
	Name        string `json:"name"`
	Author      string `json:"author"`
	Publication string `json:"publication"`
}

func init() {
	config.ConnectDB()
	db = config.GetDB()
	db.AutoMigrate(&Book{})
}

func (book *Book) CreateBook() *Book {
	db.NewRecord(book)
	db.Create(&book)
	return book
}

func GetAllBooks() []Book {
	var books []Book
	db.Find(&books)
	return books
}

func GetBookByID(id int) (*Book, *gorm.DB) {
	var foundBook Book
	db := db.Where("ID=?", id).Find(&foundBook)
	return &foundBook, db
}

func DeleteBookByID(id int) Book {
	var book Book
	db.Where("ID=?", id).Delete(book)
	return book
}

func UpdateBookByID(id int, book *Book) *gorm.DB {
	return db.Model(&book).Where("ID=?", id).
		Updates(Book{Name: book.Name, Model: book.Model, Author: book.Author, Publication: book.Publication})
}

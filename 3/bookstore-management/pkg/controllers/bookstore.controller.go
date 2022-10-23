package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"tutorial.com.3.bookstore-management/pkg/models"
	"tutorial.com.3.bookstore-management/pkg/utils"
)

func GetAllBooksHandler(w http.ResponseWriter, r *http.Request) {
	books := models.GetAllBooks()
	jsonResponse, _ := json.Marshal(books)
	w.Header().Set("Content-type", "application/json")
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)

}

func GetBookByIDHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	id, err := strconv.Atoi(params["bookID"])

	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		json.NewEncoder(w).Encode(map[string]string{"bookID": "Book ID is required"})
	}

	foundBook, _ := models.GetBookByID(id)
	if foundBook == nil {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	json.NewEncoder(w).Encode(foundBook)
	w.WriteHeader(http.StatusFound)

}

func DeleteBookByIDHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	bookID, err := strconv.Atoi(vars["bookID"])
	if err != nil {
		w.WriteHeader(http.StatusUnprocessableEntity)
		jsonResponse, _ := json.Marshal(map[string]string{"bookID": "Book Id is required"})
		w.Write(jsonResponse)
		return
	}
	deletedBooks := models.DeleteBookByID(bookID)
	w.WriteHeader(http.StatusOK)
	jsonResponse, _ := json.Marshal(deletedBooks)
	w.Write(jsonResponse)
}

func UpdateBookByIDHandler(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["bookID"])

	parsedBook := &models.Book{}
	utils.ParseBody(r, parsedBook)

	db := models.UpdateBookByID(id, parsedBook)

	jsonResponse, _ := json.Marshal(db)

	w.WriteHeader(http.StatusCreated)

	w.Write(jsonResponse)
}

func CreateBookHandler(w http.ResponseWriter, r *http.Request) {
	parsedBook := &models.Book{}
	utils.ParseBody(r, parsedBook)
	createdBook := parsedBook.CreateBook()
	jsonResponse, err := json.Marshal(createdBook)
	if err != nil {
		log.Fatal(err)
	}
	w.WriteHeader(http.StatusOK)
	w.Write(jsonResponse)

}

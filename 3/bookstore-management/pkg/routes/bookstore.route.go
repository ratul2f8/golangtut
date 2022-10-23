package routes

import (
	"github.com/gorilla/mux"
	"tutorial.com.3.bookstore-management/pkg/controllers"
)

var RegisterBookStoreRoutes = func(router *mux.Router) {
	router.HandleFunc("/books", controllers.GetAllBooksHandler).Methods("GET")
	router.HandleFunc("/books/{bookID}", controllers.GetBookByIDHandler).Methods("GET")
	router.HandleFunc("/books/{bookID}", controllers.UpdateBookByIDHandler).Methods("PUT")
	router.HandleFunc("/books/{bookID}", controllers.DeleteBookByIDHandler).Methods("DELETE")
	router.HandleFunc("/books", controllers.CreateBookHandler).Methods("POST")
}

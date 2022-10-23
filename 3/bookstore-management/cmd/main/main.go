package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"tutorial.com.3.bookstore-management/pkg/routes"
)

func main() {
	router := mux.NewRouter()
	routes.RegisterBookStoreRoutes(router)
	http.Handle("/", router)

	fmt.Println("Server is listening on port: 5000")
	log.Fatal(http.ListenAndServe("localhost:5000", router))
}

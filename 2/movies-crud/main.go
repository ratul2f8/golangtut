package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
)

type Movie struct {
	ID       string    `json:"id"`
	ISBN     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

var movies []Movie

func greetingsHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode("Server is up and running")
}

func getAllMoviesHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func deleteMovieHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	parms := mux.Vars(r)
	var found *Movie = nil

	for index, element := range movies {
		if element.ID == parms["id"] {
			found = &movies[index]
			movies = append(movies[:index], movies[index+1:]...)
			break
		}
	}
	if found == nil {
		http.Error(w, "Movie not found", http.StatusNotFound)
		return
	}

	json.NewEncoder(w).Encode(found)
}

func createNewMovieHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-type", "application/json")
	var parsedMovie Movie
	_ = json.NewDecoder(r.Body).Decode(&parsedMovie)
	parsedMovie.ID = strconv.Itoa(rand.Intn(10000000))
	movies = append(movies, parsedMovie)

	json.NewEncoder(w).Encode(parsedMovie)

}

func updateMovieHandler(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	var id string = params["id"]

	var updatedData Movie
	json.NewDecoder(r.Body).Decode(&updatedData)

	for index, element := range movies {
		if element.ID == id {
			movies[index] = updatedData
			json.NewEncoder(w).Encode(movies[index])
			break
		} else if index == (len(movies) - 1) {
			http.Error(w, "Movie not found", http.StatusNotFound)
			return
		}
	}

}

func main() {

	movies = append(movies, Movie{ID: "1", ISBN: "1234", Title: "Movie 1",
		Director: &Director{FirstName: "Tom", LastName: "Cruise"}})
	movies = append(movies, Movie{ID: "2", ISBN: "4567", Title: "Movie 2",
		Director: &Director{FirstName: "Jennifer", LastName: "Lawrence"}})

	router := mux.NewRouter()
	router.HandleFunc("/", greetingsHandler)
	router.HandleFunc("/movies", getAllMoviesHandler).Methods("GET")
	router.HandleFunc("/movies/{id}", deleteMovieHandler).Methods("DELETE")
	router.HandleFunc("/movies", createNewMovieHandler).Methods("POST")
	router.HandleFunc("/movies/{id}", updateMovieHandler).Methods("PUT")

	var port int = 5000

	fmt.Println(`Server is Listening on port:`, strconv.Itoa(port))
	err := http.ListenAndServe(":"+strconv.Itoa(port), router)
	if err != nil {
		log.Fatal(err)
	}
}

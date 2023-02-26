package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

// import (
// 	"fmt"
// 	"log"
// 	"encoding/json"
// 	"math/rand"
// 	"net/http"
// 	"strconv"
// 	"github.com/gorilla/mux"
// )

type Movie struct {
	ID       string    `json:"id"`
	Isbn     string    `json:"isbn"`
	Title    string    `json:"title"`
	Director *Director `json:"director"`
}

type Director struct {
	Firstname string `json:"firstname"`
	Lastname  string `json:"lastname"`
}

var movies []Movie

func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

func main() {
	r := mux.newRouter()

	movies = append(movies, Movie{ID: "1", Isbn: "438227", Title: "Movie One", Director: &Director{Firstname: "John", Lastname: "Doe"}})
	movies = append(movies, Movie{ID: "2", Isbn: "438228", Title: "Movie Two", Director: &Director{Firstname: "Steve", Lastname: "Smith"}})

	// router
	r.handleFunc("/movies", getMovies).Methods("GET")
	r.handleFunc("/movies/{id}", getMovie).Methods("GET")
	r.handleFunc("/movies", createMovie).Methods("POST")
	r.handleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.handleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	fmt.Printf("Starting server on port 8000")
	log.Fatal(http.ListenAndServe(":8000", r))
}

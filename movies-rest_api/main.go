package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

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

// get all movies
func getMovies(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(movies)
}

//delete a movie

func deleteMovie(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "applicatiob/json")
	params := mux.Vars(r)

	for index, Item := range movies {
		if Item.ID == params["id"] {
			movies = append(movies[:index], movies[index+1:]...)
			break
		}

	}

}

func main() {
	r := mux.NewRouter()

	movies = append(movies, Movie{ID: "1", Isbn: "234567", Title: "The fleek", Director: &Director{Firstname: "John", Lastname: "Doe"}})
	movies = append(movies, Movie{ID: "2", Isbn: "978654", Title: "Merlin", Director: &Director{Firstname: "nelly", Lastname: "Faith"}})
	r.HandleFunc("/movies", getMovies).Methods("GET")
	//r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	//r.HandleFunc("/movies", createMovie).Methods("POST")
	//r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	//start server
	fmt.Printf("Starting server at port 8000")
	log.Fatal(http.ListenAndServe(":8000", r))

}
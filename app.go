package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

//GET
//ENDPOINT: http:localhost:8080/movies
func allMoviesEndPoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Muestras todas las peliculas")
}

//GET
//ENDPOINT: http:localhost:8080/movies/{id}
//ENDPOINT: http:localhost:8080/movies/4
func findMovieEndPoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Muestra pelicula específica segun id")
}

// POST
// ENDPOINT: http:localhost:8080/movies
func createMovieEndPoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Registra una película")
}

// PUT
// ENDPOINT: http:localhost:8080/movies
func updateMovieEndPoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Actualiza una película")
}

// DELETE
// ENDPOINT: http:localhost:8080/movies/{id}
func deleteMovieEndPoint(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Elimina una película segun id")
}
func main() {
	r := mux.NewRouter()

	r.HandleFunc("/movies", allMoviesEndPoint).Methods("GET")
	r.HandleFunc("/movies/{ID}", findMovieEndPoint).Methods("GET")
	r.HandleFunc("/movies", createMovieEndPoint).Methods("POST")
	r.HandleFunc("/movies", updateMovieEndPoint).Methods("UPDATE")
	r.HandleFunc("/movies/{ID}", deleteMovieEndPoint).Methods("DELETE")
	if err := http.ListenAndServe(":3000", r); err != nil {
		log.Fatal(err)
	}
}

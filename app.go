package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	. "github.com/wilsonv244/API_3/config"
	. "github.com/wilsonv244/API_3/dao"
	. "github.com/wilsonv244/API_3/models"
	"gopkg.in/mgo.v2/bson"
)

var config = Config{}
var dao = MoviesDAO{}

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
	defer r.Body.Close()
	var movie Movie
	if err := json.NewDecoder(r.Body).Decode(&movie); err != nil {
		respondWithError(w, http.StatusBadRequest, "Invalid request payload")
		return
	}
	movie.ID = bson.NewObjectId()
	if err := dao.Insert(movie); err != nil {
		respondWithError(w, http.StatusInternalServerError, err.Error())
		return
	}
	respondWithJson(w, http.StatusCreated, movie)
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
func respondWithError(w http.ResponseWriter, code int, msg string) {
	respondWithJson(w, code, map[string]string{"error": msg})
}

func respondWithJson(w http.ResponseWriter, code int, payload interface{}) {
	response, _ := json.Marshal(payload)
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(code)
	w.Write(response)
}

// Parse the configuration file 'config.toml', and establish a connection to DB
func init() {
	config.Read()
	dao.Server = config.Server
	dao.Database = config.Database
	dao.Connect()
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

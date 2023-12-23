package controllers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func New() http.Handler {
	r := mux.NewRouter()

	r.HandleFunc("/movies", getAllMovies).Methods("GET")
	r.HandleFunc("/movies/{id}", getMovie).Methods("GET")
	r.HandleFunc("/movies", createMovie).Methods("POST")
	r.HandleFunc("/movies/{id}", updateMovie).Methods("PUT")
	r.HandleFunc("/movies/{id}", deleteMovie).Methods("DELETE")

	r.HandleFunc("/director", getAllDirectors).Methods("GET")
	r.HandleFunc("/director/{id}", getDirector).Methods("GET")
	r.HandleFunc("/director", createDirector).Methods("POST")
	r.HandleFunc("/director/{id}", updateDirector).Methods("PUT")
	r.HandleFunc("/director/{id}", deleteDirector).Methods("DELETE")

	fmt.Printf("Starting server at port 8000\n")
	log.Fatal(http.ListenAndServe(":8000", r))
	return r
}

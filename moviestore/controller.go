package moviestore

import (
	"encoding/json"
	"io"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"../dao/factory"
	"../models"
)

//Controller ...
type Controller struct {}

var db = factory.FactoryDao("mongodb")
// Index GET /
func (c *Controller) Index(w http.ResponseWriter, r *http.Request) {

	movies, err := db.GetAll() // list of all movies
	if err != nil {
		log.Println("Error Index",err)
	}
	log.Println(movies)
	data, _ := json.Marshal(movies)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	w.Write(data)
	return
}

// AddMovie POST /
func (c *Controller) AddMovie(w http.ResponseWriter, r *http.Request) {

	var movie models.Movie
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576)) // read the body of the request
	log.Println(body)
	if err != nil {
		log.Println("Error AddMovie", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if err := r.Body.Close(); err != nil {
		log.Println("Error AddMovie", err)
	}
	if err := json.Unmarshal(body, &movie); err != nil { // unmarshall body contents as a type Candidate
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			log.Println("Error AddMovie unmarshalling data", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}
	 // adds the movie to the DB
	if err := db.AddMovie(movie); err != nil {
		log.Println(err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	return
}

// UpdateMovie UPDATE /
func (c *Controller) UpdateMovie(w http.ResponseWriter, r *http.Request) {
	var movie models.Movie
	body, err := ioutil.ReadAll(io.LimitReader(r.Body, 1048576)) // read the body of the request
	if err != nil {
		log.Println("Error UpdateAlbum", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if err := r.Body.Close(); err != nil {
		log.Println("Error UpdateMovie", err)
	}
	if err := json.Unmarshal(body, &movie); err != nil { // unmarshall body contents as a type Candidate
		w.Header().Set("Content-Type", "application/json; charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if err := json.NewEncoder(w).Encode(err); err != nil {
			log.Println("Error UpdateMovie unmarshalling data", err)
			w.WriteHeader(http.StatusInternalServerError)
			return
		}
	}

	 // adds the movie to the DB
	if err := db.UpdateMovie(movie); err != nil {
		log.Println("Error UpdateMovie updating db", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	return
}

// DeleteMovie DELETE /
func (c *Controller) DeleteMovie(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"] // param id

	if err := db.DeleteMovie(id); err != nil { // delete a album by id
		log.Println("Error DeleteAlbum",err)
		w.WriteHeader(http.StatusNotFound)
		return
	}

	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)

	return
}

// GetMovieByTitle GET /
func (c *Controller) GetMovieByTitle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	title := vars["title"]

	movie,err := db.GetByTitle(title)
	if err != nil {
		log.Println("Error GetMovieByTitle", err)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	log.Println(movie)
	data, _ := json.Marshal(movie)
	w.Header().Set("Content-Type", "application/json; charset=UTF-8")
	w.Header().Set("Access-Control-Allow-Origin", "*")
	w.WriteHeader(http.StatusOK)
	w.Write(data)

	return
}

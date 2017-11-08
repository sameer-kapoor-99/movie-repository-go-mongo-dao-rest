package moviestore

import (
	"encoding/json"

	"log"
	"net/http"

	"github.com/gorilla/mux"
	"../dao/interfaces"
	"../models"
)

//Controller ...
type Controller struct {
	dao interfaces.UserDao
}


// Index GET /
func (c *Controller) Index(w http.ResponseWriter, r *http.Request) {

	movies, err := c.dao.GetAll() // list of all movies
	if err != nil {
		log.Println("Error Index",err)
	}
	log.Println(movies)

	json.NewEncoder(w).Encode(movies)
	return
}

// AddMovie POST /
func (c *Controller) AddMovie(w http.ResponseWriter, r *http.Request) {


	log.Println("in AddMovie")
	var movie models.Movie
	err := json.NewDecoder(r.Body).Decode(&movie)

	log.Println(movie.ID)
	if err != nil {
		log.Println("Error AddMovie", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if err := r.Body.Close(); err != nil {
		log.Println("Error AddMovie", err)
	}

	 // adds the movie to the DB
	if err := c.dao.AddMovie(movie); err != nil {
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
	err := json.NewDecoder(r.Body).Decode(&movie)
	if err != nil {
		log.Println("Error UpdateAlbum", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if err := r.Body.Close(); err != nil {
		log.Println("Error UpdateMovie", err)
	}

	 // adds the movie to the DB
	if err := c.dao.UpdateMovie(movie); err != nil {
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

	if err := c.dao.DeleteMovie(id); err != nil { // delete a album by id
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
	movie,err := c.dao.GetByTitle(title)
	if err != nil {
		log.Println("Error GetMovieByTitle", err)
		w.WriteHeader(http.StatusNotFound)
		return
	}
	log.Println(movie)
	json.NewEncoder(w).Encode(movie)

	return
}

package main

import (
	"log"
	"./moviestore"
	"net/http"
)



func main() {

	router := moviestore.NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))


}

package main

import (
	"log"
	"./moviestore"
	"net/http"
	"./dao/factory"
)



func main() {


	router := moviestore.NewRouter(factory.FactoryDao("mongodb"))

	log.Fatal(http.ListenAndServe(":8080", router))


}

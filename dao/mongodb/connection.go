package mongodb

import (
	"gopkg.in/mgo.v2"
	"log"
)

const SERVER = "localhost:27017"


func get() *mgo.Session {
	session, err := mgo.Dial(SERVER)
	if err != nil {
		log.Fatalln(err)
		return nil
	}

	return session
}

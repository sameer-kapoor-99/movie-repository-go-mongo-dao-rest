package mongodb

import (
	"../../models"
	"gopkg.in/mgo.v2/bson"
)

type UserImplMongodb struct {
}

const DBNAME = "movierepo"

const DOCNAME = "movies"

func (dao UserImplMongodb) AddMovie(m models.Movie) error {
	db := get()
	defer db.Close()

	return db.DB(DBNAME).C(DOCNAME).Insert(m)

}

func (dao UserImplMongodb) GetAll() ([]models.Movie, error) {

	db := get()
	defer db.Close()

	res := []models.Movie{}

	if err := db.DB(DBNAME).C(DOCNAME).Find(nil).All(&res); err != nil {
		return nil, err
	}

	return res, nil

}

func (dao UserImplMongodb) GetByTitle(title string) (*models.Movie, error) {
	db := get()
	defer db.Close()
	res := models.Movie{}

	if err := db.DB(DBNAME).C(DOCNAME).Find(bson.M{"_title":title}).One(&res); err != nil {
		return nil, err
	}

	return &res, nil
}

func (dao UserImplMongodb) GetById(id string) (*models.Movie, error) {
	db := get()
	defer db.Close()
	res := models.Movie{}

	if err := db.DB(DBNAME).C(DOCNAME).Find(bson.M{"_id":id}).One(&res); err != nil {
		return nil, err
	}

	return &res, nil
}

func (dao UserImplMongodb) UpdateMovie(m models.Movie) error {
	db := get()
	defer db.Close()

	if err := db.DB(DBNAME).C(DOCNAME).UpdateId(bson.M{"_id":m.ID}, m); err != nil {
		return err
	}

	return nil
}

func (dao UserImplMongodb) DeleteMovie(id string) error {
	db := get()
	defer db.Close()

	if err := db.DB(DBNAME).C(DOCNAME).RemoveId(bson.M{"_id":id}); err != nil {
		return err
	}

	return nil
}

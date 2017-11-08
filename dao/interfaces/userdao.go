package interfaces

import models "../../models"

type UserDao interface {
	AddMovie(m models.Movie) error
	/*Update(u *models.User) error
	Delete(i int) error*/
	GetByTitle(t string) (*models.Movie, error)
	GetById(i string) (*models.Movie,error)
	UpdateMovie(m models.Movie) error
	DeleteMovie(i string) error
	GetAll() ([]models.Movie, error)
}

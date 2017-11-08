package interfaces

import models "../../models"

// Interface for all DAO methods
type UserDao interface {
	AddMovie(m models.Movie) error
	GetByTitle(t string) (*models.Movie, error)
	GetById(i string) (*models.Movie,error)
	UpdateMovie(m models.Movie) error
	DeleteMovie(i string) error
	GetAll() ([]models.Movie, error)
}

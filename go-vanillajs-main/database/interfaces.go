package database

import "GoProjects/ReelingIt/models"

type MovieStorage interface {
	GetTopMovies() ([]models.Movie, error)
	GetRandomMovies() ([]models.Movie, error)
	// GetMovieByID(id int) (models.Movie, error)
	// SearchMoviesByName(name string) ([]models.Movie, error)
	// GetAllGenres() ([]models.Genre, error)
}
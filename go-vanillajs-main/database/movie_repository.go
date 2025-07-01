package database

import (
	"GoProjects/ReelingIt/logger"
	"GoProjects/ReelingIt/models"
	"database/sql"

	
)
type MovieRepository struct {
	db     *sql.DB
	logger *logger.Logger
}


// Factory function to create a new MovieRepository
func NewMovieRepository(db *sql.DB, logger *logger.Logger) (*MovieRepository, error) {
	return &MovieRepository{
		db:     db,
		logger: logger,
	}, nil
}


const defaultLimit = 20

func (r *MovieRepository) GetTopMovies() ([]models.Movie, error) {
	// Fetch movies
	query := `
		SELECT id, tmdb_id, title, tagline, release_year, overview, score, 
		       popularity, language, poster_url, trailer_url
		FROM movies
		ORDER BY popularity DESC
		LIMIT $1
	`
	return r.getMovies(query)
}

func (r *MovieRepository) GetRandomMovies() ([]models.Movie, error) {
	// Fetch random movies
	query := `
		SELECT id, tmdb_id, title, tagline, release_year, overview, score, 
		       popularity, language, poster_url, trailer_url
		FROM movies
		ORDER BY RANDOM()
		LIMIT $1
	`
	return r.getMovies(query)
}


func (r *MovieRepository) getMovies(query string) ([]models.Movie, error){
	rows, err:= r.db.Query(query, defaultLimit)
	if err != nil {
		r.logger.Error("Failed to execute query", err)
		return nil, err
	}

	defer rows.Close()

	var movies []models.Movie
	for rows.Next(){
		var m models.Movie
		err := rows.Scan(&m.ID, &m.TMBD_ID,&m.Title, &m.Tagline, &m.ReleaseYear,
			&m.Overview, &m.Score, &m.Popularity, &m.Language,
			&m.PosterURL, &m.TrailerURL)

		if err != nil {
			r.logger.Error("Failed to scan row", err)
			return nil, err
		}

		movies = append(movies, m)
	}

	return movies, nil
}
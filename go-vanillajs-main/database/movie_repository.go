package database

import (
	"GoProjects/ReelingIt/logger"
	"GoProjects/ReelingIt/models"
	"database/sql"
	"errors"
	"strconv"
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



func (r *MovieRepository) SearchMoviesByName(name string, order string, genre *int) ([]models.Movie, error){
	if name == "" {
		return nil, nil // Return empty slice if no name is provided
	}

	// Prepare the base query
	query := `
		SELECT id, tmdb_id, title, tagline, release_year, overview, score,
		       popularity, language, poster_url, trailer_url
		FROM movies
		WHERE LOWER(title) LIKE LOWER($1)
	`

	// Add genre filter if provided
	if genre != nil {
		query += ` AND id IN (SELECT movie_id FROM movie_genres WHERE genre_id = $2)`
	}

	// Add ordering
	switch order {
	case "popularity":
		query += ` ORDER BY popularity DESC`
		case "score":
		query += ` ORDER BY score DESC`
		case "date":
		query += ` ORDER BY release_year DESC`
		case "name":
		query += ` ORDER BY title ASC`
		default:
		query += ` ORDER BY popularity DESC` // Default ordering
	}

	query += ` LIMIT $3`

	args := []interface{}{"%" + name + "%", defaultLimit}
	if genre != nil {
		args = append(args, *genre)
	}

	rows, err := r.db.Query(query, args...)
	if err != nil {
		r.logger.Error("Failed to execute search query", err)
		return nil, err
	}
	defer rows.Close()

	var movies []models.Movie
	for rows.Next() {
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


func (r *MovieRepository) GetAllGenres() ([]models.Genre, error) {
	query := `SELECT id, name FROM genres ORDER BY id`
	rows, err := r.db.Query(query)
	if err != nil {
		r.logger.Error("Failed to query all genres", err)
		return nil, err
	}
	defer rows.Close()

	var genres []models.Genre
	for rows.Next() {
		var g models.Genre
		if err := rows.Scan(&g.ID, &g.Name); err != nil {
			r.logger.Error("Failed to scan genre row", err)
			return nil, err
		}
		genres = append(genres, g)
	}
	return genres, nil
}


func (r *MovieRepository) GetMovieByID(id int) (models.Movie, error) {
	query := `SELECT id, tmdb_id, title, tagline, release_year, overview, score,
		       popularity, language, poster_url, trailer_url
			   FROM movies WHERE id = $1`
	row:= r.db.QueryRow(query, id)
	
	var m models.Movie

	err := row.Scan(&m.ID, &m.TMBD_ID, &m.Title, &m.Tagline, &m.ReleaseYear,
		&m.Overview, &m.Score, &m.Popularity, &m.Language, &m.PosterURL, &m.TrailerURL)

	if err != nil {
		if err == sql.ErrNoRows {
			r.logger.Error("No movie found with the given ID", err)
			return models.Movie{}, nil // Return empty movie if not found
		}
		r.logger.Error("Failed to scan row", err)
		return models.Movie{}, err
	}

	// Fetch related data
	if err := r.fetchMovieRelations(&m); err != nil {
		return models.Movie{}, err
	}

	return m, nil

}





// fetchMovieRelations fetches genres, actors, and keywords for a movie
func (r *MovieRepository) fetchMovieRelations(m *models.Movie) error {
	// Fetch genres
	genreQuery := `
		SELECT g.id, g.name 
		FROM genres g
		JOIN movie_genres mg ON g.id = mg.genre_id
		WHERE mg.movie_id = $1
	`
	genreRows, err := r.db.Query(genreQuery, m.ID)
	if err != nil {
		r.logger.Error("Failed to query genres for movie "+strconv.Itoa(m.ID), err)
		return err
	}
	defer genreRows.Close()
	for genreRows.Next() {
		var g models.Genre
		if err := genreRows.Scan(&g.ID, &g.Name); err != nil {
			r.logger.Error("Failed to scan genre row", err)
			return err
		}
		m.Genres = append(m.Genres, g)
	}

	// Fetch actors
	actorQuery := `
		SELECT a.id, a.first_name, a.last_name, a.image_url
		FROM actors a
		JOIN movie_cast mc ON a.id = mc.actor_id
		WHERE mc.movie_id = $1
	`
	actorRows, err := r.db.Query(actorQuery, m.ID)
	if err != nil {
		r.logger.Error("Failed to query actors for movie "+strconv.Itoa(m.ID), err)
		return err
	}
	defer actorRows.Close()
	for actorRows.Next() {
		var a models.Actor
		if err := actorRows.Scan(&a.ID, &a.FirstName, &a.LastName, &a.ImageURL); err != nil {
			r.logger.Error("Failed to scan actor row", err)
			return err
		}
		m.Casting = append(m.Casting, a)
	}

	// Fetch keywords
	keywordQuery := `
		SELECT k.word
		FROM keywords k
		JOIN movie_keywords mk ON k.id = mk.keyword_id
		WHERE mk.movie_id = $1
	`
	keywordRows, err := r.db.Query(keywordQuery, m.ID)
	if err != nil {
		r.logger.Error("Failed to query keywords for movie "+strconv.Itoa(m.ID), err)
		return err
	}
	defer keywordRows.Close()
	for keywordRows.Next() {
		var k string
		if err := keywordRows.Scan(&k); err != nil {
			r.logger.Error("Failed to scan keyword row", err)
			return err
		}
		m.Keywords = append(m.Keywords, k)
	}

	return nil
}

var (
	ErrMovieNotFound = errors.New("movie not found")
)
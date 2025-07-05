package handlers

import (
	"GoProjects/ReelingIt/database"
	"GoProjects/ReelingIt/logger"
	"GoProjects/ReelingIt/models"
	"strconv"

	"encoding/json"
	"net/http"
)
type MovieHandler struct {
	Storage database.MovieStorage // Using interface lets us change the database later
	Logger  *logger.Logger
}

func (h * MovieHandler) writeJSONResponse(w http.ResponseWriter, data interface{}) error{
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(data); err != nil {
	    h.Logger.Error("Failed to encode JSON response", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return  err
	}
	return nil
}


func (h *MovieHandler) handleStorageError(w http.ResponseWriter, err error, context string) bool {
	if err != nil {
		if err == database.ErrMovieNotFound {
			http.Error(w, context, http.StatusNotFound)
			return true
		}
		h.Logger.Error(context, err)
		http.Error(w, "Internal server error", http.StatusInternalServerError)
		return true
	}
	return false
}

func (h *MovieHandler) parseID(w http.ResponseWriter, idStr string) (int, bool) {
	id, err := strconv.Atoi(idStr)
	if err != nil {
		h.Logger.Error("Invalid ID format", err)
		http.Error(w, "Invalid ID", http.StatusBadRequest)
		return 0, false
	}
	return id, true
}


func (h *MovieHandler) GetTopMovies(w http.ResponseWriter, r *http.Request){
	movies, err := h.Storage.GetTopMovies()
	if err != nil {
		h.Logger.Error("Failed to get top movies", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}
	h.writeJSONResponse(w, movies)

}

func (h *MovieHandler) GetRandomMovies(w http.ResponseWriter, r *http.Request) {
	Movies, err := h.Storage.GetRandomMovies()
	if err != nil {
		h.Logger.Error("Failed to get random movies", err)
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}


	h.writeJSONResponse(w, Movies)
}

func (h *MovieHandler) SearchMovies(w http.ResponseWriter, r *http.Request) {
	query := r.URL.Query().Get("q")
	order := r.URL.Query().Get("order")
	genreStr := r.URL.Query().Get("genre")

	var genre *int
	if genreStr != "" {
		genreInt, ok := h.parseID(w, genreStr)
		if !ok {
			return
		}
		genre = &genreInt
	}

	var movies []models.Movie
	var err error
	if query != "" {
		movies, err = h.Storage.SearchMoviesByName(query, order, genre)
	}
	if h.handleStorageError(w, err, "Failed to get movies") {
		return
	}
	if h.writeJSONResponse(w, movies) == nil {
		h.Logger.Info("Successfully served movies")
	}
}

func (h *MovieHandler) GetMovieByID(w http.ResponseWriter, r *http.Request) {
	id := r.URL.Query().Get("id")
	if id == "" {
		http.Error(w, "ID is required", http.StatusBadRequest)
		return
	}
	movieID, valid := h.parseID(w, id)
	if !valid {
		return
	}

	movie, err := h.Storage.GetMovieByID(movieID)
	if h.handleStorageError(w, err, "Failed to get movie by ID") {
		return
	}

	h.writeJSONResponse(w, movie)

}


func (h *MovieHandler) GetGenres(w http.ResponseWriter, r *http.Request) {
	genres, err := h.Storage.GetAllGenres()
	if h.handleStorageError(w, err, "Failed to get genres") {
		return
	}
	if h.writeJSONResponse(w, genres) == nil {
		h.Logger.Info("Successfully served genres")
	}
}

func NewMovieHandler(storage database.MovieStorage, log *logger.Logger) *MovieHandler {
	return &MovieHandler{
		Storage: storage,
		Logger:  log,
	}
}
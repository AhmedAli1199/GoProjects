package handlers

import (
	"GoProjects/ReelingIt/models"
	"encoding/json"
	"net/http"
)
type MovieHandler struct {
}

func (h * MovieHandler) writeJSONResponse(w http.ResponseWriter, data interface{}){
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(data); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
	}
}

func (h *MovieHandler) GetTopMovies(w http.ResponseWriter, r *http.Request){
	movies := []models.Movie{
		{
			ID:          1,
			TMBD_ID:    123456,
			Title:       "Inception",
			Tagline:     "Your mind is the scene of the crime.",
			ReleaseYear: 2010,
			Genres: []models.Genre{
				{ID: 1, Name: "Science Fiction"},
				{ID: 2, Name: "Action"},
			},
			Overview:    nil,
			Score:       nil,
			Popularitey: nil,
			Keywords: []string{"dream", "thief", "subconscious"},
		},
	}
	h.writeJSONResponse(w, movies)

}

func (h *MovieHandler) GetRandomMovies(w http.ResponseWriter, r *http.Request) {
	topMovies := []models.Movie{
		{
			ID:          1,
			TMBD_ID:    123456,
			Title:       "Inception Random",
			Tagline:     "Your mind is the scene of the crime.",
			ReleaseYear: 2010,
			Genres: []models.Genre{
				{ID: 1, Name: "Science Fiction"},
				{ID: 2, Name: "Action"},
			},
			Overview:    nil,
			Score:       nil,
			Popularitey: nil,
			Keywords: []string{"dream", "thief", "subconscious"},
		},
	}
	h.writeJSONResponse(w, topMovies)
}

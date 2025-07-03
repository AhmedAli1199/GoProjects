package api

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/go-chi/chi/v5"
)

type WorkoutHandler struct {
}

func NewWorkoutHandler() *WorkoutHandler {

	return &WorkoutHandler{}

}

func (wh *WorkoutHandler) HandleWorkoutByID(w http.ResponseWriter, r *http.Request){

	paramWorkoutID := chi.URLParam(r, "id")
	if paramWorkoutID == "" {
		http.NotFound(w,r)
		return
	}

	workoutID, err := strconv.ParseInt(paramWorkoutID, 10, 64)
	if err != nil {
		http.NotFound(w,r)
		return
	}

	fmt.Fprintf(w, "Workout id is %d", workoutID)

}

func (wh *WorkoutHandler) CreateWorkout(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w ,"New workout Created")

}
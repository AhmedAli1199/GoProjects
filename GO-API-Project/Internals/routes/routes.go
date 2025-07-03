package routes

import (
	"github.com/AhmedAli1199/WorkoutManager/Internals/app"
	"github.com/go-chi/chi/v5"
)

func SetupRoutes(app *app.Application) *chi.Mux{
	r := chi.NewRouter()

	// Health check route
	r.Get("/health", app.HealthChecker)

	r.Get("/workouts/{id}", app.WorkoutHandler.HandleWorkoutByID)
	r.Post("/workouts", app.WorkoutHandler.CreateWorkout)

	
	return r


}
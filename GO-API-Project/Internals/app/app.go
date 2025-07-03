package app

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/AhmedAli1199/WorkoutManager/Internals/api"
)

type Application struct {
	Logger *log.Logger
	WorkoutHandler *api.WorkoutHandler
}

func NewApplication() (*Application, error) {

	logger := log.New(os.Stdout, "[APP] ", log.LstdFlags|log.Lshortfile|log.Ldate|log.Ltime)


	// Create Handlers here:
	workout_handler := api.NewWorkoutHandler()

	app := &Application{
		Logger: logger,
		WorkoutHandler: workout_handler,
	}

	return app, nil
}

func (a *Application) HealthChecker(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response := map[string]string{"status": "healthy"}
	json.NewEncoder(w).Encode(response)
	fmt.Println("Health check endpoint hit")
}
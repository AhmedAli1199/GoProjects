package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/AhmedAli1199/WorkoutManager/Internals/app"
)

func main(){

	app, err := app.NewApplication()

	if err != nil{

		panic(err)
	}

	app.Logger.Println("Application started successfully")

	http.HandleFunc("/health", HealthChecker)

	server := http.Server{
		Addr:   ":8000",
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	server.ListenAndServe()
	app.Logger.Println("Server is running on port 8000")

}

func HealthChecker(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	response := map[string]string{"status": "healthy"}
	json.NewEncoder(w).Encode(response)
	fmt.Println("Health check endpoint hit")
}

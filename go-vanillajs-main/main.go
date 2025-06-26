package main

import (
	"GoProjects/ReelingIt/handlers"
	"GoProjects/ReelingIt/logger"
	"log"
	"net/http"
)

func initializeLogger() *logger.Logger {
	logFilePath := "movies.log"
	logger, err := logger.NewLogger(logFilePath)

	if err != nil {
		log.Fatalf("Failed to initialize logger: %v", err)
	}

	log.Println("Logger initialized successfully")

	return logger

}

func main() {

	loggerInstance := initializeLogger()

	movieHandler := handlers.MovieHandler{}

	http.HandleFunc("/api/movies/top", movieHandler.GetTopMovies)

	http.HandleFunc("/api/movies/random", movieHandler.GetRandomMovies)

	http.Handle("/", http.FileServer(http.Dir("public")))

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
		loggerInstance.Error("Failed to start server", err)
	}

	log.Println("Server started on :8080")

}

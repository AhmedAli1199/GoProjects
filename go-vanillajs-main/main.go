package main

import (
	"GoProjects/ReelingIt/database"
	"GoProjects/ReelingIt/handlers"
	"GoProjects/ReelingIt/logger"
	"database/sql"
	"log"
	"net/http"
	"os"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq" // PostgreSQL driver
	// Importing the PostgreSQL driver for database connection
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

	// Logger initialization
	loggerInstance := initializeLogger()

	//Environment variables loading
	if err:=godotenv.Load();err!=nil{
		log.Fatalf("Error loading .env file: %v", err)
	}

	dbString := os.Getenv("DATABASE_URL")
	if dbString == "" {
		log.Fatal("DATABASE_URL environment variable is not set")
	} 

	db, err := sql.Open("postgres", dbString)
	if err != nil {
		log.Fatalf("Failed to connect to the database: %v", err)
	}

	defer db.Close()
	


	//Initialize MovieRepository
	movieRepo, err := database.NewMovieRepository(db, loggerInstance)
	if err != nil {
		log.Fatalf("Failed to create movie repository: %v", err)
		loggerInstance.Error("Failed to create movie repository", err)
	}
      

	//Initialize a movie handler which will use the movie repository to handle requests
	// This allows us to change the database implementation later without changing the handler code
	movieHandler := handlers.MovieHandler{
		Storage: movieRepo,
		Logger:  loggerInstance,
	}

	http.HandleFunc("/api/movies/top", movieHandler.GetTopMovies)

	http.HandleFunc("/api/movies/random", movieHandler.GetRandomMovies)

	http.HandleFunc("/api/movies", movieHandler.GetMovieByID)

	http.HandleFunc("/api/genres", movieHandler.GetGenres)

	http.HandleFunc("/api/movies/search", movieHandler.SearchMovies)


	http.Handle("/", http.FileServer(http.Dir("public")))

	err_server := http.ListenAndServe(":8080", nil)
	if err_server != nil {
		log.Fatalf("Failed to start server: %v", err_server)
		loggerInstance.Error("Failed to start server", err_server)
	}

	log.Println("Server started on :8080")

}

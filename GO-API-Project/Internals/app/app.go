package app

import (
	"log"
	"os"
)

type Application struct {
	Logger *log.Logger
}

func NewApplication() (*Application, error) {

	logger := log.New(os.Stdout, "[APP] ", log.LstdFlags|log.Lshortfile|log.Ldate|log.Ltime)

	app := &Application{
		Logger: logger,
	}

	return app, nil
}
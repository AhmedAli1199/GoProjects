package main

import (

	"flag"
	"fmt"
	"net/http"
	"time"

	"github.com/AhmedAli1199/WorkoutManager/Internals/app"
	"github.com/AhmedAli1199/WorkoutManager/Internals/routes"
)

func main(){

	var port int
	flag.IntVar(&port, "port", 8000, "Port to run the server on")
	flag.Parse()


	app, err := app.NewApplication()

	if err != nil{

		panic(err)
	}

	app.Logger.Println("Application started successfully")

	
	r := routes.SetupRoutes(app)
	server := http.Server{
		Addr:      fmt.Sprintf(":%d", port),
		Handler:   r,
		ReadTimeout: 10 * time.Second,
		WriteTimeout: 30 * time.Second,
	}

	server.ListenAndServe()
	app.Logger.Println("Server is running on port", port)

}



package main

import (
	"log"
	"net/http"
)


func  main()  {

	http.Handle("/", http.FileServer(http.Dir("public")))

	err :=http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Failed to start server: %v", err)
	}
	
	log.Println("Server started on :8080")
	
	
}
package main

import (
	"image-service/internal/server"
	"log"
	"net/http"

	"github.com/joho/godotenv"
)

// portNumber is the port number that the server will listen on
const portNumber = ":8000"

// main is the entry point for the application
func main() {
	log.Println("Starting server on port", portNumber)

	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Declare the app
	s := server.New()

	// Declare the server
	srv := &http.Server{
		Addr:    portNumber,
		Handler: s.Routes(),
	}

	// Start the server
	err = srv.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}

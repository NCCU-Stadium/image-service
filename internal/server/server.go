package server

import (
	"image-service/internal/config"
	"image-service/internal/database"
	"log"
	"net/http"
)

type Server struct {
	database database.Database
	config   config.AppConfig
}

func New() Server {
	config := config.New()
	db := database.New(config)

	return Server{
		config:   config,
		database: db,
	}
}

// routes defines the routes for the application
func (s *Server) Routes() http.Handler {
	// Declare a new router
	mux := http.NewServeMux()

	mux.HandleFunc("/", s.home)
	mux.HandleFunc("/upload", s.uploadImage)
	mux.HandleFunc("/display", s.displayImage)
	mux.HandleFunc("/delete", s.deleteImage)

	return mux
}

type Response struct {
	Message string `json:"message"`
	Id      string `json:"id"`
}

// Define the middlewares for the application
func (s *Server) middlewares(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Print("Executing middleware")

		//Allow CORS here By * or specific origin
		w.Header().Set("Access-Control-Allow-Origin", "*")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")

		next.ServeHTTP(w, r)
		// log.Print("Executing middlewareTwo again")
	})
}

// home is the handler for the home page
func (s *Server) home(w http.ResponseWriter, r *http.Request) {
	log.Println("Home page accessed")
	w.Write([]byte("Hello from HandleFunc #1"))
}

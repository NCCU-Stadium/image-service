package server

import (
	"encoding/json"
	"fmt"
	"image-service/internal/models"
	"io"
	"log"
	"net/http"
	"time"
)

func (s *Server) uploadImage(w http.ResponseWriter, r *http.Request) {
	// Check if the request is a POST request
	if r.Method != http.MethodPost {
		s.httpError(&w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Get the collection name from the query string
	collectionName := r.URL.Query().Get("collection")

	// Parse our multipart form, 10 << 20 specifies a maximum upload of 10 MB files.
	r.ParseMultipartForm(10 << 20)

	// Get the file from the formdata
	file, handler, err := r.FormFile("image")
	if err != nil {
		fmt.Println("Error Retrieving the File")
		fmt.Println(err)
	}
	defer file.Close()

	// Read the image data & type
	imageData, err := io.ReadAll(file)
	if err != nil {
		fmt.Println("Error Reading the File")
		fmt.Println(err)
	}
	imageType := http.DetectContentType(imageData)
	log.Println(imageType)

	// Construct the new image
	image := models.MongoFields{
		Name:     handler.Filename,
		Data:     imageData,
		Type:     imageType,
		Uploaded: time.Now(),
	}

	// Insert the image into the database
	primitiveId, err := s.database.Insert(image, collectionName)
	if err != nil {
		s.httpError(&w, "Error inserting image data into MongoDB", http.StatusInternalServerError)
		return
	}

	// Respond with success message
	res := Response{
		Message: "Image uploaded successfully",
		Id:      primitiveId,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)

	return
}

func (s *Server) displayImage(w http.ResponseWriter, r *http.Request) {
	// Check if the request is a GET request
	if r.Method != http.MethodGet {
		s.httpError(&w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse the query string parameter _id and collection
	queryValues := r.URL.Query()
	targetId := queryValues.Get("_id")
	collectionName := queryValues.Get("collection")

	// Retrieve the image from MongoDB by _id
	result, err := s.database.FindById(targetId, collectionName)

	// Set appropriate headers
	w.Header().Set("Content-Type", result.Type)
	w.Header().Set("Content-Disposition", "inline; filename="+result.Name)

	// Write the image binary data directly to the response
	imageData := result.Data
	_, err = w.Write(imageData)
	if err != nil {
		s.httpError(&w, "Error writing image data to response", http.StatusInternalServerError)
	}
}

// deleteImage is the handler for the delete route
func (s *Server) deleteImage(w http.ResponseWriter, r *http.Request) {
	// Check if the request is a DELETE request
	if r.Method != http.MethodDelete {
		s.httpError(&w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	// Parse the query string parameter _id and collection
	queryValues := r.URL.Query()
	targetId := queryValues.Get("_id")
	collectionName := queryValues.Get("collection")

	// Delete the image from MongoDB by _id
	deletedCount, err := s.database.DeleteById(targetId, collectionName)
	if err != nil {
		s.httpError(&w, "Error deleting image from MongoDB", http.StatusInternalServerError)
	}

	// Respond with success Message
	res := Response{
		Message: fmt.Sprintf("Deleted %d documents", deletedCount),
		Id:      targetId,
	}
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(res)
}

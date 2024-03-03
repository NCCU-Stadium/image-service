package database

import (
	"context"
	"image-service/internal/config"
	"log"
	"time"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Database struct {
	db *mongo.Database
}

func New(config config.AppConfig) Database {
	client := Connect(config.MongoURI)

	return Database{
		db: client.Database(config.MongoDatabase),
	}
}

// connectToDB connects to the MongoDB database and returns the collection
func Connect(uri string) *mongo.Client {
	log.Println("Connecting to MongoDB...")

	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// Use this for when running in Docker
	// clientOptions := options.Client().ApplyURI("mongodb://root:rootpassword@mongo:27017/")
	// client, err := mongo.Connect(ctx, clientOptions)

	// Use this for when running locally
	// client, err := mongo.Connect(ctx, options.Client().ApplyURI("mongodb://127.0.0.1:27017/"))

	// Use this for connecting to MongoDB Atlas
	clientOptions := options.Client().ApplyURI(uri)
	client, err := mongo.Connect(context.TODO(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	// Check the connection
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Connected to MongoDB!")

	// collection := client.Database("GolangImageTest").Collection("images")

	return client
}

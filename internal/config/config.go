package config

import "os"

type AppConfig struct {
	MongoURI      string
	MongoDatabase string
}

func New() AppConfig {
	mongoURI := os.Getenv("MONGO_URI")
	mongoDatabase := os.Getenv("MONGO_DATABASE")

	return AppConfig{
		MongoURI:      mongoURI,
		MongoDatabase: mongoDatabase,
	}
}

package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// MongoFields is the struct that defines the fields in the MongoDB database
type MongoFields struct {
	Id       primitive.ObjectID `bson:"_id,omitempty" json:"id,omitempty"`
	Name     string             `bson:"name" json:"name"`
	Data     []byte             `bson:"data" json:"data"`
	Type     string             `bson:"type" json:"type"`
	Uploaded time.Time          `bson:"uploaded" json:"uploaded"`
}

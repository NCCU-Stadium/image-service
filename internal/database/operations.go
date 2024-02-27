package database

import (
	"context"
	"image-service/internal/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func (d *Database) toObjectID(id string) (primitive.ObjectID, error) {
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return primitive.ObjectID{}, err
	}
	return objectId, nil
}

func (d *Database) Insert(imageDocument models.MongoFields, collectionName string) (string, error) {
	storedImage, err := d.db.Collection(collectionName).InsertOne(context.Background(), imageDocument)
	if err != nil {
		return "", err
	}
	primitiveId := storedImage.InsertedID.(primitive.ObjectID)
	return primitiveId.Hex(), nil
}

func (d *Database) FindById(id string, collectionName string) (models.MongoFields, error) {
	docId, err := d.toObjectID(id)
	if err != nil {
		return models.MongoFields{}, err
	}
	var result models.MongoFields
	err = d.db.Collection(collectionName).FindOne(context.Background(), bson.M{"_id": docId}).Decode(&result)
	if err != nil {
		return models.MongoFields{}, err
	}
	return result, nil
}

func (d *Database) DeleteById(id string, collectionName string) (int64, error) {
	docId, err := d.toObjectID(id)
	if err != nil {
		return -1, err
	}
	result, err := d.db.Collection(collectionName).DeleteOne(context.Background(), bson.M{"_id": docId})
	if err != nil {
		return -1, err
	}
	return result.DeletedCount, nil
}

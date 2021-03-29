package repository

import (
	"context"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// MongoUserRepository gives mongo implementation for user operations
type MongoUserRepository struct {
	DB *mongo.Database
}

// AddUser adds new user to the database
func (usr *MongoUserRepository) AddUser(ctx context.Context, user UserEntity) (string, error) {
	result, err := usr.DB.Collection("users").InsertOne(context.Background(), user)
	if err != nil {
		return "", err
	}

	id := result.InsertedID.(primitive.ObjectID).Hex()
	return id, nil
}

// GetUser fetches user info from the database
func (usr *MongoUserRepository) GetUser(ctx context.Context, id string) (UserDocument, error) {
	var userDoc UserDocument

	objID, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return userDoc, err
	}

	err = usr.DB.Collection("users").FindOne(context.Background(), bson.M{"_id": objID}).Decode(&userDoc)

	return userDoc, err
}

// UpdateUser updates user data in the database
func (usr *MongoUserRepository) UpdateUser(ctx context.Context, id string, user UserEntity) error {
	var userDoc UserDocument

	objID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return err
	}

	update := bson.M{
		"$set": user,
	}

	err = usr.DB.Collection("users").FindOneAndUpdate(context.Background(), bson.M{"_id": objID}, update).Decode(&userDoc)

	return err
}

// DeleteUser deletes user document for given id
func (usr *MongoUserRepository) DeleteUser(ctx context.Context, id string) error {
	objID, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return err
	}

	_, err = usr.DB.Collection("users").DeleteOne(ctx, bson.M{"_id": objID})
	return err
}

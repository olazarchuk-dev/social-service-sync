package register

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"social-service-sync/server/model/entity"
)

func CreateUser(ctx context.Context, collection *mongo.Collection, user entity.User) (string, error) {
	result, err := collection.InsertOne(ctx, user)
	if err != nil {
		return "0", err
	}

	/**
	 * get the inserted ID string
	 */
	oid, _ := result.InsertedID.(primitive.ObjectID)
	return fmt.Sprintf("%v", oid.Hex()), err
}

func GetUser(ctx context.Context, collection *mongo.Collection, id string) (entity.User, error) {
	var u entity.User
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return u, err
	}

	err = collection.
		FindOne(ctx, bson.D{{"_id", objectId}}).
		Decode(&u)
	if err != nil {
		return u, err
	}
	return u, nil
}

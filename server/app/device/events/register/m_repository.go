package register

import (
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"social-service-sync/server/model/entity"
)

func CreateUser(u entity.User) (string, error) {
	result, err := UsersCollection.InsertOne(Ctx, u)
	if err != nil {
		return "0", err
	}

	/**
	 * get the inserted ID string
	 */
	oid, _ := result.InsertedID.(primitive.ObjectID)
	return fmt.Sprintf("%v", oid.Hex()), err
}

func GetUser(id string) (entity.User, error) {
	var u entity.User
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return u, err
	}

	err = UsersCollection.
		FindOne(Ctx, bson.D{{"_id", objectId}}).
		Decode(&u)
	if err != nil {
		return u, err
	}
	return u, nil
}

package login

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"social-service-sync/server/model/entity"
)

type NameUsers struct {
	Name  string `bson:"username"`
	Users []entity.User
}

func RepositoryGet(ctx context.Context, collection *mongo.Collection, username string) (entity.User, error) {
	var u entity.User
	err := collection.
		FindOne(ctx, bson.D{{"username", username}}).
		Decode(&u)
	if err != nil {
		return u, err
	}
	return u, nil
}

func RepositoryFindByName(ctx context.Context, collection *mongo.Collection, name string) ([]entity.User, error) {
	matchStage := bson.D{{"$match", bson.D{{"name", name}}}}

	lookupStage := bson.D{{"$lookup",
		bson.D{{"from", "users"},
			{"localField", "name"},
			{"foreignField", "username"},
			{"as", "users"}}}}

	showLoadedCursor, err := collection.Aggregate(ctx, mongo.Pipeline{matchStage, lookupStage})
	if err != nil {
		return nil, err
	}

	var a []NameUsers
	if err = showLoadedCursor.All(ctx, &a); err != nil { // https://jira.mongodb.org/browse/GODRIVER-1129
		return nil, err
	}
	return a[0].Users, err
}

package login

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"social-service-sync/server/model/entity"
)

type NameUsers struct {
	Name  string `bson:"username"`
	Users []entity.User
}

func FindNameUsers(name string) ([]entity.User, error) {
	matchStage := bson.D{{"$match", bson.D{{"name", name}}}}

	lookupStage := bson.D{{"$lookup",
		bson.D{{"from", "users"},
			{"localField", "name"},
			{"foreignField", "username"},
			{"as", "users"}}}}

	showLoadedCursor, err := UsersCollection.Aggregate(Ctx, mongo.Pipeline{matchStage, lookupStage})
	if err != nil {
		return nil, err
	}

	var a []NameUsers
	if err = showLoadedCursor.All(Ctx, &a); err != nil { // https://jira.mongodb.org/browse/GODRIVER-1129
		return nil, err
	}
	return a[0].Users, err
}

func GetUser(username string) (entity.User, error) {
	var u entity.User
	err := UsersCollection.FindOne(Ctx, bson.D{{"username", username}}).Decode(&u)
	if err != nil {
		return u, err
	}
	return u, nil
}

//func GetUsers() ([]entity.User, error) {
//	var user entity.User
//	var users []entity.User
//
//	cursor, err := UsersCollection.Find(Ctx, bson.D{})
//	if err != nil {
//		defer cursor.Close(Ctx)
//		return users, err
//	}
//
//	for cursor.Next(Ctx) {
//		err := cursor.Decode(&user)
//		if err != nil {
//			return users, err
//		}
//		users = append(users, user)
//	}
//
//	return users, nil
//}

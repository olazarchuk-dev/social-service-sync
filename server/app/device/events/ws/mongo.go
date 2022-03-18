package ws

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
)

type Mongo struct {
	Ctx        context.Context   `json:"ctx"`
	Collection *mongo.Collection `json:"collection"`
}

func NewMongo(mongoDb *mongo.Database) *Mongo {
	return &Mongo{
		Ctx:        context.TODO(),
		Collection: mongoDb.Collection("social_settings"),
	}
}

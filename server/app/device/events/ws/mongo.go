package ws

import (
	"context"
	"go.mongodb.org/mongo-driver/mongo"
)

type Mongo struct {
	Ctx        context.Context   `json:"ctx"`
	Collection *mongo.Collection `json:"collection"`
}

func NewMongo(db *mongo.Database) *Mongo {
	return &Mongo{
		Ctx:        context.TODO(),
		Collection: db.Collection("social_settings"),
	}
}

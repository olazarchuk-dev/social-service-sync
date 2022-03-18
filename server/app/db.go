package app

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"social-service-sync/server/app/config"
)

var (
	db  *mongo.Database
	ctx = context.TODO()
)

func init() {

	defer func() {
		if r := recover(); r != nil {
			fmt.Println("recovered in f", r)
		}
	}()

	config, err := config.LoadConfig()
	if err != nil {
		log.Fatal(err)
	}

	connectionURI := "mongodb://" + config.MongoHost + ":" + config.MongoPort + "/"
	clientOptions := options.Client().ApplyURI(connectionURI)
	client, err := mongo.Connect(ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	db = client.Database(config.MongoDatabase)
}

func DbConn() *mongo.Database {
	return db
}

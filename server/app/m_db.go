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
	MongoDb  *mongo.Database
	MongoCtx = context.TODO()
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
	client, err := mongo.Connect(MongoCtx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}
	err = client.Ping(MongoCtx, nil)
	if err != nil {
		log.Fatal(err)
	}

	MongoDb = client.Database(config.MongoDatabase)
}

func MongoConn() *mongo.Database {
	return MongoDb
}

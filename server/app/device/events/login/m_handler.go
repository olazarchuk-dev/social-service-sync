package login

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
	"social-service-sync/server/app/config"
)

var (
	UsersCollection *mongo.Collection
	Ctx             = context.TODO()
)

/*Setup opens a database connection to mongodb*/
func Setup() {

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
	client, err := mongo.Connect(Ctx, clientOptions)
	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(Ctx, nil)
	if err != nil {
		log.Fatal(err)
	}

	db := client.Database(config.MongoDatabase)
	UsersCollection = db.Collection("users")
}

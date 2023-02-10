package database

import (
	"context"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func MongoDB() *mongo.Client {
	MONGOURI := os.Getenv("MONGO_URI")

	client, ok := mongo.Connect(context.TODO(), options.Client().ApplyURI(MONGOURI))

	if ok != nil {
		panic(ok)
	}

	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}

	return client
}

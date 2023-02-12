package database

import (
	"context"
	"os"

	"github.com/redis/go-redis/v9"
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

func Redis() *redis.Client {

	REDIS_ADDRESS := os.Getenv("REDIS_ADDRESS")
	REDIS_USERNAME := os.Getenv("REDIS_USERNAME")
	REDIS_PASSWORD := os.Getenv("REDIS_PASSWORD")

	return redis.NewClient(&redis.Options{

		Addr:     REDIS_ADDRESS,
		Username: REDIS_USERNAME,
		Password: REDIS_PASSWORD,
		DB:       0,
	})
}

package controllers

import (
	ctx "context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gokul1630/go-url-shortner/services"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func Redirect(context *gin.Context) {

	paramHash := context.Param("url")

	mongoClient := context.MustGet("mongoDB").(*mongo.Client)

	redisClient := context.MustGet("redis").(*redis.Client)

	collection := mongoClient.Database("url-schema").Collection("url")

	if value, _ := redisClient.Get(ctx.Background(), paramHash).Result(); value != "" {

		redirect(context, value)

	} else {

		findHash := bson.D{primitive.E{Key: "hash", Value: paramHash}}

		var result UrlSchema

		services.FindOneFromDB(context, collection, findHash).Decode(&result)

		if err := redisClient.Set(ctx.Background(), result.Hash, result.Url, time.Minute*10).Err(); err != nil {
			panic(nil)
		}

		redirect(context, result.Url)
	}
}

func redirect(context *gin.Context, value string) {

	context.Redirect(http.StatusPermanentRedirect, value)

}

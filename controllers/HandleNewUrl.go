package controllers

import (
	ctx "context"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/gokul1630/go-url-shortner/services"
	"github.com/gokul1630/go-url-shortner/utils"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

type Data struct {
	Url string `json:"url"`
}

type UrlSchema struct {
	Url  string
	Hash string
}

func HandleNewUrl(context *gin.Context) {

	var decodedUrl Data

	hash := utils.GenerateHash(10)

	context.BindJSON(&decodedUrl)

	mongoClient := context.MustGet("mongoDB").(*mongo.Client)

	redisClient := context.MustGet("redis").(*redis.Client)

	if value, _ := redisClient.Get(ctx.Background(), decodedUrl.Url).Result(); value != "" {

		sendResponse(context, value)

	} else {
		collection := mongoClient.Database("url-schema").Collection("url")

		findHash := bson.D{primitive.E{Key: "url", Value: decodedUrl.Url}}

		var result UrlSchema

		services.FindOneFromDB(context, collection, findHash).Decode(&result)

		data := UrlSchema{Hash: hash, Url: decodedUrl.Url}

		storeCache(redisClient, decodedUrl.Url, result.Hash)

		if result.Hash != "" {
			storeCache(redisClient, result.Hash, decodedUrl.Url)
		} else {
			storeCache(redisClient, hash, decodedUrl.Url)
		}

		if hash != "" && result.Url != data.Url {
			_, ok := services.InsertIntoDB(context, collection, data)

			if ok != nil {
				panic(ok)
			}

			sendResponse(context, hash)
		} else {
			sendResponse(context, result.Hash)
		}
	}

}

func sendResponse(context *gin.Context, result any) {
	context.JSON(http.StatusOK, gin.H{"url": result})
}

func storeCache(redisClient *redis.Client, key string, value any) {

	if err := redisClient.Set(ctx.Background(), key, value, time.Minute*10).Err(); err != nil {
		panic(nil)
	}
}

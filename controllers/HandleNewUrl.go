package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gokul1630/go-url-shortner/services"
	"github.com/gokul1630/go-url-shortner/utils"
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

	client := context.MustGet("database").(*mongo.Client)

	collection := client.Database("url-schema").Collection("url")

	findHash := bson.D{primitive.E{Key: "url", Value: decodedUrl.Url}}

	var result UrlSchema

	services.FindOneFromDB(context, collection, findHash).Decode(&result)

	data := UrlSchema{Hash: hash, Url: decodedUrl.Url}

	if hash != "" && result.Url != data.Url {
		_, ok := services.InsertIntoDB(context, collection, data)

		if ok != nil {
			panic(ok)
		}

		context.JSON(http.StatusOK, gin.H{"url": hash})
	} else {
		context.JSON(http.StatusOK, gin.H{"url": result.Hash})
	}

}

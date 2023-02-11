package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/gokul1630/go-url-shortner/services"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

func Redirect(context *gin.Context) {

	paramHash := context.Param("url")

	client := context.MustGet("database").(*mongo.Client)

	collection := client.Database("url-schema").Collection("url")

	findHash := bson.D{primitive.E{Key: "hash", Value: paramHash}}

	var result UrlSchema

	services.FindOneFromDB(context, collection, findHash).Decode(&result)

	context.Redirect(http.StatusPermanentRedirect, result.Url)

}

package main

import (
	"context"
	"crypto/rand"
	"fmt"
	"log"
	"math/big"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

type Data struct {
	Url string `json:"url"`
}

type UrlSchema struct {
	Url  string
	Hash string
}

func main() {

	MONGOURI := os.Getenv("MONGO_URI")

	PORT := os.Getenv("PORT")

	client, error := mongo.Connect(context.TODO(), options.Client().ApplyURI(MONGOURI))

	err(error)

	if err := client.Ping(context.TODO(), readpref.Primary()); err != nil {
		panic(err)
	}

	router := gin.Default()

	router.Use(DbMiddleWare(client))

	router.GET("/", func(context *gin.Context) {
		context.File("./index.html")
	})

	router.GET("/:url", func(context *gin.Context) {

		paramHash := context.Param("url")

		collection := client.Database("url-schema").Collection("url")

		findHash := bson.D{primitive.E{Key: "hash", Value: paramHash}}

		var result UrlSchema

		collection.FindOne(context, findHash).Decode(&result)

		context.Redirect(http.StatusPermanentRedirect, result.Url)

	})

	router.POST("/url", handleNewUrl)

	log.Fatal(router.Run(fmt.Sprint(":", PORT)))

	defer func() {
		error := client.Disconnect(context.TODO())
		err(error)
	}()
}

func handleNewUrl(context *gin.Context) {

	var decodedUrl Data

	hash := generateUrl(10)

	context.BindJSON(&decodedUrl)

	client := context.MustGet("database").(*mongo.Client)

	collection := client.Database("url-schema").Collection("url")

	findHash := bson.D{primitive.E{Key: "url", Value: decodedUrl.Url}}

	var result UrlSchema

	collection.FindOne(context, findHash).Decode(&result)

	data := UrlSchema{Hash: hash, Url: decodedUrl.Url}

	if hash != "" && result.Url != data.Url {
		_, error := collection.InsertOne(context, data)
		err(error)

		context.JSON(http.StatusOK, gin.H{"url": hash})
	} else {
		context.JSON(http.StatusOK, gin.H{"url": result.Hash})
	}

}

func generateUrl(n int) string {

	generatedString := make([]byte, n)

	for i := range generatedString {
		randomInt, ok := rand.Int(rand.Reader, big.NewInt(int64(len(letters))))

		err(ok)

		generatedString[i] = letters[randomInt.Int64()]

	}

	return string(generatedString)
}

func err(ok any) {
	if ok != nil {
		panic(ok)
	}
}

func DbMiddleWare(client *mongo.Client) gin.HandlerFunc {
	return func(context *gin.Context) {
		context.Set("database", client)
		context.Next()
	}
}

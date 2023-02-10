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
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

type Data struct {
	Url string `json:"url"`
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

	router.GET("/", func(context *gin.Context) {
		context.File("./index.html")
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

	context.BindJSON(&decodedUrl)

	context.JSON(http.StatusOK, gin.H{"url": generateUrl(10)})
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

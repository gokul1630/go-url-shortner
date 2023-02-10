package main

import (
	"crypto/rand"
	"log"
	"math/big"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

const letters = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

type Data struct {
	Url string `json:"url"`
}

func main() {

	ok := godotenv.Load()
	err(ok)

	router := gin.Default()

	router.GET("/", func(context *gin.Context) {
		context.File("./index.html")
	})

	router.POST("/url", handleNewUrl)

	log.Fatal(router.Run(":3000"))

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

package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/gokul1630/go-url-shortner/database"
	"github.com/gokul1630/go-url-shortner/middlewares"
	"github.com/gokul1630/go-url-shortner/routes"
	// "github.com/joho/godotenv"
)

type Data struct {
	Url string `json:"url"`
}

type UrlSchema struct {
	Url  string
	Hash string
}

func main() {

	// godotenv.Load()

	PORT := os.Getenv("PORT")

	router := gin.Default()

	mongoClient := database.MongoDB()

	redisClient := database.Redis()

	router.Use(middlewares.DBMiddleWare(mongoClient, redisClient))

	routes.Routes(router)

	log.Fatal(router.Run(fmt.Sprint(":", PORT)))

	defer func() {
		ok := mongoClient.Disconnect(context.TODO())
		if ok != nil {
			panic(ok)
		}
	}()
}

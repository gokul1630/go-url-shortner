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
)

type Data struct {
	Url string `json:"url"`
}

type UrlSchema struct {
	Url  string
	Hash string
}

func main() {

	PORT := os.Getenv("PORT")

	router := gin.Default()

	client := database.MongoDB()

	router.Use(middlewares.DBMiddleWare(client))

	routes.Routes(router)

	log.Fatal(router.Run(fmt.Sprint(":", PORT)))

	defer func() {
		ok := client.Disconnect(context.TODO())
		if ok != nil {
			panic(ok)
		}
	}()
}

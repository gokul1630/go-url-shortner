package middlewares

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func DBMiddleWare(client *mongo.Client) gin.HandlerFunc {
	return func(context *gin.Context) {
		context.Set("database", client)
		context.Next()
	}
}

package middlewares

import (
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"go.mongodb.org/mongo-driver/mongo"
)

func DBMiddleWare(client *mongo.Client, redisClient *redis.Client) gin.HandlerFunc {
	return func(context *gin.Context) {
		context.Set("mongoDB", client)
		context.Set("redis", redisClient)
		context.Next()
	}
}

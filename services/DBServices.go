package services

import (
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

func InsertIntoDB(context *gin.Context, collection *mongo.Collection, data any) (*mongo.InsertOneResult, error) {
	return collection.InsertOne(context, data)
}

func FindOneFromDB(context *gin.Context, collection *mongo.Collection, data any) *mongo.SingleResult {
	return collection.FindOne(context, data)
}

package controllers

import (
	"github.com/gin-gonic/gin"
)

func Home(context *gin.Context) {
	context.File("./index.html")
}

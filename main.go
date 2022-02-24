package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	route := gin.Default()
	route.GET("/", func(context *gin.Context) {
		context.JSON(http.StatusOK, gin.H{
			"message": "successful",
		})
	})
	route.Run("localhost:8080")
}

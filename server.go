package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	server.GET("/test", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "OK!!!",
		})
	})

	server.Run("localhost:8080")
}

package main

import (
	"github.com/daffafaizan/go-gin-tutorial/controller"
	"github.com/daffafaizan/go-gin-tutorial/service"
	"github.com/gin-gonic/gin"
)

var (
	videoService    service.VideoService       = service.New()
	videoController controller.VideoController = controller.New(videoService)
)

func main() {
	server := gin.Default()

	server.GET("/posts", func(c *gin.Context) {
		c.JSON(200, videoController.FindAll())
	})

	server.POST("/posts", func(c *gin.Context) {
		c.JSON(200, videoController.Save(c))
	})

	server.Run("localhost:8080")
}

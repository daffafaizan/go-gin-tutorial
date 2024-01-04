package controller

import (
	"github.com/daffafaizan/go-gin-tutorial/entity"
	"github.com/daffafaizan/go-gin-tutorial/service"
	"github.com/gin-gonic/gin"
)

type VideoController interface {
	Save(c *gin.Context) entity.Video
	FindAll() []entity.Video
}

type videoController struct {
	service service.VideoService
}

func New(service service.VideoService) VideoController {
	return videoController{
		service: service,
	}
}

func (controller videoController) Save(c *gin.Context) entity.Video {
	var video entity.Video
	c.BindJSON(&video)
	controller.service.Save(video)
	return video
}

func (controller videoController) FindAll() []entity.Video {
	return controller.service.FindAll()
}

package controller

import (
	"github.com/daffafaizan/go-gin-tutorial/entity"
	"github.com/daffafaizan/go-gin-tutorial/service"
	"github.com/daffafaizan/go-gin-tutorial/validators"
	"github.com/gin-gonic/gin"
	"github.com/go-playground/validator/v10"
)

type VideoController interface {
	Save(c *gin.Context) error
	FindAll() []entity.Video
	ShowAll(c *gin.Context)
}

type videoController struct {
	service service.VideoService
}

var validate *validator.Validate

func New(service service.VideoService) VideoController {
	validate = validator.New()
	validate.RegisterValidation("is-programming", validators.ValidateIsProgramming)
	return videoController{
		service: service,
	}
}

func (controller videoController) Save(c *gin.Context) error {
	var video entity.Video
	err := c.ShouldBindJSON(&video)
	if err != nil {
		return err
	}
	err = validate.Struct(video)
	if err != nil {
		return err
	}
	controller.service.Save(video)
	return nil
}

func (controller videoController) FindAll() []entity.Video {
	return controller.service.FindAll()
}

func (controller videoController) ShowAll(c *gin.Context) {
	videos := controller.service.FindAll()
	data := gin.H{
		"title":  "Video Page",
		"videos": videos,
	}
	c.HTML(200, "index.html", data)
}

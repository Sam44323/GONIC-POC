package controllers

import (
	"github.com/Sam44323/gin-POC/models"
	"github.com/Sam44323/gin-POC/service"
	"github.com/gin-gonic/gin"
)

type VideoController interface {
	FindAll() []models.Video
	Save(ctx *gin.Context)
}

type controller struct {
	service service.VideoService
}

func New(service service.VideoService) VideoController {
	return controller{
		service: service,
	}
}

func FindAll() []models.Video {
	return service.FindAll()
}

func Save(ctx *gin.Context) {
	var video models.Video
	ctx.BindJSON(&video)
	service.Save(video)
}

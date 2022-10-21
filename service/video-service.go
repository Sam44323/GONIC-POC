package service

import "github.com/Sam44323/gin-POC/models";

type VideoService interface {
  Save(video models.Video) models.Video
}

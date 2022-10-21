package service

import "github.com/Sam44323/gin-POC/models";

type VideoService interface {
  Save(video models.Video) models.Video
  FindAll() []models.Video
};

type videoService struct{
  videos []models.Video
};

func New() VideoService{
  return &videoService{};
}

// implementing functions for the VideoService interface

func (service *videoService) Save(video models.Video) models.Video{
  service.videos = append(service.videos, video);
  return video;
}

func (service *videoService) FindAll() []models.Video{
  return service.videos;
}

package main

import (
  "github.com/gin-gonic/gin"
  "github.com/Sam44323/gin-POC/service"
  "github.com/Sam44323/gin-POC/controllers"
  );

var(
  videoService service.VideoService = service.New()
  videoController controllers.VideoController = controllers.New(videoService)
  )

func main(){
  server := gin.Default();

  server.GET("/videos", func(ctx *gin.Context){
    ctx.JSON(200, videoController.FindAll())
  });

  server.POST("/videos", func(ctx *gin.Context){
    ctx.JSON(200, videoController.Save(ctx))
  })


  server.Run(":8080");
}

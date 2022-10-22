package controllers

import (
	"github.com/Sam44323/gin-POC/services"
	"github.com/gin-gonic/gin"
);

type UserController struct{
  UserService services.UserService
}

func New(userserive services.UserService) UserController {
  return UserController{UserService: userserive}
}


func (u *UserController) CreateUser(ctx *gin.Context){
  ctx.JSON(200, nil);
}

func (u *UserController) GetUser(ctx *gin.Context){
  ctx.JSON(200, nil);
}

func (u *UserController) GetAll(ctx *gin.Context) {
  ctx.JSON(200, nil);
}

func (u *UserController) UpdateUser(ctx *gin.Context) {
  ctx.JSON(200, nil);
}

func (u *UserController) DeleteUser(ctx *gin.Context) {
  ctx.JSON(200, nil);
}

func (u * UserController) RegisterUserRoutes(rg *gin.RouterGroup){
  userRoute := rg.Group("/user"); // creating a base-path for the user-route
  userRoute.POST("/create", u.CreateUser);
  userRoute.GET("/get/:name", u.GetUser);
  userRoute.GET("/getAll", u.GetAll);
  userRoute.PUT("/update/:name", u.UpdateUser);
  userRoute.DELETE("/delete/:name", u.DeleteUser);
}

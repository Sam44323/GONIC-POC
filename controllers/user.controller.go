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
}

func (u *UserController) GetUser(*gin.Context){
}

func (u *UserController) GetAll(*gin.Context) {
}

func (u *UserController) UpdateUser(*gin.Context) {
}

func (u *UserController) DeleteUser(*gin.Context) {
}

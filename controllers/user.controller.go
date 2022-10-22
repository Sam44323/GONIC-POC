package controllers

import (
  "github.com/Sam44323/gin-POC/services"
  );

type UserController struct{
  UserService services.UserService
}

func New(userserive services.UserService) UserController {
  return UserController{UserService: userserive}
}

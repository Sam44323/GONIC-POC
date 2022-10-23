package main

//go:generate go run main.go

import (
	"context"

	"github.com/Sam44323/gin-POC/controllers"
	"github.com/Sam44323/gin-POC/services"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
)

var (
	server         *gin.Engine
	userservice    services.UserService
	usercontroller controllers.UserController
	ctx            context.Context
	usercollection *mongo.Collection
	mongoclient    *mongo.Client
	err            error
)

func main() {
	server := gin.Default()

	server.Run(":8080")
}

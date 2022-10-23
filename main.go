package main

//go:generate go run main.go

import (
	"context"
	"fmt"
	"log"

	"github.com/Sam44323/gin-POC/controllers"
	"github.com/Sam44323/gin-POC/services"
	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
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

func init() {
	ctx = context.TODO() // a simple context creation with no cancellation or timeouts
	mongoconnection := options.Client().ApplyURI("mongodb://localhost:27017")
	mongoclient, err = mongo.Connect(ctx, mongoconnection)

	if err != nil {
		log.Fatal(err)
	}

	err = mongoclient.Ping(ctx, readpref.Primary())
	if err != nil {
		log.Fatal(err) // ping the server to verify that the connection is established to the primary-database(not the read-replica) or not
	}

	fmt.Println("Connected to MongoDB!")

	usercollection := mongoclient.Database("userdb").Collection("users")
	userservice := services.NewUserServiceImpl(usercollection, ctx)
	controllers.New(userservice)
	server := gin.Default()
	log.Fatal(server.Run(":8080"))
}

func main() {
	defer mongoclient.Disconnect(ctx) // disconnect if the mongo shut-downs

	basepath := server.Group("/v1")
	usercontroller.RegisterUserRoutes(basepath)
}

package controllers

import (
	"github.com/Sam44323/gin-POC/models"
	"github.com/Sam44323/gin-POC/services"
	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserService services.UserService
}

func New(userserive services.UserService) UserController {
	return UserController{UserService: userserive}
}

func (u *UserController) CreateUser(ctx *gin.Context) {
	var user models.User
	err := ctx.ShouldBindJSON(&user) // binding the user object to the data from the model
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}

	e := u.UserService.CreateUser(&user)
	if e != nil {
		ctx.JSON(400, gin.H{"error": e.Error()})
		return
	}
	ctx.JSON(200, gin.H{"message": "User created successfully"})
}

func (u *UserController) GetUser(ctx *gin.Context) {
	name := ctx.Param("name")
	user, err := u.UserService.GetUser(&name)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{
		"message": "User found!",
		"user":    user,
	})
}

func (u *UserController) GetAll(ctx *gin.Context) {
	var users []*models.User
	users, err := u.UserService.GetAll()
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{
		"message": "Users found!",
		"users":   users,
	})
}

func (u *UserController) UpdateUser(ctx *gin.Context) {
	var user models.User
	err := ctx.ShouldBindJSON(&user) // binding the user object to the data from the model
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	e := u.UserService.UpdateUser(&user)
	if e != nil {
		ctx.JSON(400, gin.H{"error": e.Error()})
		return
	}
	ctx.JSON(200, gin.H{"message": "User updated successfully"})
}

func (u *UserController) DeleteUser(ctx *gin.Context) {
	name := ctx.Param("name")
	err := u.UserService.DeleteUser(&name)
	if err != nil {
		ctx.JSON(400, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(200, gin.H{"message": "User deleted successfully"})
}

func (u *UserController) RegisterUserRoutes(rg *gin.RouterGroup) {
	userRoute := rg.Group("/user") // creating a base-path for the user-route
	userRoute.POST("/create", u.CreateUser)
	userRoute.GET("/get/:name", u.GetUser)
	userRoute.GET("/getAll", u.GetAll)
	userRoute.PATCH("/update/:name", u.UpdateUser)
	userRoute.DELETE("/delete/:name", u.DeleteUser)
}

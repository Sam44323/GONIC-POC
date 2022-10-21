package main

//go:generate go run main.go

import (
	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	server.Run(":8080")
}

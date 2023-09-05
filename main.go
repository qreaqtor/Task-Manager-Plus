package main

import (
	"log"
	"net/http"
	"task-manager-plus-auth-users/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	acPath := server.Group("/auth")
	ac := controllers.NewAuthController()
	ac.RegisterAuthRoutes(acPath)

	ucPath := server.Group("/users")
	uc := controllers.NewUserController()
	uc.RegisterUserRoutes(ucPath)

	testPath := server.Group("/test")
	testPath.GET("/get-some", func(ctx *gin.Context) { ctx.JSON(http.StatusOK, gin.H{"message": "some text"}) })

	log.Fatal(server.Run(":8082"))
}

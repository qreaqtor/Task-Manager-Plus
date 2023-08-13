package main

import (
	"log"
	"task-manager-plus/controllers"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()

	acPath := server.Group("/auth")
	ac := controllers.NewAuthController()
	ac.RegisterAuthRoutes(acPath)

	ucPath := server.Group("/users")
	ucPath.Use(ac.JwtAuthMiddleware())
	uc := controllers.NewUserController()
	uc.RegisterUserRoutes(ucPath)

	log.Fatal(server.Run(":8080"))
}

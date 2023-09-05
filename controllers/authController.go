package controllers

import (
	"fmt"
	"net/http"
	"task-manager-plus-auth-users/models"
	"task-manager-plus-auth-users/services"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	AuthService services.AuthService
}

func NewAuthController() AuthController {
	return AuthController{
		AuthService: services.NewAuthService(),
	}
}

func (ac *AuthController) RegisterUser(ctx *gin.Context) {
	var user *models.UserCreate
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	err := ac.AuthService.CreateUser(user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (ac *AuthController) Login(c *gin.Context) {
	var input models.LoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	token, err := ac.AuthService.LoginCheck(input)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("username or password is incorrect: %s", err)})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": token})
}

func (ac *AuthController) RegisterAuthRoutes(rg *gin.RouterGroup) {
	rg.POST("/register", ac.RegisterUser)
	rg.POST("/login", ac.Login)
}

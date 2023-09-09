package controllers

import (
	"net/http"
	"task-manager-plus-auth-users/models"
	"task-manager-plus-auth-users/services"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	UserService services.UserService
}

func NewUserController() UserController {
	return UserController{
		UserService: services.NewUserService(),
	}
}

func (uc *UserController) GetUser(ctx *gin.Context) {
	username := ctx.Param("username")
	user, err := uc.UserService.GetUser(username)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, user)
}

func (uc *UserController) UpdateUser(ctx *gin.Context) {
	username := ctx.Param("username")
	var user models.UserUpdate
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	err := uc.UserService.UpdateUser(username, &user)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (uc *UserController) DeleteUser(ctx *gin.Context) {
	username := ctx.Param("username")
	err := uc.UserService.DeleteUser(username)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (uc *UserController) RegisterUserRoutes(rg *gin.RouterGroup) {
	rg.GET("/get/:username", uc.GetUser)
	rg.PATCH("/update/:username", uc.UpdateUser)
	rg.DELETE("/delete/:username", uc.DeleteUser)
}

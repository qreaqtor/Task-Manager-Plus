package controllers

import (
	"net/http"
	"task-manager-plus/models"
	"task-manager-plus/services"

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
	uc.UserService.TestSearch()
	var userId string = ctx.Param("id")
	user, err := uc.UserService.GetUser(userId)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, user)
}

func (uc *UserController) UpdateUser(ctx *gin.Context) {
	var userId string = ctx.Param("id")
	var user models.UserUpdate
	if err := ctx.ShouldBindJSON(&user); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	err := uc.UserService.UpdateUser(userId, &user)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (uc *UserController) DeleteUser(ctx *gin.Context) {
	var userId string = ctx.Param("id")
	err := uc.UserService.DeleteUser(userId)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (uc *UserController) RegisterUserRoutes(rg *gin.RouterGroup) {
	rg.GET("/get/:id", uc.GetUser)
	rg.PATCH("/update/:id", uc.UpdateUser)
	rg.DELETE("/delete/:id", uc.DeleteUser)
}

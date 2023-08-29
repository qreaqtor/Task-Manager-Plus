package controllers

import (
	"net/http"
	"task-manager-plus-auth-users/models"
	"task-manager-plus-auth-users/services"

	"github.com/gin-gonic/gin"
	"go.mongodb.org/mongo-driver/bson/primitive"
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
	userId, _ := primitive.ObjectIDFromHex(ctx.Param("id"))
	user, err := uc.UserService.GetUser(userId)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, user)
}

func (uc *UserController) GetMe(ctx *gin.Context) {
	userId := ctx.MustGet("userId").(primitive.ObjectID)
	user, err := uc.UserService.GetUser(userId)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, user)
}

func (uc *UserController) UpdateUser(ctx *gin.Context) {
	userId := ctx.MustGet("userId").(primitive.ObjectID)
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
	userId := ctx.MustGet("userId").(primitive.ObjectID)
	err := uc.UserService.DeleteUser(userId)
	if err != nil {
		ctx.JSON(http.StatusBadGateway, gin.H{"message": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (uc *UserController) RegisterUserRoutes(rg *gin.RouterGroup) {
	rg.GET("/get/:id", uc.GetUser)
	rg.GET("/get/me", uc.GetMe)
	rg.PATCH("/update/", uc.UpdateUser)
	rg.DELETE("/delete/", uc.DeleteUser)
}

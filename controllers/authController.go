package controllers

import (
	"fmt"
	"net/http"
	"strings"
	"task-manager-plus/models"
	"task-manager-plus/services"

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
	var input *models.RegisterInput
	if err := ctx.ShouldBindJSON(&input); err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	user := models.UserCreate{
		Username: input.Username,
		Password: input.Password,
	}
	err := ac.AuthService.CreateUser(&user)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	ctx.JSON(http.StatusOK, gin.H{"message": "success"})
}

func (ac *AuthController) Login(c *gin.Context) {
	var input *models.LoginInput
	if err := c.ShouldBindJSON(&input); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	token, err := ac.AuthService.LoginCheck(input.Username, input.Password)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": fmt.Sprintf("username or password is incorrect: %s", err)})
		return
	}
	c.JSON(http.StatusOK, gin.H{"token": token})
}

func (ac *AuthController) JwtAuthMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		token := ac.extractToken(c)
		err := ac.AuthService.TokenValid(token)
		if err != nil {
			c.String(http.StatusUnauthorized, "Unauthorized")
			c.Abort()
			return
		}
		userId, err := ac.AuthService.ExtractTokenID(token)
		if err != nil {
			c.String(http.StatusUnauthorized, fmt.Sprintf("error: %s", err.Error()))
			c.Abort()
			return
		}
		if err = ac.AuthService.IsUserExists(userId); err != nil {
			c.String(http.StatusUnauthorized, fmt.Sprintf("error: %s", err.Error()))
			c.Abort()
			return
		}
		c.Set("userId", userId)
		c.Next()
	}
}

// func (ac *AuthController) currentUser(c *gin.Context) {
// 	user_id, err := ac.AuthService.ExtractTokenID(ac.extractToken(c))
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}
// 	user, err := ac.AuthService.GetUserByID(user_id)
// 	if err != nil {
// 		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
// 		return
// 	}
// 	c.JSON(http.StatusOK, gin.H{"message": "success", "data": user})
// }

func (ac *AuthController) extractToken(c *gin.Context) string {
	bearerToken := c.Request.Header.Get("Authorization")
	if len(strings.Split(bearerToken, " ")) == 2 {
		return strings.Split(bearerToken, " ")[1]
	}
	return ""
}

func (ac *AuthController) RegisterAuthRoutes(rg *gin.RouterGroup) {
	rg.POST("/register", ac.RegisterUser)
	rg.POST("/login", ac.Login)
}

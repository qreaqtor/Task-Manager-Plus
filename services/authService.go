package services

import (
	"context"
	"errors"
	"task-manager-plus-auth-users/models"
	"time"

	"github.com/golang-jwt/jwt/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"golang.org/x/crypto/bcrypt"
)

var _ IAuthService = (*AuthService)(nil)

type IAuthService interface {
	CreateUser(user *models.UserCreate) error
}

type AuthService struct {
	users *mongo.Collection
	ctx   *context.Context
}

func NewAuthService() AuthService {
	return AuthService{
		users: usersCollection,
		ctx:   &ctx,
	}
}

func (as *AuthService) CreateUser(userCreate *models.UserCreate) error {
	var user models.User
	query := bson.M{"username": userCreate.Username}
	if err := as.users.FindOne(*as.ctx, query).Decode(&user); err == nil {
		return errors.New("this username already exists")
	}
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userCreate.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	userCreate.Password = string(hashedPassword)
	_, err = as.users.InsertOne(*as.ctx, userCreate)
	return err
}

func (as *AuthService) LoginCheck(loginInput models.LoginInput) (string, error) {
	var user *models.User
	filter := bson.M{"username": loginInput.Username}
	err := as.users.FindOne(*as.ctx, filter).Decode(&user)
	if err != nil {
		return "", err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(loginInput.Password))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}
	token, err := generateToken(user.Username)
	if err != nil {
		return "", err
	}
	return token, nil
}

func generateToken(username string) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["username"] = username
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(TOKEN_HOUR_LIFESPAN)).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(API_SECRET))
}

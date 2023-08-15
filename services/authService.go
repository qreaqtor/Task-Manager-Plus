package services

import (
	"context"
	"errors"
	"fmt"
	"task-manager-plus/models"
	"time"

	"github.com/dgrijalva/jwt-go"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
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

func (as *AuthService) CreateUser(user *models.UserCreate) error {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), bcrypt.DefaultCost)
	if err != nil {
		return err
	}
	user.Password = string(hashedPassword)
	_, err = as.users.InsertOne(*as.ctx, user)
	return err
}

func (as *AuthService) LoginCheck(username string, password string) (string, error) {
	var err error
	var user *models.User
	filter := bson.M{"user_name": username}
	err = as.users.FindOne(*as.ctx, filter).Decode(&user)
	if err != nil {
		return "", err
	}
	err = bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))
	if err != nil && err == bcrypt.ErrMismatchedHashAndPassword {
		return "", err
	}
	token, err := generateToken(user.ID)
	if err != nil {
		return "", err
	}
	return token, nil
}

func generateToken(user_id primitive.ObjectID) (string, error) {
	claims := jwt.MapClaims{}
	claims["authorized"] = true
	claims["user_id"] = user_id
	claims["exp"] = time.Now().Add(time.Hour * time.Duration(TOKEN_HOUR_LIFESPAN)).Unix()
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	return token.SignedString([]byte(API_SECRET))
}

func (as *AuthService) TokenValid(token string) error {
	_, err := jwt.Parse(token, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(API_SECRET), nil
	})
	if err != nil {
		return err
	}
	return nil
}

func (as *AuthService) ExtractTokenID(tokenString string) (primitive.ObjectID, error) {
	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(API_SECRET), nil
	})
	if err != nil {
		return primitive.NilObjectID, err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if ok && token.Valid {
		userIdStr, ok := claims["user_id"].(string)
		if !ok {
			return primitive.NilObjectID, fmt.Errorf("user id claim is not a string")
		}
		userId, err := primitive.ObjectIDFromHex(userIdStr)
		if err != nil {
			return primitive.NilObjectID, err
		}
		return userId, nil
	}
	return primitive.NilObjectID, nil
}

func (as *AuthService) IsUserExists(userId primitive.ObjectID) error {
	var user models.UserRead
	query := bson.M{"_id": userId}
	if err := as.users.FindOne(*as.ctx, query).Decode(&user); err != nil {
		return errors.New("current user not found")
	}
	return nil
}

// func (as *AuthService) GetUserByID(userId primitive.ObjectID) (models.UserRead, error) {
// 	var user models.UserRead
// 	query := bson.M{"_id": userId}
// 	if err := as.users.FindOne(*as.ctx, query).Decode(&user); err != nil {
// 		return user, errors.New("user not found")
// 	}
// 	return user, nil
// }

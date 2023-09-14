package services

import (
	"context"
	"errors"
	"task-manager-plus-auth-users/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

var _ IUserService = (*UserService)(nil)

type IUserService interface {
	GetUser(username string) (*models.UserRead, error)
	UpdateUser(username string, user *models.UserUpdate) error
	DeleteUser(username string) error
}

type UserService struct {
	users *mongo.Collection
	ctx   *context.Context
}

func NewUserService() UserService {
	return UserService{
		users: usersCollection,
		ctx:   &ctx,
	}
}

func (us *UserService) GetUser(username string) (*models.UserRead, error) {
	var user *models.UserRead
	filter := bson.M{"username": username}
	err := us.users.FindOne(*us.ctx, filter).Decode(&user)
	return user, err
}

func (us *UserService) UpdateUser(username string, user *models.UserUpdate) error {
	filter := bson.M{"username": username}
	updateBSON := bson.M{"$set": user}
	result, err := us.users.UpdateOne(*us.ctx, filter, updateBSON)
	if err != nil {
		return err
	}
	if result.MatchedCount != 1 {
		return errors.New("no matched document found for update")
	}
	return nil
}

func (us *UserService) DeleteUser(username string) error {
	filter := bson.M{"username": username}
	result, err := us.users.DeleteOne(*us.ctx, filter)
	if err != err {
		return err
	}
	if result.DeletedCount != 1 {
		return errors.New("no matched document found for delete")
	}
	return nil
}

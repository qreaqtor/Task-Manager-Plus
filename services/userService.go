package services

import (
	"context"
	"errors"
	"task-manager-plus/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var _ IUserService = (*UserService)(nil)

type IUserService interface {
	GetUser(id primitive.ObjectID) (*models.UserRead, error)
	UpdateUser(id primitive.ObjectID, user *models.UserUpdate) error
	DeleteUser(id primitive.ObjectID) error
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

func (us *UserService) GetUser(userId primitive.ObjectID) (*models.UserRead, error) {
	var user *models.UserRead
	filter := bson.M{"_id": userId}
	err := us.users.FindOne(*us.ctx, filter).Decode(&user)
	return user, err
}

func (us *UserService) UpdateUser(userId primitive.ObjectID, user *models.UserUpdate) error {
	filter := bson.M{"_id": userId}
	update, err := bson.Marshal(user)
	if err != err {
		return err
	}
	var updateBSON bson.M
	err = bson.Unmarshal(update, &updateBSON)
	if err != nil {
		return err
	}
	updateBSON = bson.M{"$set": updateBSON}
	result, err := us.users.UpdateOne(*us.ctx, filter, updateBSON)
	if err != nil {
		return err
	}
	if result.MatchedCount != 1 {
		return errors.New("no matched document found for update")
	}
	return nil
}

func (us *UserService) DeleteUser(userId primitive.ObjectID) error {
	filter := bson.M{"_id": userId}
	result, err := us.users.DeleteOne(*us.ctx, filter)
	if err != err {
		return err
	}
	if result.DeletedCount != 1 {
		return errors.New("no matched document found for delete")
	}
	return nil
}

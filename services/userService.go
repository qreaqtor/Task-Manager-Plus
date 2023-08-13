package services

import (
	"context"
	"errors"
	"fmt"
	"log"
	"task-manager-plus/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

var _ IUserService = (*UserService)(nil)

type IUserService interface {
	GetUser(id string) (*models.UserRead, error)
	UpdateUser(id string, user *models.UserUpdate) error
	DeleteUser(id string) error
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

func (us *UserService) GetUser(userIdStr string) (*models.UserRead, error) {
	var user *models.UserRead
	userId, err := primitive.ObjectIDFromHex(userIdStr)
	if err != nil {
		return user, err
	}
	filter := bson.M{"_id": userId}
	err = us.users.FindOne(*us.ctx, filter).Decode(&user)
	return user, err
}

func (us *UserService) UpdateUser(userIdStr string, user *models.UserUpdate) error {
	userId, err := primitive.ObjectIDFromHex(userIdStr)
	if err != nil {
		return err
	}
	filter := bson.M{"_id": userId}
	update := bson.M{
		"$set": user.ToBSONM(),
	}
	result, _ := us.users.UpdateOne(*us.ctx, filter, update)
	if result.MatchedCount != 1 {
		return errors.New("no matched document found for update")
	}
	return nil
}

func (us *UserService) DeleteUser(userIdStr string) error {
	userId, err := primitive.ObjectIDFromHex(userIdStr)
	if err != nil {
		return err
	}
	filter := bson.M{"_id": userId}
	result, _ := us.users.DeleteOne(*us.ctx, filter)
	if result.DeletedCount != 1 {
		return errors.New("no matched document found for delete")
	}
	return nil
}

func (us *UserService) TestSearch() {
	model := mongo.IndexModel{
		Keys: bson.M{"user_name": "text"},
	}
	_, err := us.users.Indexes().CreateOne(*us.ctx, model)
	if err != nil {
		panic(err)
	}
	filter := bson.M{"$text": bson.M{"$search": "r"}}
	cursor, err := us.users.Find(context.Background(), filter)
	result := make([]string, 0)
	for cursor.Next(*us.ctx) {
		var user models.User
		err := cursor.Decode(&user)
		if err != nil {
			log.Fatal(err)
		}
		result = append(result, fmt.Sprintf("ID: %s, Username: %s\n", user.ID.Hex(), user.Username))
	}
	fmt.Println(result)
}

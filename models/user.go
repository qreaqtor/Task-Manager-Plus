package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID       primitive.ObjectID `json:"id" bson:"_id" binding:"required"`
	Username string             `json:"username" bson:"user_name" binding:"required"`
	Password string             `json:"password" bson:"password" binding:"required"`
}

type UserCreate struct {
	Username string `json:"username" bson:"username" binding:"required"`
	Password string `json:"password" bson:"password" binding:"required"`
}

type UserRead struct {
	Username string             `json:"username" bson:"username" binding:"required"`
	ID       primitive.ObjectID `json:"id" bson:"_id" binding:"required"`
}

type UserUpdate struct {
	Username string `json:"username" bson:"username" binding:"required"`
}

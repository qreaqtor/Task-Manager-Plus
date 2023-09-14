package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID        primitive.ObjectID `json:"id" bson:"_id" binding:"required"`
	Username  string             `json:"username" bson:"username" binding:"required"`
	FirstName string             `json:"firstName" bson:"firstName" binding:"required"`
	LastName  string             `json:"lastName" bson:"lastName" binding:"required"`
	Password  string             `json:"password" bson:"password" binding:"required"`
}

type UserCreate struct {
	Username  string `json:"username" bson:"username" binding:"required"`
	FirstName string `json:"firstName" bson:"firstName" binding:"required"`
	LastName  string `json:"lastName" bson:"lastName" binding:"required"`
	Password  string `json:"password" bson:"password" binding:"required"`
}

type UserRead struct {
	Username  string `json:"username" bson:"username" binding:"required"`
	FirstName string `json:"firstName" bson:"firstName" binding:"required"`
	LastName  string `json:"lastName" bson:"lastName" binding:"required"`
}

type UserUpdate struct {
	FirstName string `json:"firstName" bson:"firstName,omitempty" binding:"-"`
	LastName  string `json:"lastName" bson:"lastName,omitempty" binding:"-"`
}

package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Task struct {
	UserId  primitive.ObjectID `json:"user_id" bson:"user_id"`
	Title   string             `json:"title" bson:"title" binding:"required"`
	Content string             `json:"content" bson:"content" binding:"required"`
}

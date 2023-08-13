package models

import (
	"reflect"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID       primitive.ObjectID `json:"id" bson:"_id" binding:"required"`
	Username string             `json:"username" bson:"user_name" binding:"required"`
	Password string             `json:"password" bson:"password" binding:"required"`
}

type UserCreate struct {
	Username string `json:"username" bson:"user_name" binding:"required"`
	Password string `json:"password" bson:"password" binding:"required"`
}

type UserRead struct {
	Username string             `json:"username" bson:"user_name" binding:"required"`
	ID       primitive.ObjectID `json:"id" bson:"_id" binding:"required"`
}

type UserUpdate struct {
	Username string `json:"username" bson:"user_name" binding:"required"`
}

func (u *UserUpdate) ToBSONM() bson.M {
	result := bson.M{}
	v := reflect.ValueOf(*u)
	t := v.Type()
	for i := 0; i < t.NumField(); i++ {
		field := t.Field(i)
		value := v.Field(i)
		bsonKey := field.Tag.Get("bson")
		if bsonKey != "" {
			result[bsonKey] = value.Interface()
		}
	}
	return result
}

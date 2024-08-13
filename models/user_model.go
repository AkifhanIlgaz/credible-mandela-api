package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type User struct {
	Id           primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Username     string             `json:"username" bson:"username"`
	PasswordHash string             `json:"-" bson:"passwordHash"`
	Roles        string             `json:"roles" bson:"roles"`
	Address      string             `json:"address" bson:"address"`
}

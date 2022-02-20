package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Unit struct {
	Id   primitive.ObjectID `json:"id" bson:"_id"`
	Name string             `json:"name" bson:"name"`
}

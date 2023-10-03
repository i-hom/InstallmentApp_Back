package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Card struct {
	ID      primitive.ObjectID `bson:"_id"`
	Number  string             `json:"number"`
	ExpDate string             `json:"expDate"`
	Balance int                `json:"balance"`
}

type AddCard struct {
	Number  string             `json:"number"`
	ExpDate string             `json:"expDate"`
	Balance int                `json:"balance"`
	OwnerId primitive.ObjectID `bson:"ownerid"`
}

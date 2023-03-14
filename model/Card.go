package model

import "go.mongodb.org/mongo-driver/bson/primitive"

//==================BSON=======================

type BCard struct {
	ID      primitive.ObjectID `bson:"_id"`
	Number  string             `json:"number"`
	ExpDate string             `json:"expDate"`
	Value   int                `json:"value"`
	ownerID	primitive.ObjectID `bson:"ownerId"`
}

//===================JSON=======================

type JCard struct {
	ID      string `json:"id"`
	Number  string `json:"number"`
	ExpDate string `json:"expDate"`
	Value   int    `json:"value"`
}
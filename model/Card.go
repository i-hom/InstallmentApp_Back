package model

import (
	"context"
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

//==================BSON=======================

type BCard struct {
	ID      primitive.ObjectID `bson:"_id"`
	Number  string             `json:"number"`
	ExpDate string             `json:"expDate"`
	Balance int                `json:"balance"`
}

//===================JSON=======================

type AddCard struct {
	Number  string             `json:"number"`
	ExpDate string             `json:"expDate"`
	Value   int                `json:"value"`
	OwnerId primitive.ObjectID `json:"ownerId"`
}

func CardAdd(params interface{}, db *mongo.Database) RPCResponse {
	var cardData AddCard
	json.Unmarshal(GetRaw(params), &cardData)

	if cardData.ExpDate == "" || cardData.Number == "" {
		return RPCResponse{Error: &RPCError{Code: 1, Message: "Missing one of params"}}
	}

	if len(cardData.Number) != 16 {
		return RPCResponse{Error: &RPCError{Code: 2, Message: "Card number should be 16 digits"}}
	}

	db.Collection("Cards").InsertOne(context.TODO(), cardData)

	return RPCResponse{Result: cardData}
}

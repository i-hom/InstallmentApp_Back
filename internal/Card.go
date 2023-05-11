package internal

import (
	"context"
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
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
		return RPCResponse{Code: 1, Message: "Missing one of params"}
	}

	if len(cardData.Number) != 16 {
		return RPCResponse{Code: 2, Message: "Card number should be 16 digits"}
	}

	db.Collection("Cards").InsertOne(context.TODO(), cardData)

	return RPCResponse{Code: 0, Message: "Card added"}
}

func GetCards(ownerID primitive.ObjectID, db *mongo.Database) []BCard {
	var cards []BCard
	curr, _ := db.Collection("Cards").Find(context.TODO(), bson.M{"ownerid": ownerID})
	curr.All(context.TODO(), &cards)
	return cards
}

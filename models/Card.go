package models

import (
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"installment_back/src"
)

//==================BSON=======================

type Card struct {
	ID      primitive.ObjectID `bson:"_id"`
	Number  string             `json:"number"`
	ExpDate string             `json:"expDate"`
	Balance int                `json:"balance"`
}

//===================JSON=======================

type AddCard struct {
	Number  string             `json:"number"`
	ExpDate string             `json:"expDate"`
	Balance int                `json:"balance"`
	OwnerId primitive.ObjectID `bson:"ownerid"`
}

func (card *Card) Add(params interface{}, db *src.DataBase) src.RPCResponse {
	var cardData AddCard
	json.Unmarshal(src.GetRaw(params), &cardData)

	if cardData.ExpDate == "" || cardData.Number == "" {
		return src.Missing_parameter
	}

	if len(cardData.Number) != 16 {
		return src.Card_number_length_not_correct
	}

	cardData.Balance = 1000000

	db.Insert("Cards", cardData)

	return src.RPCResponse{Code: 0, Message: "Card added"}
}

func CardsGet(ownerID primitive.ObjectID, db *src.DataBase) []Card {
	var cards []Card
	db.FindAll("Cards", bson.M{"ownerid": ownerID}, &cards)
	return cards
}

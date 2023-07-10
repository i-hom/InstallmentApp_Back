package controllers

import (
	"context"
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"installment_back/models"
)

type Card models.Card

func (this *Card) Add(params interface{}, db *DataBase) models.RPCResponse {
	var cardData models.AddCard
	json.Unmarshal(GetRaw(params), &cardData)

	if cardData.ExpDate == "" || cardData.Number == "" {
		return models.Missing_parameter
	}

	if len(cardData.Number) != 16 {
		return models.Card_number_length_not_correct
	}

	cardData.Balance = 1000000

	db.Insert("Cards", cardData)

	return models.RPCResponse{Code: 0, Message: "Card added"}
}

func (this *Card) Get(ownerID primitive.ObjectID, db *DataBase) []models.Card {
	var cards []models.Card
	db.FindAll("Cards", bson.M{"ownerid": ownerID}).All(context.TODO(), &cards)
	return cards
}

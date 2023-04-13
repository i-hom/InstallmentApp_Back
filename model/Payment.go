package model

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
)

type Payment struct {
	ElmakonID string `json:"elmakonid"`
	Amount    int    `json:"amount"`
	Method    string `json:"method"`
	Date      string `json:"date"`
}

func GetPayments(elmakonID string, db *mongo.Database) []Payment {
	var payments []Payment
	curr, _ := db.Collection("Payments").Find(context.TODO(), bson.M{"elmakonid": elmakonID})
	curr.All(context.TODO(), &payments)
	return payments
}

package models

import (
	"go.mongodb.org/mongo-driver/bson"
	"installment_back/src"
)

type Payment struct {
	ElmakonID string `json:"elmakonid"`
	Amount    int    `json:"amount"`
	Method    string `json:"method"`
	Date      string `json:"date"`
}

func PaymentsGet(elmakonID string, db *src.DataBase) []Payment {
	var payments []Payment
	db.FindAll("Payments", bson.M{"elmakonid": elmakonID}, &payments)
	return payments
}

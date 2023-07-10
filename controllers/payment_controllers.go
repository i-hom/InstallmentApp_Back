package controllers

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"installment_back/models"
)

type Payment models.Payment

func (this *Payment) Get(elmakonID string, db *DataBase) []models.Payment {
	var payments []models.Payment
	db.FindAll("Payments", bson.M{"elmakonid": elmakonID}).All(context.TODO(), &payments)
	return payments
}

package models

import "go.mongodb.org/mongo-driver/bson/primitive"

type Payment struct {
	InstallmentID primitive.ObjectID `bson:"installment_id"`
	Amount        int                `json:"amount"`
	CardID        primitive.ObjectID `bson:"method"`
	Date          string             `json:"date"`
}

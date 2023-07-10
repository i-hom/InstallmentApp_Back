package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//==================BSON=======================

type BInstallment struct {
	ID             primitive.ObjectID `bson:"_id"`
	ElmakonID      string             `json:"elmakonid"`
	Balance        int                `json:"balance"`
	IsActive       bool               `json:"isActive"`
	MonthlyPayment int                `json:"monthlyPayment"`
	OwnerID        primitive.ObjectID `json:"ownerId"`
	ItemID         primitive.ObjectID `json:"itemId"`
}

func (installment *BInstallment) ToJInstallment() Installment {
	return Installment{
		ID:             installment.ID,
		ElmakonID:      installment.ElmakonID,
		Balance:        installment.Balance,
		IsActive:       installment.IsActive,
		MonthlyPayment: installment.MonthlyPayment,
	}
}

//===================JSON=======================

type Installment struct {
	ID             primitive.ObjectID `bson:"_id"`
	ElmakonID      string             `json:"elmakonid"`
	Item           Item               `json:"item"`
	Balance        int                `json:"balance"`
	IsActive       bool               `json:"isActive"`
	MonthlyPayment int                `json:"monthlyPayment"`
	Payments       []Payment          `json:"payments"`
}

type InstallmentPayment struct {
	InstallmentID primitive.ObjectID `json:"installment_id"`
	CardID        primitive.ObjectID `json:"card_id"`
	Amount        int                `json:"amount"`
}

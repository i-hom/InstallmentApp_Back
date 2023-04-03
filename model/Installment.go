package model

import (
	"context"
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
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

func (installment *BInstallment) ToJInstallment() JInstallment {
	return JInstallment{
		ID:             installment.ID,
		ElmakonID:      installment.ElmakonID,
		Balance:        installment.Balance,
		IsActive:       installment.IsActive,
		MonthlyPayment: installment.MonthlyPayment,
	}
}

//===================JSON=======================

type JInstallment struct {
	ID             primitive.ObjectID `bson:"_id"`
	ElmakonID      string             `json:"elmakonId"`
	Item           JItem              `json:"item"`
	Balance        int                `json:"balance"`
	IsActive       bool               `json:"isActive"`
	MonthlyPayment int                `json:"monthlyPayment"`
}

type InstallmentPay struct {
	InstallmentID primitive.ObjectID `json:"installmentId"`
	CardID        primitive.ObjectID `json:"number"`
	Value         int                `json:"value"`
}

func InstallmentPayment(params interface{}, db *mongo.Database) RPCResponse {
	var installmentData InstallmentPay
	json.Unmarshal(GetRaw(params), &installmentData)
	if installmentData.InstallmentID.IsZero() || installmentData.CardID.IsZero() {
		return RPCResponse{Error: &RPCError{Code: 1, Message: "Missing one of params"}}
	}
	var cardData BCard
	db.Collection("Cards").FindOne(context.TODO(), bson.M{"_id": installmentData.CardID}).Decode(&cardData)

	var installment BInstallment
	db.Collection("Installments").FindOne(context.TODO(), bson.M{"_id": installmentData.InstallmentID}).Decode(&installment)

	if cardData.Balance < installmentData.Value {
		return RPCResponse{Error: &RPCError{Code: 3, Message: "Insufficient balance"}}
	}

	if installment.Balance < installmentData.Value {
		return RPCResponse{Error: &RPCError{Code: 6, Message: "U paid a lot"}}
	}

	cardData.Balance -= installmentData.Value
	installment.Balance -= installmentData.Value

	_, cardDeposit := db.Collection("Cards").UpdateOne(context.TODO(), bson.M{"_id": installmentData.CardID}, bson.M{"$set": cardData})
	_, installmentDeposit := db.Collection("Installments").UpdateOne(context.TODO(), bson.M{"_id": installmentData.InstallmentID}, bson.M{"$set": installment})

	if cardDeposit != nil || installmentDeposit != nil {
		err := RPCError{Code: 5, Message: "Failed to deposit"}
		return RPCResponse{Error: &err}
	}

	if installment.Balance == 0 {
		installment.IsActive = false
		db.Collection("Installments").UpdateOne(context.TODO(), bson.M{"elmakonid": installment.ID}, bson.M{"$set": installment})
	}

	return RPCResponse{Result: "Successfully paid installment"}
}

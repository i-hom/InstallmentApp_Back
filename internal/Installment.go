package internal

import (
	"context"
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
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
	ElmakonID      string             `json:"elmakonid"`
	Item           JItem              `json:"item"`
	Balance        int                `json:"balance"`
	IsActive       bool               `json:"isActive"`
	MonthlyPayment int                `json:"monthlyPayment"`
	Payments       []Payment          `json:"payments"`
}

type InstallmentPay struct {
	InstallmentID primitive.ObjectID `json:"installment_id"`
	CardID        primitive.ObjectID `json:"card_id"`
	Amount        int                `json:"amount"`
}

func InstallmentPayment(params interface{}, db *mongo.Database) RPCResponse {
	var installmentData InstallmentPay
	json.Unmarshal(GetRaw(params), &installmentData)
	if installmentData.InstallmentID.IsZero() || installmentData.CardID.IsZero() {
		return RPCResponse{Code: 1, Message: "Missing one of params"}
	}
	var cardData BCard
	db.Collection("Cards").FindOne(context.TODO(), bson.M{"_id": installmentData.CardID}).Decode(&cardData)

	var installment BInstallment
	db.Collection("Installments").FindOne(context.TODO(), bson.M{"_id": installmentData.InstallmentID}).Decode(&installment)

	if cardData.Balance < installmentData.Amount {
<<<<<<< refs/remotes/origin/master:model/Installment.go
		return RPCResponse{Error: &RPCError{Code: 3, Message: "Insufficient balance"}}
	}

	if installment.Balance < installmentData.Amount {
		return RPCResponse{Error: &RPCError{Code: 6, Message: "U paid a lot"}}
=======
		return RPCResponse{Code: 3, Message: "Insufficient balance"}
	}

	if installment.Balance < installmentData.Amount {
		return RPCResponse{Code: 6, Message: "U paid a lot"}
>>>>>>> Update logic and cleanup code:internal/Installment.go
	}

	cardData.Balance -= installmentData.Amount
	installment.Balance -= installmentData.Amount

	_, cardDeposit := db.Collection("Cards").UpdateOne(context.TODO(), bson.M{"_id": installmentData.CardID}, bson.M{"$set": cardData})
	_, installmentDeposit := db.Collection("Installments").UpdateOne(context.TODO(), bson.M{"_id": installmentData.InstallmentID}, bson.M{"$set": installment})

	if cardDeposit != nil || installmentDeposit != nil {
		return RPCResponse{Code: 5, Message: "Failed to deposit"}
	}

	if installment.Balance == 0 {
		installment.IsActive = false
		db.Collection("Installments").UpdateOne(context.TODO(), bson.M{"elmakonid": installment.ID}, bson.M{"$set": installment})
	}

	now := time.Now().Format("2006-01-02 15:04:05")

	db.Collection("Payments").InsertOne(context.TODO(), Payment{
		installment.ElmakonID, installmentData.Amount, cardData.Number, now,
	})

<<<<<<< refs/remotes/origin/master:model/Installment.go
	return RPCResponse{Result: "Successfully paid installment"}
=======
	return RPCResponse{Code: 0, Message: "Successfully paid installment"}
>>>>>>> Update logic and cleanup code:internal/Installment.go
}

package models

import (
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"installment_back/src"
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

func (inst *Installment) Pay(params interface{}, db *src.DataBase) src.RPCResponse {
	var installmentData InstallmentPayment
	json.Unmarshal(src.GetRaw(params), &installmentData)

	if installmentData.InstallmentID.IsZero() || installmentData.CardID.IsZero() {
		return src.RPCResponse{Code: 1, Message: "Missing one of params"}
	}
	var cardData Card
	db.FindOne("Cards", bson.M{"_id": installmentData.CardID}, &cardData)

	var installment BInstallment
	db.FindOne("Installments", bson.M{"_id": installmentData.InstallmentID}, &installment)

	if cardData.Balance < installmentData.Amount {
		return src.Insufficient_balance
	}
	if installment.Balance < installmentData.Amount {
		return src.Payment_greater_than_balance
	}

	cardData.Balance -= installmentData.Amount
	installment.Balance -= installmentData.Amount

	installmentDeposit, err := db.Update("Installments", bson.M{"_id": installmentData.InstallmentID}, installment)
	if err != nil {
		return src.Failde_to_deposite
	}
	cardDeposit, err := db.Update("Cards", bson.M{"_id": installmentData.CardID}, cardData)
	if err != nil {
		return src.Failde_to_deposite
	}

	if cardDeposit == nil || installmentDeposit == nil {
		return src.Failde_to_deposite
	}

	if installment.Balance == 0 {
		installment.IsActive = false
		db.Update("Installments", bson.M{"elmakonid": installment.ID}, installment)
	}

	now := time.Now().Format("2006-01-02 15:04:05")

	db.Insert("Payments", Payment{
		installment.ElmakonID, installmentData.Amount, cardData.Number, now,
	})

	return src.RPCResponse{Code: 0, Message: "Successfully paid installment"}
}

func InstallmentsGet(ownerID primitive.ObjectID, db *src.DataBase) ([]Installment, error) {
	var installments []BInstallment
	var jinstallments []Installment
	err := db.FindAll("Installments", bson.M{"ownerid": ownerID}, &installments)
	for _, i := range installments {
		var installment = i.ToJInstallment()
		installment.Item, _ = ItemGet(i.ItemID, db)
		installment.Payments = PaymentsGet(i.ElmakonID, db)
		jinstallments = append(jinstallments, installment)
	}
	return jinstallments, err
}

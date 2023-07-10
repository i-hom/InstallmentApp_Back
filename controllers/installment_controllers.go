package controllers

import (
	"context"
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"installment_back/models"
	"time"
)

type Installment models.Installment

func (inst *Installment) Pay(params interface{}, db *DataBase) models.RPCResponse {
	var installmentData models.InstallmentPayment
	json.Unmarshal(GetRaw(params), &installmentData)

	if installmentData.InstallmentID.IsZero() || installmentData.CardID.IsZero() {
		return models.Missing_parameter
	}
	var cardData models.Card
	db.FindOne("Cards", bson.M{"_id": installmentData.CardID}).Decode(&cardData)

	var installment models.BInstallment
	db.FindOne("Installments", bson.M{"_id": installmentData.InstallmentID}).Decode(&installment)

	if cardData.Balance < installmentData.Amount {
		return models.Insufficient_balance
	}
	if installment.Balance < installmentData.Amount {
		return models.Payment_greater_than_balance
	}

	cardData.Balance -= installmentData.Amount
	installment.Balance -= installmentData.Amount

	installmentDeposit, err := db.Update("Installments", bson.M{"_id": installmentData.InstallmentID}, installment)
	if err != nil {
		return models.Failde_to_deposite
	}
	cardDeposit, err := db.Update("Cards", bson.M{"_id": installmentData.CardID}, cardData)
	if err != nil {
		return models.Failde_to_deposite
	}

	if cardDeposit == nil || installmentDeposit == nil {
		return models.Failde_to_deposite
	}

	if installment.Balance == 0 {
		installment.IsActive = false
		db.Update("Installments", bson.M{"elmakonid": installment.ID}, installment)
	}

	now := time.Now().Format("2006-01-02 15:04:05")

	db.Insert("Payments", models.Payment{
		installment.ElmakonID, installmentData.Amount, cardData.Number, now,
	})

	return models.RPCResponse{Code: 0, Message: "Successfully paid installment"}
}

func (this *Installment) Get(ownerID primitive.ObjectID, db *DataBase) ([]models.Installment, error) {
	var installments []models.BInstallment
	var jinstallments []models.Installment
	var item Item
	var payment Payment
	db.FindAll("Installments", bson.M{"ownerid": ownerID}).All(context.TODO(), &installments)
	for _, i := range installments {
		var installment = i.ToJInstallment()
		installment.Item, _ = item.Get(i.ItemID, db)
		installment.Payments = payment.Get(i.ElmakonID, db)
		jinstallments = append(jinstallments, installment)
	}
	return jinstallments, nil
}

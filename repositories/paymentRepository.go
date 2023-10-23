package repositories

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"installment_back/errors"
	"installment_back/models"
	"installment_back/storage"
	"time"
)

type PaymentRepository struct {
	db *storage.DataBase
}

func NewPaymentRepository(db *storage.DataBase) *PaymentRepository {
	return &PaymentRepository{db: db}
}

func (ip *PaymentRepository) GetAll(installmentID primitive.ObjectID) ([]models.Payment, models.RPCResponse) {
	var payments []models.Payment
	if err := ip.db.FindAll("Payments", bson.M{"installment_id": installmentID}, &payments); err != nil {
		return []models.Payment{}, errors.Payment_not_found
	}
	return payments, errors.Success
}

func (ip *PaymentRepository) Add(paymentData models.InstallmentPayment) models.RPCResponse {

	now := time.Now().Format("2006-01-02 15:04:05")

	err := ip.db.Insert("Payments", models.Payment{
		InstallmentID: paymentData.InstallmentID, Amount: paymentData.Amount, CardID: paymentData.CardID, Date: now,
	})
	if err != nil {
		return errors.Failed_to_add_payment
	}

	return errors.Success
}

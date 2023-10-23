package repositories

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"installment_back/errors"
	"installment_back/models"
	"installment_back/storage"
)

type InstallmentRepository struct {
	db *storage.DataBase
}

func NewInstallmentRepository(db *storage.DataBase) *InstallmentRepository {
	return &InstallmentRepository{db: db}
}

func (ir *InstallmentRepository) GetAll(ownerID primitive.ObjectID) ([]models.BInstallment, models.RPCResponse) {
	var installments []models.BInstallment
	if err := ir.db.FindAll("Installments", bson.M{"ownerid": ownerID}, &installments); err != nil {
		return []models.BInstallment{}, errors.Installment_not_found
	}
	return installments, errors.Success
}

func (ir *InstallmentRepository) Get(installmentID primitive.ObjectID) (models.BInstallment, models.RPCResponse) {
	var installment models.BInstallment
	if err := ir.db.FindOne("Installments", bson.M{"_id": installmentID}, &installment); err != nil {
		return models.BInstallment{}, errors.Installment_not_found
	}
	return installment, errors.Success
}

func (ir *InstallmentRepository) Deposit(installmentID primitive.ObjectID, amount int) models.RPCResponse {

	installment, iErr := ir.Get(installmentID)
	if iErr != errors.Success {
		return iErr
	}

	if installment.Balance < amount {
		return errors.Payment_greater_than_balance
	}

	installment.Balance -= amount

	if installment.Balance == 0 {
		installment.IsActive = false
	}

	installmentDeposit, err := ir.db.Update("Installments", bson.M{"_id": installmentID}, installment)
	if err != nil && installmentDeposit == 0 {
		return errors.Failed_to_deposite
	}
	return errors.Success
}

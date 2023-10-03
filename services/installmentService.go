package services

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"installment_back/errors"
	"installment_back/models"
	"installment_back/repositories"
)

type InstallmentService struct {
	installmentRepository *repositories.InstallmentRepository
	cardService           *CardService
	paymentRepository     *repositories.PaymentRepository
	itemRepository        *repositories.ItemRepository
}

func NewInstallmentService(installmentRepository *repositories.InstallmentRepository, cardService *CardService, paymentRepository *repositories.PaymentRepository, itemRepository *repositories.ItemRepository) *InstallmentService {
	return &InstallmentService{installmentRepository: installmentRepository, cardService: cardService, paymentRepository: paymentRepository, itemRepository: itemRepository}
}

func (is *InstallmentService) Pay(paymentData models.InstallmentPayment) models.RPCResponse {

	cErr := is.cardService.Deposit(paymentData.CardID, paymentData.Amount)
	iErr := is.installmentRepository.Deposit(paymentData.InstallmentID, paymentData.Amount)

	if cErr != errors.Success && iErr != errors.Success {
		return errors.Failed_to_deposite
	}

	return is.paymentRepository.Add(paymentData)
}

func (is *InstallmentService) GetAll(ownerID primitive.ObjectID) ([]models.Installment, models.RPCResponse) {
	var jinstallments []models.Installment
	installments, _ := is.installmentRepository.GetAll(ownerID)
	for _, i := range installments {
		var installment = i.ToJInstallment()
		installment.Item, _ = is.itemRepository.Get(i.ItemID)
		installment.Payments, _ = is.paymentRepository.GetAll(i.ID)
		jinstallments = append(jinstallments, installment)
	}
	return jinstallments, errors.Success
}

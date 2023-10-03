package services

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"installment_back/errors"
	"installment_back/models"
	"installment_back/repositories"
)

type CardService struct {
	cardRepository *repositories.CardRepository
}

func NewCardService(cardRepository *repositories.CardRepository) *CardService {
	return &CardService{cardRepository: cardRepository}
}

func (cs *CardService) Add(cardData models.AddCard) models.RPCResponse {
	cardData.Balance = 1000000

	return cs.cardRepository.Add(cardData)
}

func (cs *CardService) Deposit(cardID primitive.ObjectID, amount int) models.RPCResponse {
	card, err := cs.cardRepository.Get(cardID)
	if err != errors.Success {
		return err
	}
	if card.Balance < amount {
		return errors.Insufficient_balance
	}
	return cs.cardRepository.BalanceOperations(cardID, -amount)
}

func (cs *CardService) Fill(cardID primitive.ObjectID, amount int) models.RPCResponse {
	return cs.cardRepository.BalanceOperations(cardID, amount)
}

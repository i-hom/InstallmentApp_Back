package repositories

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"installment_back/errors"
	"installment_back/models"
	"installment_back/storage"
)

type CardRepository struct {
	db *storage.DataBase
}

func NewCardRepository(db *storage.DataBase) *CardRepository {
	return &CardRepository{db: db}
}

func (cr *CardRepository) Add(cardData models.AddCard) models.RPCResponse {
	if err := cr.db.Insert("Cards", cardData); err != nil {
		return errors.Failed_to_add_card
	}

	return models.RPCResponse{Code: 0, Message: "Card added"}
}

func (cr *CardRepository) GetAll(ownerID primitive.ObjectID) ([]models.Card, models.RPCResponse) {
	var cards []models.Card
	if err := cr.db.FindAll("Cards", bson.M{"ownerid": ownerID}, &cards); err != nil {
		return []models.Card{}, errors.Failed_to_get_cards
	}
	return cards, errors.Success
}

func (cr *CardRepository) Get(cardID primitive.ObjectID) (models.Card, models.RPCResponse) {
	var card models.Card
	if err := cr.db.FindOne("Cards", bson.M{"_id": cardID}, &card); err != nil {
		return models.Card{}, errors.Card_not_found
	}
	return card, errors.Success
}

func (cr *CardRepository) BalanceOperations(cardID primitive.ObjectID, amount int) models.RPCResponse {
	var card models.Card
	if err := cr.db.FindOne("Cards", bson.M{"_id": cardID}, &card); err != nil {
		return errors.Card_not_found
	}
	card.Balance += amount
	if _, err := cr.db.Update("Cards", bson.M{"_id": cardID}, card); err != nil {
		return errors.Failed_to_deposite
	}
	return errors.Success
}

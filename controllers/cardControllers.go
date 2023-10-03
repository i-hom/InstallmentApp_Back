package controllers

import (
	"encoding/json"
	"installment_back/errors"
	"installment_back/models"
	"installment_back/services"
)

type CardController struct {
	cardService *services.CardService
}

func NewCardController(cardService *services.CardService) *CardController {
	return &CardController{cardService: cardService}
}

func (cc *CardController) Add(params interface{}) models.RPCResponse {
	var cardData models.AddCard
	if err := json.Unmarshal(models.GetRaw(params), &cardData); err != nil {
		return errors.Invalid_parameter
	}

	if cardData.ExpDate == "" || cardData.Number == "" {
		return errors.Missing_parameter
	}

	if len(cardData.Number) != 16 {
		return errors.Card_number_length_not_correct
	}

	if err := cc.cardService.Add(cardData); err != errors.Success {
		return errors.Failed_to_add_card
	}

	return models.RPCResponse{Code: 0, Message: "Card added"}
}

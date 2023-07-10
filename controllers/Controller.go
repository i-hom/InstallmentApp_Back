package controllers

import (
	"encoding/json"
	"installment_back/models"
	"io"
	"log"
	"net/http"
)

type Controller struct {
	UserModel        *User
	CardModel        *Card
	InstallmentModel *Installment
	DataBaseModel    DataBase
}

func (ctrl *Controller) Handler(w http.ResponseWriter, r *http.Request) {
	data, _ := io.ReadAll(r.Body)
	var request models.RPCRequest
	var response models.RPCResponse
	json.Unmarshal(data, &request)

	switch request.Method {
	case "card.add":
		{
			response = ctrl.CardModel.Add(request.Params, &ctrl.DataBaseModel)
		}
		break
	case "installment.pay":
		{
			response = ctrl.InstallmentModel.Pay(request.Params, &ctrl.DataBaseModel)
		}
		break
	case "user.get":
		response = ctrl.UserModel.Get(request.Params, &ctrl.DataBaseModel)

		break
	default:
		response = models.Method_not_found

	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("Failed to encode JSON response: %v", err)
	}
}

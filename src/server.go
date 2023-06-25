package src

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
)

func WebServer() {
	var ctrl Controller
	ctrl.DataBaseModel.Connect("Installment_Front")
	url := "0.0.0.0:7777"
	fmt.Printf("Server started! %s/EndPoint", url)
	http.HandleFunc("/EndPoint", ctrl.Handler)
	if err := http.ListenAndServe(url, nil); err != nil {
		log.Fatal(err)
	}
	defer ctrl.DataBaseModel.Disconnect()
}

func (ctrl *Controller) Handler(w http.ResponseWriter, r *http.Request) {
	data, _ := io.ReadAll(r.Body)
	var request RPCRequest
	var response RPCResponse
	json.Unmarshal(data, &request)

	switch request.Method {
	case "card.add":
		{
			response = ctrl.CardModel.Add(request.Params, ctrl.DataBaseModel)
		}
		break
	case "installment.pay":
		{
			response = ctrl.InstallmentModel.Pay(request.Params, ctrl.DataBaseModel)
		}
		break
	case "user.get":
		response = ctrl.UserModel.Get(request.Params, ctrl.DataBaseModel)

		break
	default:
		response = RPCResponse{Code: 1, Message: "Method not found"}

	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("Failed to encode JSON response: %v", err)
	}
}

func GetRaw(params interface{}) []byte {
	data, _ := json.Marshal(params)
	return data
}

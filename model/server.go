package server

import (
	"context"
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"io"
	"log"
	"model/model"
	"net/http"
)

func WebServer() {
	fmt.Println("Server started! http://localhost:777/EndPoint")
	http.HandleFunc("/EndPoint", Handler)
	err := http.ListenAndServe(":777", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func DataBase() *mongo.Client {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI("mongodb://localhost:27017").SetServerAPIOptions(serverAPI)
	conn, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		panic(err)
	}
	return conn
}

func Handler(w http.ResponseWriter, r *http.Request) {
	data, _ := io.ReadAll(r.Body)
	var request model.RPCRequest
	var response model.RPCResponse
	conn := DataBase()
	db := conn.Database("Elmakon")
	json.Unmarshal(data, &request)

	switch request.Method {
	case "card.add":
		{
			response = model.CardAdd(request.Params, db)
		}
		break
	case "installment.pay":
		{
			response = model.InstallmentPayment(request.Params, db)
		}
		break
	case "user.get":
		{
			response = model.UserAuth(request.Params, db)
		}
		break
	}
	w.Header().Set("Content-Type", "application/json")
	if err := json.NewEncoder(w).Encode(response); err != nil {
		log.Printf("Failed to encode JSON response: %v", err)
	}
	defer conn.Disconnect(context.TODO())
}

func GetRaw(params interface{}) []byte {
	data, _ := json.Marshal(params)
	return data
}

package model

import (
	"context"
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"io"
	"log"
	"net/http"
)

type RPCRequest struct {
	Method string      `json:"method"`
	Params interface{} `json:"params"`
}

type RPCResponse struct {
	Result interface{} `json:"result,omitempty"`
	Error  *RPCError   `json:"error,omitempty"`
}

type RPCError struct {
	Code    int    `json:"error_code"`
	Message string `json:"error_message"`
}

func WebServer() {
	fmt.Println("Server started! https://localhost:7777/EndPoint")
	http.HandleFunc("/EndPoint", Handler)
	//err := http.ListenAndServe("192.168.0.77:7777", nil)
	err := http.ListenAndServe("192.168.233.88:7777", nil)
	//err := http.ListenAndServe("localhost:7777", nil)
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
	var request RPCRequest
	var response RPCResponse
	conn := DataBase()
	db := conn.Database("Elmakon")
	json.Unmarshal(data, &request)
	switch request.Method {
	case "card.add":
		{
			response = CardAdd(request.Params, db)
		}
		break
	case "installment.pay":
		{
			response = InstallmentPayment(request.Params, db)
		}
		break
	case "user.get":
		{
			response = UserAuth(request.Params, db)
		}
		break
	default:
		{
			response = RPCResponse{Error: &RPCError{Code: 1, Message: "Method not found"}}
		}
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

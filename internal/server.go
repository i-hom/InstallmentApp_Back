package internal

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
	Code    int         `json:"code,omitempty"`
	Message string      `json:"msg,omitempty"`
	Data    interface{} `json:"data,omitempty"`
}

func WebServer() {
<<<<<<< refs/remotes/origin/master:model/server.go
	fmt.Println("Server started! https://localhost:7777/EndPoint")
	http.HandleFunc("/EndPoint", Handler)
	//err := http.ListenAndServe("192.168.0.77:7777", nil)
	err := http.ListenAndServe("192.168.233.88:7777", nil)
	//err := http.ListenAndServe("localhost:7777", nil)
	if err != nil {
=======
	urls := []string{"192.168.0.77:7777", "192.168.0.162:7777", "192.168.233.88:7777", "localhost:7777"}
	url := urls[0]
	fmt.Printf("Server started! %s/EndPoint", url)
	http.HandleFunc("/EndPoint", Handler)
	if err := http.ListenAndServe(url, nil); err != nil {
>>>>>>> Update logic and cleanup code:internal/server.go
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
	db := conn.Database("Installment_App")
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
<<<<<<< refs/remotes/origin/master:model/server.go
			response = RPCResponse{Error: &RPCError{Code: 1, Message: "Method not found"}}
=======
			response = RPCResponse{Code: 1, Message: "Method not found"}
>>>>>>> Update logic and cleanup code:internal/server.go
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

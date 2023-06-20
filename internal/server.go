package installment_back

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

type Controller struct {
	UserModel        *User
	CardModel        *Card
	InstallmentModel *Installment
	db               *mongo.Database
}

func WebServer() {
	var ctrl Controller

	var conn *mongo.Client
	ctrl.db, conn = DataBase("Installment_Front")
	defer conn.Disconnect(context.TODO())

	url := "0.0.0.0:7777"
	fmt.Printf("Server started! %s/EndPoint", url)
	http.HandleFunc("/EndPoint", ctrl.Handler)
	if err := http.ListenAndServe(url, nil); err != nil {
		log.Fatal(err)
	}
}

func DataBase(db_name string) (*mongo.Database, *mongo.Client) {
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI("mongodb://localhost:27017").SetServerAPIOptions(serverAPI)
	conn, err := mongo.Connect(context.TODO(), opts)
	if err != nil {
		log.Panic(err)
	}
	return conn.Database(db_name), conn
}

func (ctrl *Controller) Handler(w http.ResponseWriter, r *http.Request) {
	data, _ := io.ReadAll(r.Body)
	var request RPCRequest
	var response RPCResponse
	json.Unmarshal(data, &request)

	switch request.Method {
	case "card.add":
		{
			response = ctrl.CardModel.Add(request.Params, ctrl.db)
		}
		break
	case "installment.pay":
		{
			response = ctrl.InstallmentModel.Pay(request.Params, ctrl.db)
		}
		break
	case "user.get":
		response = ctrl.UserModel.Get(request.Params, ctrl.db)

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

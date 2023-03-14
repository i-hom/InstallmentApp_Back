package main

import (
	"context"
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"io"
	"log"
	"model/model"
	"net/http"
)

func main() {
	WebServer()
}

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
			response = CardAdd(request.Params, db)
		}
		break
	case "installment.pay":
		{
			response = InstallmentPay(request.Params, db)
		}
		break
	case "user.get":
		{
			response = UserGet(request.Params, db)
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

func CardAdd(params interface{}, db *mongo.Database) model.RPCResponse {
	var cardData model.JCard
	json.Unmarshal(GetRaw(params), &cardData)
	fmt.Println(cardData)
	if cardData.ExpDate == "" || cardData.Number == "" {
		err := model.RPCError{Code: 1, Message: "Missing one of params"}
		fmt.Println(err)
		return model.RPCResponse{Error: &err}
	}

	if len(cardData.Number) != 16 {
		err := model.RPCError{Code: 2, Message: "Card number should be 16 digits"}
		fmt.Println(err)
		return model.RPCResponse{Error: &err}
	}

	db.Collection("Cards").InsertOne(context.TODO(), cardData)

	return model.RPCResponse{Result: cardData}
}

func InstallmentPay(params interface{}, db *mongo.Database) model.RPCResponse {
	var installmentData model.InstallmentPay
	json.Unmarshal(GetRaw(params), &installmentData)
	if installmentData.InstallmentID.IsZero() || installmentData.CardID.IsZero() {
		err := model.RPCError{Code: 1, Message: "Missing one of params"}
		fmt.Println(err)
		return model.RPCResponse{Error: &err}
	}
	var cardData model.BCard
	db.Collection("Cards").FindOne(context.TODO(), bson.M{"_id": installmentData.CardID}).Decode(&cardData)

	var installment model.BInstallment
	db.Collection("Installments").FindOne(context.TODO(), bson.M{"_id": installmentData.InstallmentID}).Decode(&installment)

	fmt.Println(installmentData)
	fmt.Println(cardData)
	fmt.Println(installment)

	if cardData.Value < installmentData.Value {
		return model.RPCResponse{Error: &model.RPCError{Code: 3, Message: "Insufficient balance"}}
	}

	if installment < installmentData.Value {
		return model.RPCResponse{Error: &model.RPCError{Code: 6, Message: "U paid a lot"}}
	}

	cardData.Value -= installmentData.Value
	installment.Balance -= installmentData.Value

	_, cardDeposit := db.Collection("Cards").UpdateOne(context.TODO(), bson.M{"cardnumber": installmentData.CardNumber}, bson.M{"$set": cardData})
	_, installmentDeposit := db.Collection("Installments").UpdateOne(context.TODO(), bson.M{"elmakonid": installmentData.InstallmentID}, bson.M{"$set": installment})

	if cardDeposit != nil || installmentDeposit != nil {
		err := model.RPCError{Code: 5, Message: "Failed to deposit"}
		return model.RPCResponse{Error: &err}
	}

	if installment.Balance == 0 {
		installment.IsActive = false
		db.Collection("Installments").UpdateOne(context.TODO(), bson.M{"elmakonid": installment.ID}, bson.M{"$set": installment})
	}

	return model.RPCResponse{Result: "Successfully paid installment"}
}

func UserGet(params interface{}, db *mongo.Database) model.RPCResponse {
	var userGet model.UserGet
	json.Unmarshal(GetRaw(params), &userGet)
	if userGet.PhoneNumber == "" || userGet.Password == "" {
		err := model.RPCError{Code: 1, Message: "Missing one of params"}
		fmt.Println(err)
		return model.RPCResponse{Error: &err}
	}
	var user model.User
	if db.Collection("Users").FindOne(context.TODO(), bson.M{"phonenumber": userGet.PhoneNumber, "password": userGet.Password}).Decode(&user) != nil {
		return model.RPCResponse{Error: &model.RPCError{Code: 4, Message: "User not found"}}
	}
	fmt.Println(user)
	curr, _ := db.Collection("Installments").Find(context.TODO(), bson.M{"ownerid": user.ID})
	curr.All(context.TODO(), &user.UserInstallment)

	return model.RPCResponse{Result: user}
}

//InstallmentGet
//CardGet

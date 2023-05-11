package internal

import (
	"context"
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

//==================BSON=======================

type BUser struct {
	ID          primitive.ObjectID `bson:"_id"`
	FullName    string             `json:"fullName"`
	PhoneNumber string             `json:"phoneNumber"`
	PassID      string             `json:"passId"`
	CashBack    int                `json:"cashBack"`
	Password    string             `json:"password"`
}

func (user *BUser) ToJUser() JUser {
	return JUser{
		ID:          user.ID,
		FullName:    user.FullName,
		PhoneNumber: user.PhoneNumber,
		PassID:      user.PassID,
		CashBack:    user.CashBack,
	}
}

//===================JSON=======================

type JUser struct {
	ID          primitive.ObjectID `bson:"_id"`
	FullName    string             `json:"fullName"`
	PhoneNumber string             `json:"phoneNumber"`
	PassID      string             `json:"passId"`
	CashBack    int                `json:"cashBack"`
	Installment []JInstallment     `json:"installments"`
	Card        []BCard            `json:"cards"`
}

type UserLog struct {
	PhoneNumber string `json:"phoneNumber"`
	Password    string `json:"password"`
}

func UserAuth(params interface{}, db *mongo.Database) RPCResponse {
	var userAuth UserLog
	json.Unmarshal(GetRaw(params), &userAuth)
	if userAuth.PhoneNumber == "" || userAuth.Password == "" {
		return RPCResponse{Code: 1, Message: "Missing one of params"}
	}

	var buser BUser
	if db.Collection("Users").FindOne(context.TODO(), bson.M{"phonenumber": userAuth.PhoneNumber, "password": userAuth.Password}).Decode(&buser) != nil {
		return RPCResponse{Code: 4, Message: "User not found"}
	}

	if buser.ID.IsZero() {
		return RPCResponse{Data: BUser{}}
	}

	var juser JUser
	juser = buser.ToJUser()

	var installments []BInstallment
	var jinstallment []JInstallment
	curr, _ := db.Collection("Installments").Find(context.TODO(), bson.M{"ownerid": buser.ID})
	curr.All(context.TODO(), &installments)
	for _, i := range installments {
		var installment = i.ToJInstallment()
		installment.Item, _ = GetItem(i.ItemID, db)
		installment.Payments = GetPayments(i.ElmakonID, db)
		jinstallment = append(jinstallment, installment)
	}
	juser.Installment = jinstallment
	juser.Card = GetCards(buser.ID, db)
	return RPCResponse{Data: juser}
}

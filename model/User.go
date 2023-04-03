package model

import (
	"context"
	"encoding/json"
	"fmt"
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
		Password:    user.Password,
	}
}

//===================JSON=======================

type JUser struct {
	ID          primitive.ObjectID `json:"id"`
	FullName    string             `json:"fullName"`
	PhoneNumber string             `json:"phoneNumber"`
	PassID      string             `json:"passId"`
	CashBack    int                `json:"cashBack"`
	Password    string             `json:"password"`
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
		err := RPCError{Code: 1, Message: "Missing one of params"}
		fmt.Println(err)
		return RPCResponse{Error: &err}
	}

	var buser BUser
	if db.Collection("Users").FindOne(context.TODO(), bson.M{"phonenumber": userAuth.PhoneNumber, "password": userAuth.Password}).Decode(&buser) != nil {
		return RPCResponse{Error: &RPCError{Code: 4, Message: "User not found"}}
	}

	var juser JUser
	juser = buser.ToJUser()

	var installments []BInstallment
	var jinstallment []JInstallment
	curr, _ := db.Collection("Installments").Find(context.TODO(), bson.M{"ownerid": buser.ID})
	curr.All(context.TODO(), &installments)
	for _, i := range installments {
		var item BItem
		var jitem JItem
		var category Category
		db.Collection("Items").FindOne(context.TODO(), bson.M{"_id": i.ItemID}).Decode(&item)
		db.Collection("Category").FindOne(context.TODO(), bson.M{"_id": item.Category}).Decode(&category)
		jitem = item.ToJItem()
		jitem.Category = category.Name
		var installment = i.ToJInstallment()
		installment.Item = jitem
		jinstallment = append(jinstallment, installment)
	}
	juser.Installment = jinstallment
	curr, _ = db.Collection("Cards").Find(context.TODO(), bson.M{"ownerid": buser.ID})
	curr.All(context.TODO(), &juser.Card)

	return RPCResponse{Result: juser}
}

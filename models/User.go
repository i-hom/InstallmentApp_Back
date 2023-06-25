package models

import (
	"encoding/json"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"installment_back/src"
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

func (user *BUser) ToJUser() User {
	return User{
		ID:          user.ID,
		FullName:    user.FullName,
		PhoneNumber: user.PhoneNumber,
		PassID:      user.PassID,
		CashBack:    user.CashBack,
	}
}

//===================JSON=======================

type User struct {
	ID          primitive.ObjectID `bson:"_id"`
	FullName    string             `json:"fullName"`
	PhoneNumber string             `json:"phoneNumber"`
	PassID      string             `json:"passId"`
	CashBack    int                `json:"cashBack"`
	Installment []Installment      `json:"installments"`
	Card        []Card             `json:"cards"`
}

type UserLog struct {
	PhoneNumber string `json:"phoneNumber"`
	Password    string `json:"password"`
}

func (user *User) Get(params interface{}, db *src.DataBase) src.RPCResponse {
	var userAuth UserLog
	var buser BUser
	var juser User

	json.Unmarshal(src.GetRaw(params), &userAuth)
	if userAuth.PhoneNumber == "" || userAuth.Password == "" {
		return src.Missing_parameter
	}

	db.FindOne("Users", bson.M{"phoneNumber": userAuth.PhoneNumber, "password": userAuth.Password}, &buser)
	if buser.ID.IsZero() {
		return src.RPCResponse{Data: BUser{}}
	}

	juser = buser.ToJUser()
	juser.Installment, _ = InstallmentsGet(buser.ID, db)
	juser.Card = CardsGet(buser.ID, db)
	return src.RPCResponse{Data: juser}
}

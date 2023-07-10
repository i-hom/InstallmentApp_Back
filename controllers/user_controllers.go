package controllers

import (
	"encoding/json"
	"fmt"
	"go.mongodb.org/mongo-driver/bson"
	"installment_back/models"
)

type User models.User

func (user *User) Get(params interface{}, db *DataBase) models.RPCResponse {
	var installment Installment
	var card Card
	var userAuth models.UserLog
	var buser models.BUser
	var juser models.User

	json.Unmarshal(GetRaw(params), &userAuth)
	if userAuth.PhoneNumber == "" || userAuth.Password == "" {
		return models.Missing_parameter
	}

	db.FindOne("Users", bson.M{"phonenumber": userAuth.PhoneNumber, "password": userAuth.Password}).Decode(&buser)
	fmt.Println(buser)
	if buser.ID.IsZero() {
		return models.RPCResponse{Data: models.BUser{}}
	}

	juser = buser.ToJUser()
	juser.Installment, _ = installment.Get(buser.ID, db)
	juser.Card = card.Get(buser.ID, db)
	return models.RPCResponse{Data: juser}
}

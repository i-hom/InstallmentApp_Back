package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
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

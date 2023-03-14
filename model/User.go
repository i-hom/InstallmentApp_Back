package model

import "go.mongodb.org/mongo-driver/bson/primitive"

//==================BSON=======================

type BUser struct {
	ID          primitive.ObjectID `bson:"_id"`
	fullName    string             `bson:"fullName"`
	phoneNumber string             `bson:"phoneNumber"`
	passID      string             `bson:"passId"`
	cashBack    int                `bson:"cashBack"`
	password    string             `bson:"password"`
}

//===================JSON=======================

type JUser struct {
	ID          string `json:"id"`
	FullName    string `json:"fullName"`
	PhoneNumber string `json:"phoneNumber"`
	PassID      string `json:"passId"`
	CashBack    int    `json:"cashBack"`
	Password    string `json:"password"`
	Installment []JInstallment `json:"installments"`
	Card []JCard `json:"cards"`
}

type UserAuth struct {
	PhoneNumber string `json:"phone_number"`
	Password    string `json:"password"`
}


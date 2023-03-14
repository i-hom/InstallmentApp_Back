package model

import "go.mongodb.org/mongo-driver/bson/primitive"

//==================BSON=======================

type BInstallment struct {
	ID			primitive.ObjectID `bson:"_id"`
	elmakonID   string             `bson:"elmakonId"`
	itemID   	primitive.ObjectID `bson:"item"`
	balance  	int                `bson:"balance"`
	ownerID  	primitive.ObjectID `bson:"ownerId"`
	isActive 	bool               `bson:"isActive"`
}

//===================JSON=======================

type JInstallment struct {
	ID			string `json:"id"`
	ElmakonID   string `json:"elmakonId"`
	Item        JItem  `json:"item"`
	Balance  	int    `json:"balance"`
	IsActive 	bool   `json:"isActive"`
}

type InstallmentPay struct {
	InstallmentID primitive.ObjectID `json:"installment_id"`
	CardID   	  primitive.ObjectID `json:"number"`
	Value         int    			 `json:"value"`
}



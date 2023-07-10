package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

//==================BSON=======================

type BItem struct {
	ID       primitive.ObjectID `bson:"_id"`
	Brand    string             `json:"brand"`
	FullName string             `json:"fullName"`
	Image    string             `json:"image"`
	Category primitive.ObjectID `bson:"category"`
}

func (item *BItem) ToJItem() Item {
	return Item{
		Brand:    item.Brand,
		FullName: item.FullName,
		Image:    item.Image,
	}
}

type Category struct {
	ID   primitive.ObjectID `bson:"_id"`
	Name string             `json:"name"`
}

//==================JSON=======================

type Item struct {
	Brand    string `json:"brand"`
	FullName string `json:"fullName"`
	Image    string `json:"image"`
	Category string `json:"category"`
}

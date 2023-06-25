package models

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"installment_back/src"
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

func ItemGet(id primitive.ObjectID, db *src.DataBase) (Item, error) {
	var item BItem
	var category Category
	var jItem Item
	err := db.FindOne("Items", bson.M{"_id": id}, &item)
	if err != nil {
		return Item{}, err
	}
	err = db.FindOne("Category", bson.M{"_id": item.Category}, &category)
	if err != nil {
		return Item{}, err
	}
	jItem = item.ToJItem()
	jItem.Category = category.Name
	return jItem, nil
}

package controllers

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"installment_back/models"
)

type Item models.Item

func (this *Item) Get(id primitive.ObjectID, db *DataBase) (models.Item, error) {
	var item models.BItem
	var category models.Category
	var jItem models.Item
	db.FindOne("Items", bson.M{"_id": id}).Decode(&item)
	db.FindOne("Category", bson.M{"_id": item.Category}).Decode(&category)
	jItem = item.ToJItem()
	jItem.Category = category.Name
	return jItem, nil
}

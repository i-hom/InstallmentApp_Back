package internal

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

//==================BSON=======================

type BItem struct {
	ID       primitive.ObjectID `bson:"_id"`
	Brand    string             `json:"brand"`
	FullName string             `json:"fullName"`
	Image    string             `json:"image"`
	Category primitive.ObjectID `bson:"category"`
}

func (item *BItem) ToJItem() JItem {
	return JItem{
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

type JItem struct {
	Brand    string `json:"brand"`
	FullName string `json:"fullName"`
	Image    string `json:"image"`
	Category string `json:"category"`
}

func GetItem(id primitive.ObjectID, db *mongo.Database) (JItem, error) {
	var item BItem
	var category Category
	var jItem JItem
	if err := db.Collection("Items").FindOne(context.TODO(), bson.M{"_id": id}).Decode(&item); err != nil {
		return JItem{}, err
	}
<<<<<<< refs/remotes/origin/master:model/Item.go
	if err := db.Collection("Categories").FindOne(context.TODO(), bson.M{"_id": item.Category}).Decode(&category); err != nil {
=======
	if err := db.Collection("Category").FindOne(context.TODO(), bson.M{"_id": item.Category}).Decode(&category); err != nil {
>>>>>>> Update logic and cleanup code:internal/Item.go
		return JItem{}, err
	}
	jItem = item.ToJItem()
	jItem.Category = category.Name
<<<<<<< refs/remotes/origin/master:model/Item.go
=======

>>>>>>> Update logic and cleanup code:internal/Item.go
	return jItem, nil
}

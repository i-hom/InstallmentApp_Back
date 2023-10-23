package repositories

import (
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"installment_back/models"
	"installment_back/storage"
)

type ItemRepository struct {
	db *storage.DataBase
}

func NewItemRepository(db *storage.DataBase) *ItemRepository {
	return &ItemRepository{db: db}
}

func (ir *ItemRepository) Get(id primitive.ObjectID) (models.Item, error) {
	var item models.BItem
	var category models.Category
	var jItem models.Item
	ir.db.FindOne("Items", bson.M{"_id": id}, &item)
	ir.db.FindOne("Category", bson.M{"_id": item.Category}, &category)
	jItem = item.ToJItem()
	jItem.Category = category.Name
	return jItem, nil
}

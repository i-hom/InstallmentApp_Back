package src

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"log"
)

type DataBase struct {
	conn *mongo.Client
	db   *mongo.Database
}

func (this *DataBase) Connect(db_name string) error {
	var err error
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI("mongodb://localhost:27017").SetServerAPIOptions(serverAPI)
	this.conn, err = mongo.Connect(context.TODO(), opts)
	if err != nil {
		log.Panic(err)
	}

	this.db = this.conn.Database(db_name)
	return nil
}

func (this *DataBase) Disconnect() {
	this.conn.Disconnect(context.TODO())
}

func (this *DataBase) FindOne(collection_name string, query bson.M, result interface{}) error {
	err := this.db.Collection(collection_name).FindOne(context.TODO(), query).Decode(&result)
	if err != nil {
		return err
	}
	return nil
}

func (this *DataBase) FindAll(collection_name string, query bson.M, result interface{}) error {
	curr, _ := this.db.Collection(collection_name).Find(context.TODO(), query)
	err := curr.All(context.TODO(), &result)
	if err != nil {
		return err
	}
	return nil
}

func (this *DataBase) Insert(collection_name string, data interface{}) error {
	_, err := this.db.Collection(collection_name).InsertOne(context.TODO(), data)
	if err != nil {
		return err
	}
	return nil
}

func (this *DataBase) Update(collection_name string, query bson.M, data interface{}) (*mongo.UpdateResult, error) {
	res, err := this.db.Collection(collection_name).UpdateOne(context.TODO(), query, bson.M{"$set": data})
	if err != nil {
		return res, err
	}
	return res, nil
}

func (this *DataBase) Delete(colletion_name string, query bson.M) error {
	_, err := this.db.Collection(colletion_name).DeleteOne(context.TODO(), query)
	if err != nil {
		return err
	}
	return nil
}

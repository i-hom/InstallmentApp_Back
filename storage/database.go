package storage

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

func (db *DataBase) Connect(dbUrl, dbName string) error {
	var err error
	serverAPI := options.ServerAPI(options.ServerAPIVersion1)
	opts := options.Client().ApplyURI(dbUrl).SetServerAPIOptions(serverAPI)
	db.conn, err = mongo.Connect(context.TODO(), opts)
	if err != nil {
		log.Panic(err)
	}

	db.db = db.conn.Database(dbName)
	return nil
}

func (db *DataBase) Disconnect() {
	err := db.conn.Disconnect(context.TODO())
	if err != nil {
		log.Panic(err)
	}
}

func (db *DataBase) FindOne(collectionName string, query bson.M, data interface{}) error {
	return db.db.Collection(collectionName).FindOne(context.TODO(), query).Decode(data)
}

func (db *DataBase) FindAll(collectionName string, query bson.M, data interface{}) error {
	curr, _ := db.db.Collection(collectionName).Find(context.TODO(), query)
	return curr.All(context.TODO(), data)
}

func (db *DataBase) Insert(collectionName string, data interface{}) error {
	_, err := db.db.Collection(collectionName).InsertOne(context.TODO(), data)
	if err != nil {
		return err
	}
	return nil
}

func (db *DataBase) Update(collectionName string, query bson.M, data interface{}) (int64, error) {
	res, err := db.db.Collection(collectionName).UpdateOne(context.TODO(), query, bson.M{"$set": data})
	if err != nil {
		return 0, err
	}
	return res.ModifiedCount, nil
}

func (db *DataBase) Delete(collectionName string, query bson.M) error {
	_, err := db.db.Collection(collectionName).DeleteOne(context.TODO(), query)
	if err != nil {
		return err
	}
	return nil
}

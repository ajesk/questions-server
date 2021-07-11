package main

import (
	"context"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var client *mongo.Client

func GetClient() *mongo.Client {
	if client != nil {
		return client
	}

	var err error
	log.Print("connecting to ", os.Getenv("MONGODB_URI"))
	clientOptions := options.Client().ApplyURI(os.Getenv("MONGODB_URI"))
	client, err = mongo.Connect(context.Background(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	return client
}

func GetDatabase() *mongo.Database {
	return GetClient().Database("questions")
}

func GetCollection(collectionName string) *mongo.Collection {
	db := GetDatabase()
	return db.Collection(collectionName)
}

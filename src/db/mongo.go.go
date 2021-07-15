package db

import (
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

func GetMongoClientByCollection(collectionName string) (*mongo.Collection, error) {
	client, err := ConnectDB(collectionName)
	if err != nil {
		log.Fatal(err)
	}
	return client, err
}

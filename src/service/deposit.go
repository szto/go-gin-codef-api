package service

import (
	"context"
	"go-gin-codef-api/src/db"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"log"
)

const TEST_STORE_COLLECTION = "testStore"

func FindDepositByFilter(filter bson.M) (*mongo.Cursor, error) {

	client, err := db.GetMongoClientByCollection(TEST_STORE_COLLECTION)
	if err != nil {
		log.Fatal(err)
		return _, err
	}

	cursor, err := client.Find(context.Background(), filter)
	if err != nil {
		log.Fatal(err)
		return _, err
	}

	return cursor, _
}

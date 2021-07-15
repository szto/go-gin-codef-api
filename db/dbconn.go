package db

import (
	"context"
	"go-gin-codef-api/config"
	"log"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"go.mongodb.org/mongo-driver/mongo/readpref"
)

func ConnectDB(mongoDBCollection string) (*mongo.Collection, error) {
	config := config.InitConfig()

	MongoDBUrl := "mongodb://" + config.MongoDBHost + ":" + config.MongoDBPort

	credential := options.Credential{
		Username: config.MongoDBUserName,
		Password: config.MongoDBPassword,
	}

	clientOptions := options.Client().ApplyURI(MongoDBUrl).SetAuth(credential)

	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return nil, err
	}

	err = client.Ping(context.TODO(), readpref.Primary())
	if err != nil {
		log.Fatal(err)
	}

	mongoDBClient := client.Database(config.MongoDBName).Collection(mongoDBCollection)

	return mongoDBClient, err
}

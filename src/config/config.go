package config

import (
	"os"
)

type Config struct {
	MongoDBHost       string
	MongoDBPort       string
	MongoDBName       string
	MongoDBUserName   string
	MongoDBPassword   string
	CodefPublicKey    string
	CodefClientId     string
	CodefClientSecret string
}

func InitConfig() Config {
	return Config{
		MongoDBHost:       os.Getenv("MONGO_DB_HOST"),
		MongoDBPort:       os.Getenv("MONGO_DB_PORT"),
		MongoDBName:       os.Getenv("MONGO_DB_NAME"),
		MongoDBUserName:   os.Getenv("MONGO_DB_USER_NAME"),
		MongoDBPassword:   os.Getenv("MONGO_DB_PASSWORD"),
		CodefPublicKey:    os.Getenv("CODEF_PUBLIC_KEY"),
		CodefClientId:     os.Getenv("CODEF_CLIENT_ID"),
		CodefClientSecret: os.Getenv("CODEF_CLIENT_SECRET"),
	}
}

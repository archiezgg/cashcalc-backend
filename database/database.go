package database

import (
	"context"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	mongoUser     = os.Getenv("MONGO_USER")
	mongoPassword = os.Getenv("MONGO_PW")
	mongoHostURL  = os.Getenv("MONGO_HOST")
	mongoDBName   = os.Getenv("MONG_DB_NAME")

	mongoClient *mongo.Client
)

// Startup is the init call of the mongo DB, supposed to be called in the main function
func Startup() *mongo.Client {
	dbSpec := fmt.Sprintf("mongodb+srv://%v:%v@%v", mongoUser, mongoPassword, mongoHostURL)

	var err error
	mongoClient, err = mongo.Connect(context.TODO(), options.Client().ApplyURI(dbSpec))
	if err != nil {
		log.Fatal("Couldn't connect to database: ", err)
	}

	err = mongoClient.Ping(context.TODO(), nil)
	if err != nil {
		log.Fatal("Database is not responding!")
	}
	log.Println(("Successfully connected to the database!"))
	return mongoClient
}

// GetCollectionByName returns a collection type from the db by its name
func GetCollectionByName(collectionName string) *mongo.Collection {
	coll := mongoClient.Database(mongoDBName).Collection(collectionName)
	return coll
}

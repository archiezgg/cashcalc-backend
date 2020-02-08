package database

import (
	"fmt"
	"log"
	"os"

	"gopkg.in/mgo.v2"
)

var (
	mongoUser     = os.Getenv("MONGO_USER")
	mongoPassword = os.Getenv("MONGO_PW")
	mongoHost     = os.Getenv("MONGO_HOST")
	mongoPort     = os.Getenv("MONGO_PORT")
	mongoDBName   = os.Getenv("MONGO_DB_NAME")

	dbSession *mgo.Session
)

// Startup is the init call of the mongo DB, supposed to be called in the main function
func Startup() *mgo.Session {
	dbSpec := fmt.Sprintf("mongodb+srv://%v:%v@%v:%v/%v", mongoUser, mongoPassword, mongoHost, mongoPort, mongoDBName)

	dbSession, err := mgo.Dial(dbSpec)
	if err != nil {
		log.Fatal("couldn't connect to database: ", err)
	}

	log.Println(("successfully connected to the database!"))
	return dbSession
}

// GetCollectionByName returns a collection type from the db by its name
func GetCollectionByName(collectionName string) *mgo.Collection {
	coll := dbSession.DB(mongoDBName).C(collectionName)
	return coll
}

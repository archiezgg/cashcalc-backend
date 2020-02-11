package database

import (
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
	mongoURI      = os.Getenv("MONGO_URI")

	dbSession *mgo.Session
)

// Startup is the init call of the mongo DB, supposed to be called in the main function
func Startup() *mgo.Session {
	var err error
	dbSession, err = mgo.Dial(mongoURI)
	if err != nil {
		log.Fatal("couldn't connect to database: ", err)
	}

	log.Println(("successfully connected to the database!"))
	return dbSession
}

// GetCollectionByName returns a collection type from the db by its name
func GetCollectionByName(collectionName string) *mgo.Collection {
	coll := dbSession.Clone().DB(mongoDBName).C(collectionName)
	return coll
}

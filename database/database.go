package database

import (
	"log"
	"os"

	"github.com/IstvanN/cashcalc-backend/services"
	"gopkg.in/mgo.v2"
)

var (
	mongoURI = os.Getenv("MONGODB_URI")

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
	dbName := services.GetDBNameFromURI(mongoURI)
	coll := dbSession.Clone().DB(dbName).C(collectionName)
	return coll
}

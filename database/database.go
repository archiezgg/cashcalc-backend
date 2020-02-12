package database

import (
	"log"
	"os"
	"strings"

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
	dbName := getDBNameFromURI(mongoURI)
	coll := dbSession.Clone().DB(dbName).C(collectionName)
	return coll
}

//TODO: write a function to strip the DB name from the URI and store it in env variable
func getDBNameFromURI(uri string) string {
	splitURI := strings.SplitAfter(uri, "/")
	return splitURI[len(splitURI)-1]
}

/*
	Cashcalc 2020
	Copyright (C) 2019-2020 Istvan Nemeth
	mailto: nemethistvanius@gmail.com
*/

package database

import (
	"log"
	"os"

	"github.com/IstvanN/cashcalc-backend/services"
	"github.com/globalsign/mgo"
)

var (
	mongoURI = os.Getenv("MONGODB_URI")

	dbSession *mgo.Session
)

// StartupMongo is the init call of the mongo DB, supposed to be called in the main function
func StartupMongo() *mgo.Session {
	var err error
	dbSession, err = mgo.Dial(mongoURI)
	if err != nil {
		log.Fatal("couldn't connect to database: ", err)
	}

	log.Println(("successfully connected to mongoDB!"))
	return dbSession
}

// GetCollectionByName returns a collection type from the db by its name
func GetCollectionByName(collectionName string) *mgo.Collection {
	dbName := services.GetMongoDBNameFromURI(mongoURI)
	coll := dbSession.DB(dbName).C(collectionName)
	return coll
}

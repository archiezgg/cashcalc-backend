package model

import (
	"context"
	"log"
	"os"

	"github.com/IstvanN/cashcalc-backend/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
)

var (
	airCountriesCollectionName  = os.Getenv("COUNTRIES_AIR_COLL")
	roadCountriesCollectionName = os.Getenv("COUNTRIES_ROAD_COLL")
)

// Country stores the countries with name and a zone number
type Country struct {
	Name       string
	ZoneNumber int
}

// GetAirCountriesFromDB returns with a slice of all elements of the airCountries collection
func GetAirCountriesFromDB() []Country {
	coll := database.GetCollectionByName(airCountriesCollectionName)
	cur, err := coll.Find(context.TODO(), bson.D{{}}, options.Find())
	if err != nil {
		log.Printf("retrieving collection %v failed: %v\n", airCountriesCollectionName, err)
	}

	var airCountries []Country
	for cur.Next(context.TODO()) {
		var c Country
		err := cur.Decode(&c)
		if err != nil {
			log.Println("error while decoding air country: ", err)
		} else {
			airCountries = append(airCountries, c)
		}
	}
	cur.Close(context.TODO())
	return airCountries
}

// GetRoadCountriesFromDB returns with an array of all the elements of the roadCountries collection
func GetRoadCountriesFromDB() []Country {
	coll := database.GetCollectionByName(roadCountriesCollectionName)
	cur, err := coll.Find(context.TODO(), bson.D{{}}, options.Find())
	if err != nil {
		log.Printf("retrieving collection %v failed: %v\n", roadCountriesCollectionName, err)
	}

	var roadCountries []Country
	for cur.Next(context.TODO()) {
		var c Country
		err := cur.Decode(&c)
		if err != nil {
			log.Println("error while decoding road country: ", err)
		} else {
			roadCountries = append(roadCountries, c)
		}
	}
	cur.Close(context.TODO())
	return roadCountries
}

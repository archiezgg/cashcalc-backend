package model

import (
	"context"
	"fmt"
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

// GetAirCountriesFromDB returns with a slice of all elements of the airCountries collection, or an error
func GetAirCountriesFromDB() ([]Country, error) {
	coll := database.GetCollectionByName(airCountriesCollectionName)

	cur, err := coll.Find(context.TODO(), bson.D{{}}, options.Find())
	defer cur.Close(context.TODO())
	if err != nil {
		return nil, fmt.Errorf("retrieving collection %v failed: %v", airCountriesCollectionName, err)
	}

	var airCountries []Country
	for cur.Next(context.TODO()) {
		var c Country
		err := cur.Decode(&c)
		if err != nil {
			return nil, fmt.Errorf("error while decoding air country: %v", err)
		}
		airCountries = append(airCountries, c)
	}
	return airCountries, nil
}

// GetRoadCountriesFromDB returns with an array of all the elements of the roadCountries collection, or an error
func GetRoadCountriesFromDB() ([]Country, error) {
	coll := database.GetCollectionByName(roadCountriesCollectionName)

	cur, err := coll.Find(context.TODO(), bson.D{{}}, options.Find())
	defer cur.Close(context.TODO())
	if err != nil {
		return nil, fmt.Errorf("retrieving collection %v failed: %v", roadCountriesCollectionName, err)
	}

	var roadCountries []Country
	for cur.Next(context.TODO()) {
		var c Country
		err := cur.Decode(&c)
		if err != nil {
			return nil, fmt.Errorf("error while decoding road country: %v", err)
		}
		roadCountries = append(roadCountries, c)
	}
	cur.Close(context.TODO())
	return roadCountries, nil
}

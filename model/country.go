package model

import (
	"fmt"
	"os"

	"github.com/IstvanN/cashcalc-backend/database"
)

var (
	countriesCollectionName = os.Getenv("COUNTRIES_COLL")
)

// Country stores the countries with name and a zone number
type Country struct {
	Name       string `bson:"name"`
	ZoneNumber int    `bson:"zoneNumber"`
}

// Countries stores both air and road lists as fields
type Countries struct {
	CountriesAir  []Country `bson:"countriesAir"`
	CountriesRoad []Country `bson:"countriesRoad"`
}

// GetCountriesAirFromDB returns with a slice of all air elements of the Countries collection, or an error
func GetCountriesAirFromDB() ([]Country, error) {
	c, err := getCountriesFromDB()
	if err != nil {
		return nil, err
	}

	return c.CountriesAir, nil
}

// GetCountriesRoadFromDB returns with an array of all road elements of the Countries collection, or an error
func GetCountriesRoadFromDB() ([]Country, error) {
	c, err := getCountriesFromDB()
	if err != nil {
		return nil, err
	}

	return c.CountriesRoad, nil
}

func getCountriesFromDB() (Countries, error) {
	coll := database.GetCollectionByName(countriesCollectionName)

	var c Countries
	err := coll.Find(nil).One(&c)
	if err != nil {
		return Countries{}, fmt.Errorf("error while retrieving collection %v from database: %v", countriesCollectionName, err)
	}

	return c, nil
}

package model

import (
	"fmt"
	"os"

	"github.com/IstvanN/cashcalc-backend/database"
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

	var airCountries []Country
	err := coll.Find(nil).All(&airCountries)
	if err != nil {
		return nil, fmt.Errorf("error while retrieving collection %v from database: %v", airCountriesCollectionName, err)
	}
	return airCountries, nil
}

// GetRoadCountriesFromDB returns with an array of all the elements of the roadCountries collection, or an error
func GetRoadCountriesFromDB() ([]Country, error) {
	coll := database.GetCollectionByName(roadCountriesCollectionName)

	var roadCountries []Country
	err := coll.Find(nil).All(&roadCountries)
	if err != nil {
		return nil, fmt.Errorf("error while retrieving collection %v from database: %v", roadCountriesCollectionName, err)
	}
	return roadCountries, nil
}

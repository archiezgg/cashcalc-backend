package repositories

import (
	"fmt"

	"github.com/IstvanN/cashcalc-backend/database"
	"github.com/IstvanN/cashcalc-backend/models"
	"github.com/IstvanN/cashcalc-backend/properties"
)

var (
	countriesCollectionName = properties.Prop.GetString(properties.CountriesCollection, "countries")
)

// GetCountries queries all countries from db
func GetCountries() (models.Countries, error) {
	coll := database.GetCollectionByName(countriesCollectionName)

	var c models.Countries
	err := coll.Find(nil).One(&c)
	if err != nil {
		return models.Countries{}, fmt.Errorf("error while retrieving collection %v from database: %v",
			countriesCollectionName, err)
	}

	return c, nil
}

// GetAirCountries returns with a slice of all air elements of the Countries collection, or an error
func GetAirCountries() ([]models.Country, error) {
	c, err := GetCountries()
	if err != nil {
		return nil, err
	}

	return c.CountriesAir, nil
}

// GetRoadCountries returns with an array of all road elements of the Countries collection, or an error
func GetRoadCountries() ([]models.Country, error) {
	c, err := GetCountries()
	if err != nil {
		return nil, err
	}

	return c.CountriesRoad, nil
}

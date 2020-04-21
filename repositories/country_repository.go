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

// GetCountriesFromDB queries all countries from db
func GetCountriesFromDB() (models.Countries, error) {
	coll := database.GetCollectionByName(countriesCollectionName)

	var c models.Countries
	err := coll.Find(nil).One(&c)
	if err != nil {
		return models.Countries{}, fmt.Errorf("error while retrieving collection %v from database: %v", countriesCollectionName, err)
	}

	return c, nil
}

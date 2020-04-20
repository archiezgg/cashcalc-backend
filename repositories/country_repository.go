package repositories

import (
	"fmt"
	"os"

	"github.com/IstvanN/cashcalc-backend/database"
	"github.com/IstvanN/cashcalc-backend/models"
)

var (
	countriesCollectionName = os.Getenv("COUNTRIES_COLL")
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

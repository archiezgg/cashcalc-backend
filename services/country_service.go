package services

import (
	"github.com/IstvanN/cashcalc-backend/models"
	"github.com/IstvanN/cashcalc-backend/repositories"
)

// GetCountries returns all countries with Countries struct
func GetCountries() (models.Countries, error) {
	c, err := repositories.GetCountriesFromDB()
	if err != nil {
		return models.Countries{}, err
	}
	return c, nil
}

// GetAirCountries returns with a slice of all air elements of the Countries collection, or an error
func GetAirCountries() ([]models.Country, error) {
	c, err := repositories.GetCountriesFromDB()
	if err != nil {
		return nil, err
	}

	return c.CountriesAir, nil
}

// GetRoadCountries returns with an array of all road elements of the Countries collection, or an error
func GetRoadCountries() ([]models.Country, error) {
	c, err := repositories.GetCountriesFromDB()
	if err != nil {
		return nil, err
	}

	return c.CountriesRoad, nil
}

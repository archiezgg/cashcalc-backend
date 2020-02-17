package repositories

import (
	"fmt"
	"os"

	"github.com/IstvanN/cashcalc-backend/database"
	"github.com/IstvanN/cashcalc-backend/models"
	"github.com/IstvanN/cashcalc-backend/services"
)

var (
	pricingsCollectionName = os.Getenv("PRICINGS_COLL")
	pricingVarsCollName    = os.Getenv("PRICING_VARS_COLL")
)

// GetAirPricingsFromDB returns with a slice of all elements of the air pricings collection, or an error
func GetAirPricingsFromDB() ([]models.Pricing, error) {
	p, err := getPricingsFromDB()
	if err != nil {
		return nil, err
	}
	return p.AirPricings, nil
}

// GetRoadPricingsFromDB returns with a slice of all elements of the road pricings collection or an error
func GetRoadPricingsFromDB() ([]models.Pricing, error) {
	p, err := getPricingsFromDB()
	if err != nil {
		return nil, err
	}
	return p.RoadPricings, nil
}

// GetAirPricingFaresByZoneNumber takes a zone number int as parameter
// and returns with the corresponding air pricing fares as slice of ints, or an error
func GetAirPricingFaresByZoneNumber(zn int) ([]models.Fare, error) {
	if err := services.ValidateZoneNumber(zn); err != nil {
		return []models.Fare{}, err
	}

	ap, err := GetAirPricingsFromDB()
	if err != nil {
		return []models.Fare{}, err
	}

	for _, p := range ap {
		if p.ZoneNumber == zn {
			return p.Fares, nil
		}
	}

	return []models.Fare{}, fmt.Errorf("can't find number %v in air pricing fares", zn)
}

// GetAirPricingDocFaresByZoneNumber takes a zone number int as parameter
// and returns with the corresponding air pricing doc fares as slice of ints, or an error
func GetAirPricingDocFaresByZoneNumber(zn int) ([]models.Fare, error) {
	if zn < 5 || zn > 9 {
		return []models.Fare{}, fmt.Errorf("the zone number %v is invalid, it doesn't contain doc fares", zn)
	}

	ap, err := GetAirPricingsFromDB()
	if err != nil {
		return []models.Fare{}, err
	}

	for _, p := range ap {
		if p.ZoneNumber == zn {
			return p.DocFares, nil
		}
	}
	return []models.Fare{}, fmt.Errorf("can't find number %v in air pricing doc fares", zn)
}

// GetRoadPricingFaresByZoneNumber takes a zone number int as parameter
// and returns with the corresponding road pricing fares as slice of ints, or an error
func GetRoadPricingFaresByZoneNumber(zn int) ([]models.Fare, error) {
	if err := services.ValidateZoneNumber(zn); err != nil {
		return []models.Fare{}, err
	}

	rp, err := GetRoadPricingsFromDB()
	if err != nil {
		return []models.Fare{}, err
	}

	for _, p := range rp {
		if p.ZoneNumber == zn {
			return p.Fares, nil
		}
	}

	return []models.Fare{}, fmt.Errorf("can't find number %v in road pricing fares", zn)
}

// GetPricingVariablesFromDB retreives the pricing variables
// that can be set by the admin
func GetPricingVariablesFromDB() (models.PricingVariables, error) {
	coll := database.GetCollectionByName(pricingVarsCollName)

	var pv models.PricingVariables
	err := coll.Find(nil).One(&pv)
	if err != nil {
		return models.PricingVariables{}, fmt.Errorf("error while retreiving collection %v from db: %v", pricingVarsCollName, err)
	}

	return pv, nil
}

func getPricingsFromDB() (models.Pricings, error) {
	coll := database.GetCollectionByName(pricingsCollectionName)

	var p models.Pricings
	err := coll.Find(nil).One(&p)
	if err != nil {
		return models.Pricings{}, fmt.Errorf("error while retrieving collection %v from database: %v", pricingsCollectionName, err)
	}

	return p, nil
}

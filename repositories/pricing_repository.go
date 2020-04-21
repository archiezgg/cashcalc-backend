package repositories

import (
	"fmt"

	"github.com/IstvanN/cashcalc-backend/database"
	"github.com/IstvanN/cashcalc-backend/models"
	"github.com/IstvanN/cashcalc-backend/properties"
	"github.com/IstvanN/cashcalc-backend/services"
)

var (
	pricingsCollectionName = properties.Prop.GetString(properties.PricingsCollection, "pricings")
	pricingVarsCollName    = properties.Prop.GetString(properties.PricingVarsCollection, "pricingVariables")
)

// GetPricings queries the db for both road and air pricings
func GetPricings() (models.Pricings, error) {
	coll := database.GetCollectionByName(pricingsCollectionName)

	var p models.Pricings
	err := coll.Find(nil).One(&p)
	if err != nil {
		return models.Pricings{}, fmt.Errorf("error while retrieving collection %v from database: %v", pricingsCollectionName, err)
	}

	return p, nil
}

// GetPricingVariablesFromDB queries the db for
// the pricing variables that can be set by the admin
func GetPricingVariablesFromDB() (models.PricingVariables, error) {
	coll := database.GetCollectionByName(pricingVarsCollName)

	var pv models.PricingVariables
	err := coll.Find(nil).One(&pv)
	if err != nil {
		return models.PricingVariables{}, fmt.Errorf("error while retreiving collection %v from db: %v", pricingVarsCollName, err)
	}

	return pv, nil
}

// GetAirPricings returns with a slice of all elements of the air pricings collection, or an error
func GetAirPricings() ([]models.Pricing, error) {
	p, err := GetPricings()
	if err != nil {
		return nil, err
	}
	return p.AirPricings, nil
}

// GetRoadPricings returns with a slice of all elements of the road pricings collection or an error
func GetRoadPricings() ([]models.Pricing, error) {
	p, err := GetPricings()
	if err != nil {
		return nil, err
	}
	return p.RoadPricings, nil
}

// GetAirFaresByZoneNumber takes a zone number int as parameter
// and returns with the corresponding air pricing fares as slice of ints, or an error
func GetAirFaresByZoneNumber(zn int) ([]models.Fare, error) {
	if err := services.ValidateAirFaresZoneNumber(zn); err != nil {
		return []models.Fare{}, err
	}

	ap, err := GetAirPricings()
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

// GetAirFaresByZoneNumberAndWeight returns the weight-base fare pairing of the given zone number and weight
func GetAirFaresByZoneNumberAndWeight(zn int, weight float64) (models.Fare, error) {
	err := services.ValidateAirFaresZoneNumber(zn)
	err2 := services.ValidateAirFaresWeight(weight)
	if err != nil || err2 != nil {
		return models.Fare{}, err
	}

	ap, err := GetAirFaresByZoneNumber(zn)
	if err != nil {
		return models.Fare{}, err
	}

	for _, p := range ap {
		if p.Weight == weight {
			return p, nil
		}
	}

	return models.Fare{}, fmt.Errorf("can't find air fare with zone number: %v and weight: %v", zn, weight)
}

// GetAirDocFaresByZoneNumber takes a zone number int as parameter
// and returns with the corresponding air pricing doc fares as slice of ints, or an error
func GetAirDocFaresByZoneNumber(zn int) ([]models.Fare, error) {
	if err := services.ValidateAirDocFaresZoneNumber(zn); err != nil {
		return []models.Fare{}, err
	}

	ap, err := GetAirPricings()
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

// GetAirDocFaresByZoneNumberAndWeight returns the weight-base fare pairing of the given zone number and weight
func GetAirDocFaresByZoneNumberAndWeight(zn int, weight float64) (models.Fare, error) {
	err := services.ValidateAirDocFaresZoneNumber(zn)
	err2 := services.ValidateAirDocFaresWeight(weight)
	if err != nil || err2 != nil {
		return models.Fare{}, err
	}

	ap, err := GetAirDocFaresByZoneNumber(zn)
	if err != nil {
		return models.Fare{}, err
	}

	for _, p := range ap {
		if p.Weight == weight {
			return p, nil
		}
	}

	return models.Fare{}, fmt.Errorf("can't find air docfare with zone number: %v and weight: %v", zn, weight)
}

// GetRoadFaresByZoneNumber takes a zone number int as parameter
// and returns with the corresponding road pricing fares as slice of ints, or an error
func GetRoadFaresByZoneNumber(zn int) ([]models.Fare, error) {
	if err := services.ValidateRoadFaresZoneNumber(zn); err != nil {
		return []models.Fare{}, err
	}

	rp, err := GetRoadPricings()
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

// GetRoadFaresByZoneNumberAndWeight returns the weight-base fare pairing of the given zone number and weight
func GetRoadFaresByZoneNumberAndWeight(zn int, weight float64) (models.Fare, error) {
	err := services.ValidateRoadFaresZoneNumber(zn)
	err2 := services.ValidateRoadFaresWeight(weight)
	if err != nil || err2 != nil {
		return models.Fare{}, err
	}

	rp, err := GetRoadFaresByZoneNumber(zn)
	if err != nil {
		return models.Fare{}, err
	}

	for _, p := range rp {
		if p.Weight == weight {
			return p, nil
		}
	}

	return models.Fare{}, fmt.Errorf("can't find road fare with zone number: %v and weight: %v", zn, weight)
}

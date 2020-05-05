package repositories

import (
	"fmt"

	"github.com/IstvanN/cashcalc-backend/database"
	"github.com/IstvanN/cashcalc-backend/models"
	"github.com/IstvanN/cashcalc-backend/properties"
	"github.com/IstvanN/cashcalc-backend/services"
)

// GetPricings queries the db for both road and air pricings
func GetPricings() (models.Pricings, error) {
	coll := database.GetCollectionByName(properties.PricingsCollection)

	var p models.Pricings
	err := coll.Find(nil).One(&p)
	if err != nil {
		return models.Pricings{}, fmt.Errorf("error while retrieving collection %v from database: %v",
			properties.PricingsCollection, err)
	}

	return p, nil
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
	if services.IsZoneNumberInvalid(zn, properties.AirFaresZnMin, properties.AirFaresZnMax) {
		return nil, fmt.Errorf("the zone number %v is invalid for air fares", zn)
	}

	ap, err := GetAirPricings()
	if err != nil {
		return nil, err
	}

	for _, p := range ap {
		if p.ZoneNumber == zn {
			return p.Fares, nil
		}
	}

	return nil, fmt.Errorf("can't find number %v in air pricing fares", zn)
}

// GetAirFaresByZoneNumberAndWeight returns the weight-base fare pairing of the given zone number and weight
func GetAirFaresByZoneNumberAndWeight(zn int, weight float64) (models.Fare, error) {
	if services.IsZoneNumberInvalid(zn, properties.AirFaresZnMin, properties.AirFaresZnMax) ||
		services.IsWeightInvalid(weight, properties.AirFaresWeightMin, properties.AirFaresWeightMax) {
		return models.Fare{}, fmt.Errorf("the zone number %v and weight %v is invalid for air fares", zn, weight)
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
	if services.IsZoneNumberInvalid(zn, properties.AirDocFaresZnMin, properties.AirDocFaresZnMax) {
		return nil, fmt.Errorf("the zone number %v is invalid for air document fares", zn)
	}

	ap, err := GetAirPricings()
	if err != nil {
		return nil, err
	}

	for _, p := range ap {
		if p.ZoneNumber == zn {
			return p.DocFares, nil
		}
	}
	return nil, fmt.Errorf("can't find number %v in air pricing doc fares", zn)
}

// GetAirDocFaresByZoneNumberAndWeight returns the weight-base fare pairing of the given zone number and weight
func GetAirDocFaresByZoneNumberAndWeight(zn int, weight float64) (models.Fare, error) {
	if services.IsZoneNumberInvalid(zn, properties.AirDocFaresZnMin, properties.AirDocFaresZnMax) ||
		services.IsWeightInvalid(weight, properties.AirDocFaresWeightMin, properties.AirDocFaresWeightMax) {
		return models.Fare{}, fmt.Errorf("the zone number %v and weight %v is invalid for air document fares", zn, weight)
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
	if services.IsZoneNumberInvalid(zn, properties.RoadFaresZnMin, properties.RoadFaresZnMax) {
		return nil, fmt.Errorf("the zone number %v is invalid for road fares", zn)
	}

	rp, err := GetRoadPricings()
	if err != nil {
		return nil, err
	}

	for _, p := range rp {
		if p.ZoneNumber == zn {
			return p.Fares, nil
		}
	}

	return nil, fmt.Errorf("can't find number %v in road pricing fares", zn)
}

// GetRoadFaresByZoneNumberAndWeight returns the weight-base fare pairing of the given zone number and weight
func GetRoadFaresByZoneNumberAndWeight(zn int, weight float64) (models.Fare, error) {
	if services.IsZoneNumberInvalid(zn, properties.RoadFaresZnMin, properties.RoadFaresZnMax) ||
		services.IsWeightInvalid(weight, properties.RoadFaresWeightMin, properties.RoadFaresWeightMax) {
		return models.Fare{}, fmt.Errorf("the zone number %v and weight %v is invalid for road fares", zn, weight)
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

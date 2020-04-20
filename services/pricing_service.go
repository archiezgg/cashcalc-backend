package services

import (
	"fmt"

	"github.com/IstvanN/cashcalc-backend/models"
	"github.com/IstvanN/cashcalc-backend/repositories"
)

// IncreaseWithVat takes a float64 and a percentage as parameter
// and returns with the vat-increased result
func IncreaseWithVat(num float64, vat float64) float64 {
	return num * (1 + (vat / 100))
}

// ValidateZoneNumber takes an int as parameter and checks if it between 0 and 9
// returns with error if not
func ValidateZoneNumber(zn int) error {
	if zn < 0 || zn > 9 {
		return fmt.Errorf("the zone number %v is invalid", zn)
	}

	return nil
}

// GetPricings returns Pricings struct of both air and road pricings
func GetPricings() (models.Pricings, error) {
	p, err := repositories.GetPricingsFromDB()
	if err != nil {
		return models.Pricings{}, err
	}

	return p, nil
}

// GetAirPricings returns with a slice of all elements of the air pricings collection, or an error
func GetAirPricings() ([]models.Pricing, error) {
	p, err := repositories.GetPricingsFromDB()
	if err != nil {
		return nil, err
	}
	return p.AirPricings, nil
}

// GetRoadPricings returns with a slice of all elements of the road pricings collection or an error
func GetRoadPricings() ([]models.Pricing, error) {
	p, err := repositories.GetPricingsFromDB()
	if err != nil {
		return nil, err
	}
	return p.RoadPricings, nil
}

// GetAirPricingFaresByZoneNumber takes a zone number int as parameter
// and returns with the corresponding air pricing fares as slice of ints, or an error
func GetAirPricingFaresByZoneNumber(zn int) ([]models.Fare, error) {
	if err := ValidateZoneNumber(zn); err != nil {
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

// GetAirPricingFaresByZoneNumberAndWeight returns the weight-base fare pairing of the given zone number and weight
func GetAirPricingFaresByZoneNumberAndWeight(zn int, weight float64) (models.Fare, error) {
	if err := ValidateZoneNumber(zn); err != nil {
		return models.Fare{}, err
	}

	ap, err := GetAirPricingFaresByZoneNumber(zn)
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

// GetAirPricingDocFaresByZoneNumber takes a zone number int as parameter
// and returns with the corresponding air pricing doc fares as slice of ints, or an error
func GetAirPricingDocFaresByZoneNumber(zn int) ([]models.Fare, error) {
	if zn < 5 || zn > 9 {
		return []models.Fare{}, fmt.Errorf("the zone number %v is invalid, it doesn't contain doc fares", zn)
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

// GetAirPricingDocFaresByZoneNumberAndWeight returns the weight-base fare pairing of the given zone number and weight
func GetAirPricingDocFaresByZoneNumberAndWeight(zn int, weight float64) (models.Fare, error) {
	if zn < 5 || zn > 9 {
		return models.Fare{}, fmt.Errorf("the zone number %v is invalid, it doesn't contain doc fares", zn)
	}

	ap, err := GetAirPricingDocFaresByZoneNumber(zn)
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

// GetRoadPricingFaresByZoneNumber takes a zone number int as parameter
// and returns with the corresponding road pricing fares as slice of ints, or an error
func GetRoadPricingFaresByZoneNumber(zn int) ([]models.Fare, error) {
	if err := ValidateZoneNumber(zn); err != nil {
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

// GetRoadPricingFaresByZoneNumberAndWeight returns the weight-base fare pairing of the given zone number and weight
func GetRoadPricingFaresByZoneNumberAndWeight(zn int, weight float64) (models.Fare, error) {
	if err := ValidateZoneNumber(zn); err != nil {
		return models.Fare{}, err
	}

	rp, err := GetRoadPricingFaresByZoneNumber(zn)
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

package service

import (
	"fmt"

	"github.com/IstvanN/cashcalc-backend/model"
)

// GetAirPricingFaresByZoneNumber takes a zone number int as parameter
// and returns with the corresponding air pricing fares as slice of ints, or an error
func GetAirPricingFaresByZoneNumber(zn int) ([]int, error) {
	if err := validateZoneNumber(zn); err != nil {
		return nil, err
	}

	ap, err := model.GetAirPricingsFromDB()
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

// GetAirPricingDocFaresByZoneNumber takes a zone number int as parameter
// and returns with the corresponding air pricing doc fares as slice of ints, or an error
func GetAirPricingDocFaresByZoneNumber(zn int) ([]int, error) {
	if zn < 5 || zn > 9 {
		return nil, fmt.Errorf("the zone number %v is invalid, it doesn't contain doc fares", zn)
	}

	ap, err := model.GetAirPricingsFromDB()
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

// GetRoadPricingFaresByZoneNumber takes a zone number int as parameter
// and returns with the corresponding road pricing fares as slice of ints, or an error
func GetRoadPricingFaresByZoneNumber(zn int) ([]int, error) {
	if err := validateZoneNumber(zn); err != nil {
		return nil, err
	}

	rp, err := model.GetRoadPricingsFromDB()
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

func validateZoneNumber(zn int) error {
	if zn < 0 || zn > 9 {
		return fmt.Errorf("the zone number %v is invalid", zn)
	}

	return nil
}

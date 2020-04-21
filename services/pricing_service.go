package services

import (
	"fmt"

	"github.com/IstvanN/cashcalc-backend/properties"
)

// IncreaseWithVat takes a float64 and a percentage as parameter
// and returns with the vat-increased result
func IncreaseWithVat(num float64, vat float64) float64 {
	return num * (1 + (vat / 100))
}

// ValidateAirZoneNumber takes an int as parameter and checks if it between 0 and 9
// returns with error if not
func ValidateAirZoneNumber(zn int) error {
	airFaresZnMin := properties.Prop.GetInt(properties.AirFaresZnMin, 0)
	airFaresZnMax := properties.Prop.GetInt(properties.AirFaresZnMax, 9)
	if zn < airFaresZnMin || zn > airFaresZnMax {
		return fmt.Errorf("the zone number %v is invalid for air zones", zn)
	}
	return nil
}

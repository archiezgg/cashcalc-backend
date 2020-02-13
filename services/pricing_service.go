package services

import (
	"fmt"
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

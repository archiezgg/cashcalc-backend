package services

import (
	"fmt"

	"github.com/IstvanN/cashcalc-backend/model"
)

func getVatAsMultiplier() float64 {
	pv := getPricingVariablesFromDB()	
	return 1 + ()
}

// ValidateZoneNumber takes an int as parameter and checks if it between 0 and 9
// returns with error if not
func ValidateZoneNumber(zn int) error {
	if zn < 0 || zn > 9 {
		return fmt.Errorf("the zone number %v is invalid", zn)
	}

	return nil
}

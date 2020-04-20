package repositories

import (
	"fmt"
	"os"

	"github.com/IstvanN/cashcalc-backend/database"
	"github.com/IstvanN/cashcalc-backend/models"
)

var (
	pricingsCollectionName = os.Getenv("PRICINGS_COLL")
	pricingVarsCollName    = os.Getenv("PRICING_VARS_COLL")
)

// GetPricingsFromDB returns with Pricings struct containing both air and road pricings
func GetPricingsFromDB() (models.Pricings, error) {
	coll := database.GetCollectionByName(pricingsCollectionName)

	var p models.Pricings
	err := coll.Find(nil).One(&p)
	if err != nil {
		return models.Pricings{}, fmt.Errorf("error while retrieving collection %v from database: %v", pricingsCollectionName, err)
	}

	return p, nil
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

/*
	Cashcalc 2020
	Copyright (C) 2019-2020 Istvan Nemeth
	mailto: nemethistvanius@gmail.com
*/

package repositories

import (
	"fmt"

	"github.com/IstvanN/cashcalc-backend/database"
	"github.com/IstvanN/cashcalc-backend/models"
	"github.com/IstvanN/cashcalc-backend/properties"
)

// GetPricingVariables queries the db for
// the pricing variables that can be set by the admin
func GetPricingVariables() (models.PricingVariables, error) {
	coll := database.GetCollectionByName(properties.PricingVarsCollection)

	var pv models.PricingVariables
	err := coll.Find(nil).One(&pv)
	if err != nil {
		return models.PricingVariables{}, fmt.Errorf("error while retreiving collection %v from db: %v",
			properties.PricingVarsCollection, err)
	}

	return pv, nil
}

// UpdatePricingVariables updates the pricing variables in database
func UpdatePricingVariables(updatedPricingVars models.PricingVariables) error {
	currentPV, err := GetPricingVariables()
	if err != nil {
		return err
	}

	coll := database.GetCollectionByName(properties.PricingVarsCollection)
	if err := coll.Update(currentPV, updatedPricingVars); err != nil {
		return fmt.Errorf("error updating pricing variables: %v", err)
	}
	return nil
}

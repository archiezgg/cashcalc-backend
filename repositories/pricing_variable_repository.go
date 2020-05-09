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

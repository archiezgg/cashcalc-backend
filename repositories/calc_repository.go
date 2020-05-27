/*
	Cashcalc 2020
	Copyright (C) 2019-2020 Istvan Nemeth
	mailto: nemethistvanius@gmail.com
*/

package repositories

import (
	"github.com/IstvanN/cashcalc-backend/models"
	"github.com/IstvanN/cashcalc-backend/services"
)

// CalcResultAir takes input data and calculates the fares for air AND hungarian delivery
func CalcResultAir(inputData models.CalcInputData) (models.CalcOutputData, error) {
	if err := services.ValidateInputData(inputData); err != nil {
		return models.CalcOutputData{}, err
	}
	pricingVars, err := GetPricingVariables()
	if err != nil {
		return models.CalcOutputData{}, err
	}

	var pricingFare models.Fare
	if inputData.IsDocument {
		pricingFare, err = GetAirDocFaresByZoneNumberAndWeight(inputData.ZoneNumber, inputData.Weight)
		if err != nil {
			return models.CalcOutputData{}, err
		}
	} else {
		pricing
	}

	baseFare := services.CalcBaseFareWithVatAndDiscountAir(inputData.ZoneNumber, pricingVars.VATPercent, inputData.DiscountPercent, inputData.IsDocument, pricingFare)
	return models.CalcOutputData{}, nil
}

func getPricingFareBasedOnInputData(inputData models.CalcInputData) (models.Fare, error) {
	if inputData.IsDocument {
		pricingFare, err := GetAirDocFaresByZoneNumberAndWeight(inputData.ZoneNumber, inputData.Weight)
		if err != nil {
			return models.Fare{}, nil
		}
		return pricingFare, nil
	}

	pricingFare, err := GetAirFaresByZoneNumberAndWeight(inputData.ZoneNumber, inputData.Weight)
	if err != nil {
		return models.Fare{}, err
	}
	return pricingFare, nil
}

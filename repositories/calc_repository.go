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

	pricingFare, err := getPricingFareBasedOnInputData(inputData)
	if err != nil {
		return models.CalcOutputData{}, err
	}

	baseFare := services.CalcBaseFareWithVatAndDiscountAir(inputData.ZoneNumber, inputData.DiscountPercent, pricingVars.VATPercent, pricingFare.BaseFare)

	return models.CalcOutputData{
		BaseFare: baseFare,
	}, nil
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

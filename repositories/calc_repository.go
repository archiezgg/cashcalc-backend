/*
	Cashcalc 2020
	Copyright (C) 2019-2020 Istvan Nemeth
	mailto: nemethistvanius@gmail.com
*/

package repositories

import (
	"fmt"

	"github.com/IstvanN/cashcalc-backend/models"
	"github.com/IstvanN/cashcalc-backend/services"
)

// CalcResult takes input data and calculates the fares
func CalcResult(inputData models.CalcInputData) (models.CalcOutputData, error) {
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

	express, err := getExpressFareBasedOnInputData(inputData)
	if err != nil {
		return models.CalcOutputData{}, err
	}

	baseFare := services.CalcBaseFareWithVatAndDiscountAir(inputData.ZoneNumber, inputData.DiscountPercent,
		pricingVars.VATPercent, pricingFare.BaseFare)
	expressFare := services.CalcExpressFare(inputData.ZoneNumber, pricingVars.VATPercent, express)
	insuranceFare := services.CalcInsuranceFare(inputData.ZoneNumber, inputData.Insurance, pricingVars.InsuranceLimit,
		pricingVars.MinInsurance, pricingVars.VATPercent)
	extFare := services.CalcExtRasTk(inputData.IsExt, inputData.ZoneNumber, pricingVars.EXT, pricingVars.VATPercent)
	rasFare := services.CalcExtRasTk(inputData.IsRas, inputData.ZoneNumber, pricingVars.RAS, pricingVars.VATPercent)
	tkFare := services.CalcExtRasTk(inputData.IsTk, inputData.ZoneNumber, pricingVars.TK, pricingVars.VATPercent)
	fuelFare := services.CalcFuelFare(baseFare, expressFare, rasFare, pricingVars.AirFuelFarePercent)
	emergencyFare := services.CalcEmergencyFare(inputData.Weight, pricingVars.EmergencyFare)
	result := services.SumFares(baseFare, expressFare, insuranceFare, extFare, rasFare, tkFare, fuelFare, emergencyFare)

	return models.CalcOutputData{
		BaseFare:      baseFare,
		ExpressFare:   expressFare,
		InsuranceFare: insuranceFare,
		ExtFare:       extFare,
		RasFare:       rasFare,
		TkFare:        tkFare,
		FuelFare:      fuelFare,
		EmergencyFare: emergencyFare,
		Result:        result,
	}, nil
}

func getPricingFareBasedOnInputData(inputData models.CalcInputData) (models.Fare, error) {
	if inputData.TransferType == models.TransferAir {
		return getAirPricingFareBasedOnInputData(inputData)
	} else if inputData.TransferType == models.TransferRoad {
		return getRoadPricingFareBasedOnInputData(inputData)
	} else {
		return models.Fare{}, fmt.Errorf("invalid transfer type in input data: %v", inputData.TransferType)
	}
}

func getAirPricingFareBasedOnInputData(inputData models.CalcInputData) (models.Fare, error) {
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

func getRoadPricingFareBasedOnInputData(inputData models.CalcInputData) (models.Fare, error) {
	pricingFare, err := GetRoadFaresByZoneNumberAndWeight(inputData.ZoneNumber, inputData.Weight)
	if err != nil {
		return models.Fare{}, err
	}
	return pricingFare, nil
}

func getExpressFareBasedOnInputData(inputData models.CalcInputData) (float64, error) {
	pv, err := GetPricingVariables()
	if err != nil {
		return 0, err
	}

	// Hungary
	if inputData.ZoneNumber == 0 {
		if inputData.ExpressType == models.Express9h {
			return float64(pv.Express9hHungarian), nil
		}
		if inputData.ExpressType == models.Express12h {
			return float64(pv.Express12hHungarian), nil
		}
	}
	// Not Hungary
	if inputData.ExpressType == models.Express9h {
		return float64(pv.Express9h), nil
	}
	if inputData.ExpressType == models.Express12h {
		return float64(pv.Express12h), nil
	}
	// Express Worldwide
	return 0, nil
}

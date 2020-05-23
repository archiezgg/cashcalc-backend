/*
	Cashcalc 2020
	Copyright (C) 2019-2020 Istvan Nemeth
	mailto: nemethistvanius@gmail.com
*/

package services

import (
	"github.com/IstvanN/cashcalc-backend/models"
)

// CalcResultAir takes input data and calculates the fares for air AND hungarian delivery
func CalcResultAir(inputData models.CalcInputData) (models.CalcOutputData, error) {
	if err := validateInputData(inputData); err != nil {
		return models.CalcOutputData{}, err
	}
	return models.CalcOutputData{}, nil
}

func isZoneEU(zn int) bool {
	return zn <= 4 && zn >= 0
}

func calcBaseFareWithVatAndDiscountAir(zn int, vatPercent float64, discountPercent float64, isDocument bool, pricingFare models.Fare) float64 {
	baseFareIncreasedWithVat := IncreaseWithVat(float64(pricingFare.BaseFare), vatPercent)
	if isZoneEU(zn) {
		return applyDiscountToBaseFare(baseFareIncreasedWithVat, discountPercent)
	}

	if isDocument && pricingFare.Weight <= 2 {
		return applyDiscountToBaseFare(baseFareIncreasedWithVat, discountPercent)
	}

	return applyDiscountToBaseFare(baseFareIncreasedWithVat, discountPercent)
}

func applyDiscountToBaseFare(baseFare float64, discountPercent float64) float64 {
	return (1 - discountPercent/100) * baseFare
}

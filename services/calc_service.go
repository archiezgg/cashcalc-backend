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
func CalcResultAir(input models.CalcInputData) (models.CalcOutputData, error) {
	return models.CalcOutputData{}, nil
}

func isZoneEU(zn int) bool {
	return zn <= 4 && zn >= 0
}

func calcBaseFareWithDiscountAir(zn int, vat float64, pricingFare models.Fare) float64 {
	var result float64
	if isZoneEU(zn) {
		result = addDiscountToBaseFare(float64(pricingFare.BaseFare) * vat)
	}
}

func addDiscountToBaseFare(discountPercent int, baseFare float64) float64 {
	return float64(1-discountPercent/100) * baseFare
}

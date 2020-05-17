/*
	Cashcalc 2020
	Copyright (C) 2019-2020 Istvan Nemeth
	mailto: nemethistvanius@gmail.com
*/

package services

import (
	"github.com/IstvanN/cashcalc-backend/models"
)

func CalcResultAir(input models.CalcInputData) (models.CalcOutputData, error) {
	return models.CalcOutputData{}, nil
}

func isZoneEU(zn int) bool {
	return zn < 5
}

func calcBaseFareWithDiscountAir(zn int, vat float64, baseFare models.Fare) float64 {
	return 0.0
}

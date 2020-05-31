/*
	Cashcalc 2020
	Copyright (C) 2019-2020 Istvan Nemeth
	mailto: nemethistvanius@gmail.com
*/

package services

import (
	"fmt"
	"math"

	"github.com/IstvanN/cashcalc-backend/models"
)

// IsZoneEU returns if zone is EU or not
func isZoneEU(zn int) bool {
	return zn <= 4 && zn >= 0
}

// CalcBaseFareWithVatAndDiscountAir calculates the basefare increased by VAT and applied discount
func CalcBaseFareWithVatAndDiscountAir(zn int, discountPercent float64, vatPercent float64, baseFare int) float64 {
	if isZoneEU(zn) {
		baseFareIncreasedWithVat := IncreaseWithVat(float64(baseFare), vatPercent)
		return math.Round(applyDiscountToBaseFare(baseFareIncreasedWithVat, discountPercent))
	}

	return math.Round(applyDiscountToBaseFare(float64(baseFare), discountPercent))
}

// ValidateInputData takes an input data model and returns with an error if there is a logical error
func ValidateInputData(input models.CalcInputData) error {
	var err error
	if isZoneEU(input.ZoneNumber) && input.IsDocument {
		err = fmt.Errorf("zone number %v, document status %v: no document delivery to EU", input.ZoneNumber, input.IsDocument)
		return err
	}

	if input.IsDocument && input.Weight > 2 {
		err = fmt.Errorf("weight %v, document status %v: document cannot have more weight than 2", input.Weight, input.IsDocument)
		return err
	}
	return nil
}

func applyDiscountToBaseFare(baseFare float64, discountPercent float64) float64 {
	return math.Round((1 - discountPercent/100) * baseFare)
}

// CalcExpressFare increases express fare with VAT if zone is EU
func CalcExpressFare(zn int, vatPercent float64, expressFare float64) float64 {
	if isZoneEU(zn) {
		return math.Round(IncreaseWithVat(expressFare, vatPercent))
	}
	return expressFare
}

// CalcInsuranceFare calculates the insurance fare based on the limit, minimum fee and vat
func CalcInsuranceFare(zn, insurance, limit, min int, vatPercent float64) float64 {
	if insurance == 0 {
		return 0
	}

	if insurance < limit {
		if isZoneEU(zn) {
			return math.Round(IncreaseWithVat(float64(min), vatPercent))
		}
		return float64(min)
	}

	if isZoneEU(zn) {
		return math.Round(0.01 * IncreaseWithVat(float64(insurance), vatPercent))
	}
	return math.Round(0.01 * float64(insurance))
}

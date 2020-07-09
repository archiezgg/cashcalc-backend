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

// ValidateInputData takes an input data model and returns with an error if there is a logical error
func ValidateInputData(input models.CalcInputData) error {
	var err error
	if input.TransferType != models.TransferAir && input.TransferType != models.TransferRoad {
		err = fmt.Errorf("transfer type could either be %v or %v, but got %v",
			models.TransferAir, models.TransferRoad, input.TransferType)
		return err
	}

	if input.ExpressType != models.ExpressWorldwide && input.ExpressType != models.Express9h && input.ExpressType != models.Express12h {
		err = fmt.Errorf("express type could either be %v, %v or %v, but got %v",
			models.ExpressWorldwide, models.Express9h, models.Express12h, input.ExpressType)
		return err
	}

	if input.TransferType == models.TransferRoad &&
		(input.IsDocument || input.IsExt || input.ExpressType != models.ExpressWorldwide) {
		err = fmt.Errorf("road transfer type cannot be document delivery, EXT, express9h or express12h")
		return err
	}

	if isZoneEU(input.ZoneNumber) && input.IsDocument {
		err = fmt.Errorf("zone number %v, document status %v: no document delivery to EU",
			input.ZoneNumber, input.IsDocument)
		return err
	}

	if input.IsDocument && input.Weight > 2 {
		err = fmt.Errorf("weight %v, document status %v: document cannot have more weight than 2",
			input.Weight, input.IsDocument)
		return err
	}
	return nil
}

// CalcBaseFareWithVatAndDiscountAir calculates the basefare increased by VAT and applied discount
func CalcBaseFareWithVatAndDiscountAir(zn int, discountPercent float64, vatPercent float64, baseFare int) float64 {
	if isZoneEU(zn) {
		baseFareIncreasedWithVat := IncreaseWithVat(float64(baseFare), vatPercent)
		return math.Round(applyDiscountToBaseFare(baseFareIncreasedWithVat, discountPercent))
	}

	return math.Round(applyDiscountToBaseFare(float64(baseFare), discountPercent))
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

// CalcExtRasTk is a function for make Ext, Ras and TK calculations,
// calculates based on shouldCalc (isExt, isRas, isTk), zone number (EU or not), the fare and VAT
func CalcExtRasTk(shouldCalc bool, zn, fare int, vatPercent float64) float64 {
	if !shouldCalc {
		return 0
	}
	if isZoneEU(zn) {
		return math.Round(IncreaseWithVat(float64(fare), vatPercent))
	}
	return float64(fare)
}

// CalcFuelFare calculates fuel fare based on base fare, express, ras and the fuelPercent given
func CalcFuelFare(baseFare, expressFare, rasFare, fuelPercent float64) float64 {
	return math.Round((fuelPercent / 100) * (baseFare + expressFare + rasFare))
}

// CalcEmergencyFare calculates fee based on whether there is emergency (COVID)
// after every started kg of weight, it calculates weight * emergency fee
func CalcEmergencyFare(isEmergency bool, weight float64, emergencyFee int) float64 {
	if !isEmergency {
		return 0
	}

	if weight == math.Round(weight) {
		return math.Round(float64(emergencyFee) * weight)
	}
	return math.Round(float64(emergencyFee) * float64(int(weight)+1)) // converting weight to integer rounds it down
}

// SumFares takes any number of float64 numbers, and adds them together
func SumFares(fares ...float64) float64 {
	var result float64
	for _, f := range fares {
		result += f
	}
	return result
}

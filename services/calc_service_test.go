/*
	Cashcalc 2020
	Copyright (C) 2019-2020 Istvan Nemeth
	mailto: nemethistvanius@gmail.com
*/

package services

import (
	"testing"

	"github.com/IstvanN/cashcalc-backend/models"
)

func TestIsZoneEu(t *testing.T) {
	testCases := []struct {
		input    int
		expected bool
	}{
		{0, true},
		{4, true},
		{-1, false},
		{5, false},
		{8, false},
	}

	for _, tc := range testCases {
		actual := isZoneEU(tc.input)
		if actual != tc.expected {
			t.Errorf("isZoneEu(%v) failed: expected %v, got %v", tc.input, tc.expected, actual)
		}
	}
}

func TestApplyDiscountToBaseFare(t *testing.T) {
	testCases := []struct {
		inputBaseFare float64
		inputDiscount float64
		expected      float64
	}{
		{100, 10, 90},
		{200, 10, 180},
		{100, 90, 10},
	}

	for _, tc := range testCases {
		actual := applyDiscountToBaseFare(tc.inputBaseFare, tc.inputDiscount)
		if tc.expected != actual {
			t.Errorf("applyDiscountToBaseFare(%v, %v) failed: expected %v, got %v", tc.inputBaseFare, tc.inputDiscount, tc.expected, actual)
		}
	}
}

func TestCalcBaseFareWithVatAndDiscountAir(t *testing.T) {
	testCases := []struct {
		zoneNumber                  int
		discountPercent, vatPercent float64
		baseFare                    int
		expected                    float64
	}{
		{0, 20, 27, 3371, 3425},
		{1, 10, 27, 14992, 17136},
		{5, 30, 27, 57184, 40029},
		{6, 10, 27, 21093, 18984},
	}

	for _, tc := range testCases {
		actual := CalcBaseFareWithVatAndDiscountAir(tc.zoneNumber, tc.discountPercent, tc.vatPercent, tc.baseFare)
		if tc.expected != actual {
			t.Errorf("CalcBaseFareWithVatAndDiscountAir(%v, %v, %v, %v) failed: expected %v, got %v", tc.zoneNumber, tc.discountPercent, tc.vatPercent, tc.baseFare,
				tc.expected, actual)
		}
	}
}

func TestValidateInputData(t *testing.T) {
	testCases := []struct {
		inputData   models.CalcInputData
		expectError bool
	}{
		{models.CalcInputData{TransferType: "road", ZoneNumber: 0, IsDocument: true, Weight: 1.5, ExpressType: models.ExpressWorldwide}, true},
		{models.CalcInputData{TransferType: "air", ZoneNumber: 3, IsDocument: false, Weight: 2, ExpressType: models.Express9h}, false},
		{models.CalcInputData{TransferType: "road", ZoneNumber: 5, IsDocument: true, Weight: 2.5, ExpressType: models.ExpressWorldwide}, true},
		{models.CalcInputData{TransferType: "air", ZoneNumber: 6, IsDocument: false, Weight: 1.5, ExpressType: models.ExpressWorldwide}, false},
		{models.CalcInputData{TransferType: "air", ZoneNumber: 7, IsDocument: true, Weight: 2, ExpressType: models.Express12h}, false},
		{models.CalcInputData{TransferType: "", Weight: 0.5, ExpressType: models.ExpressWorldwide}, true},
		{models.CalcInputData{TransferType: "non-valid", Weight: 0.5, ExpressType: models.ExpressWorldwide}, true},
		{models.CalcInputData{TransferType: "road", IsDocument: false, Weight: 0.5, IsExt: true, ExpressType: models.ExpressWorldwide}, true},
		{models.CalcInputData{TransferType: "road", IsDocument: true, Weight: 0.5, IsExt: false, ExpressType: models.ExpressWorldwide}, true},
		{models.CalcInputData{TransferType: "road", IsDocument: false, Weight: 0.5, IsExt: false, ExpressType: models.ExpressWorldwide}, false},
		{models.CalcInputData{TransferType: "air", ZoneNumber: 0, IsDocument: true, Weight: 0.5, IsExt: true, ExpressType: models.ExpressWorldwide}, true},
		{models.CalcInputData{TransferType: "air", ZoneNumber: 5, IsDocument: true, Weight: 0.5, IsExt: true, ExpressType: models.ExpressWorldwide}, false},
		{models.CalcInputData{TransferType: "road", Weight: 0.5, ExpressType: models.Express9h}, true},
		{models.CalcInputData{TransferType: "road", Weight: 0.5, ExpressType: models.Express12h}, true},
		{models.CalcInputData{TransferType: "air", Weight: 0.5, ExpressType: ""}, true},
		{models.CalcInputData{TransferType: "air", Weight: 0.5, ExpressType: "non-valid"}, true},
		{models.CalcInputData{TransferType: "air", Weight: 0}, true},
		{models.CalcInputData{TransferType: "road", Weight: 0.5, Insurance: -1000}, true},
	}

	for _, tc := range testCases {
		err := ValidateInputData(tc.inputData)
		actual := (err != nil)
		if tc.expectError != actual {
			t.Errorf("ValidateInputData(%v) failed: expected error %v, got error %v", tc.inputData, tc.expectError, err)
		}
	}
}

func TestCalcExpressFare(t *testing.T) {
	testCases := []struct {
		zoneNumber                        int
		vatPercent, expressFare, expected float64
	}{
		{0, 27, 0, 0},
		{7, 27, 0, 0},
		{3, 27, 1000, 1270},
		{5, 27, 1000, 1000},
	}

	for _, tc := range testCases {
		actual := CalcExpressFare(tc.zoneNumber, tc.vatPercent, tc.expressFare)
		if actual != tc.expected {
			t.Errorf("CalcExpressFare(%v, %v, %v) failed: expected %v, got %v", tc.zoneNumber, tc.vatPercent, tc.expressFare, tc.expected, actual)
		}
	}
}

func TestCalcInsuranceFare(t *testing.T) {
	testCases := []struct {
		zoneNumber, insurance, limit, min int
		vatPercent, expected              float64
	}{
		{0, 0, 150, 20, 27, 0},
		{2, 100, 150, 20, 27, 25},
		{5, 100, 150, 20, 27, 20},
		{3, 200, 150, 20, 27, 3},
		{6, 200, 150, 20, 27, 2},
	}

	for _, tc := range testCases {
		actual := CalcInsuranceFare(tc.zoneNumber, tc.insurance, tc.limit, tc.min, tc.vatPercent)
		if actual != tc.expected {
			t.Errorf("CalcInsuranceFare(%v, %v, %v, %v, %v) failed: expected %v, got %v", tc.zoneNumber, tc.insurance, tc.limit, tc.min, tc.vatPercent,
				tc.expected, actual)
		}
	}
}

func TestCalcExtRasTk(t *testing.T) {
	testCases := []struct {
		shouldCalc       bool
		zoneNumber, fare int
		vatPercent       float64
		expected         float64
	}{
		{true, 0, 1000, 27, 1270},
		{true, 5, 1000, 27, 1000},
		{false, 3, 1000, 27, 0},
		{false, 7, 1200, 27, 0},
	}

	for _, tc := range testCases {
		actual := CalcExtRasTk(tc.shouldCalc, tc.zoneNumber, tc.fare, tc.vatPercent)
		if actual != tc.expected {
			t.Errorf("CalcExtRasTk(%v, %v, %v, %v) failed: expected %v, got %v", tc.shouldCalc, tc.zoneNumber, tc.fare, tc.vatPercent,
				tc.expected, actual)
		}
	}
}

func TestCalcFuelFare(t *testing.T) {
	testCases := []struct {
		baseFare, expressFare, rasFare, fuelPercent, expected float64
	}{
		{100, 0, 0, 10, 10},
		{50, 10, 15, 0, 0},
		{50, 10, 15, 10, 8},
	}

	for _, tc := range testCases {
		actual := CalcFuelFare(tc.baseFare, tc.expressFare, tc.rasFare, tc.fuelPercent)
		if actual != tc.expected {
			t.Errorf("CalcFuelFare(%v, %v, %v, %v) failed: expected %v, got %v", tc.baseFare, tc.expressFare, tc.rasFare, tc.fuelPercent,
				tc.expected, actual)
		}
	}
}

func TestCalcEmergencyFare(t *testing.T) {
	testCases := []struct {
		weight       float64
		emergencyFee int
		expected     float64
	}{
		{6, 10, 60},
		{6.1, 10, 70},
		{6.8, 10, 70},
		{7, 10, 70},
		{7.1, 10, 80},
		{7.8, 10, 80},
	}

	for _, tc := range testCases {
		actual := CalcEmergencyFare(tc.weight, tc.emergencyFee)
		if actual != tc.expected {
			t.Errorf("CalcEmergencyFare(%v, %v) failed: expected %v, got %v", tc.weight, tc.emergencyFee,
				tc.expected, actual)
		}
	}
}

func TestSumFares(t *testing.T) {
	testCases := []struct {
		a, b, c, d, e, f, expected float64
	}{
		{1, 0, 2, 0, 3, 0, 6},
		{1, 1, 1, 1, 1, 1, 6},
		{1, 10, 100, 1000, 10000, 0, 11111},
		{0, 0, 0, 0, 0, 0, 0},
	}

	for _, tc := range testCases {
		actual := SumFares(tc.a, tc.b, tc.c, tc.d, tc.e, tc.f)
		if actual != tc.expected {
			t.Errorf("SumFares(%v, %v, %v, %v, %v, %v) failed: expected %v, got %v", tc.a, tc.b, tc.c, tc.d, tc.e, tc.f,
				tc.expected, actual)
		}
	}
}

func TestAreAllIntegersPositive(t *testing.T) {
	testCases := []struct {
		a, b, c, d, e, f float64
		expected         bool
	}{
		{0, 1, 2.2, 3, 4, 5.6, true},
		{-1, 0, 1.5, 2, 3, 4, false},
		{66, 78, 100000, -567.5, 88, 9999, false},
		{-7, -88, -89, -99, -10000, -67, false},
		{0, 0, 0, 0, 0, 0, true},
	}

	for _, tc := range testCases {
		actual := areAllNumbersPositive(tc.a, tc.b, tc.c, tc.d, tc.e, tc.f)
		if actual != tc.expected {
			t.Errorf("AreAllNumbersPositive(%v, %v, %v, %v, %v, %v) failed: expected %v, got %v", tc.a, tc.b, tc.c, tc.d, tc.e, tc.f,
				tc.expected, actual)
		}
	}
}

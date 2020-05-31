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
		{models.CalcInputData{ZoneNumber: 0, IsDocument: true, Weight: 1.5}, true},
		{models.CalcInputData{ZoneNumber: 3, IsDocument: false, Weight: 2}, false},
		{models.CalcInputData{ZoneNumber: 5, IsDocument: true, Weight: 2.5}, true},
		{models.CalcInputData{ZoneNumber: 6, IsDocument: false, Weight: 1.5}, false},
		{models.CalcInputData{ZoneNumber: 7, IsDocument: true, Weight: 2}, false},
	}

	for _, tc := range testCases {
		err := ValidateInputData(tc.inputData)
		actual := (err != nil)
		if tc.expectError != actual {
			t.Errorf("ValidateInputData(%v) failed: expected error %v, got error %v", tc.inputData, tc.expectError, actual)
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

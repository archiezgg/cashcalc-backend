/*
	Cashcalc 2020
	Copyright (C) 2019-2020 Istvan Nemeth
	mailto: nemethistvanius@gmail.com
*/

package services

import "testing"

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

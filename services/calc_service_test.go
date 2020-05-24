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
	}

	for _, tc := range testCases {
		actual := applyDiscountToBaseFare(tc.inputBaseFare, tc.inputDiscount)
		if tc.expected != actual {
			t.Errorf("applyDiscountToBaseFare(%v, %v) failed: expected %v, got %v", tc.inputBaseFare, tc.inputDiscount, tc.expected, actual)
		}
	}
}

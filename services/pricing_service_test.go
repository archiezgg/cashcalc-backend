/*
	Cashcalc 2020
	Copyright (C) 2019-2020 Istvan Nemeth
	mailto: nemethistvanius@gmail.com
*/

package services

import (
	"testing"
)

func TestIncreaseWithVat(t *testing.T) {
	testCases := []struct {
		input      float64
		vatPercent float64
		expected   float64
	}{
		{0, 27, 0},
		{1, 0, 1},
		{1, 27, 1.27},
		{50, 30, 65},
	}

	for _, tc := range testCases {
		output := IncreaseWithVat(tc.input, tc.vatPercent)
		if output != tc.expected {
			t.Errorf("IncreaseWithVat(%v, %v) failed: expected: %v, got: %v", tc.input, tc.vatPercent, tc.expected, output)
		}
	}
}
func TestIsZoneNumberInvalid(t *testing.T) {
	testCases := []struct {
		zn       int
		min      int
		max      int
		expected bool
	}{
		{5, 2, 100, false},
		{55, 1, 200, false},
		{5, 5, 10, false},
		{5, 1, 5, false},
		{-1, 0, 200, true},
		{0, 1, 50, true},
		{10, 11, 100, true},
	}

	for _, tc := range testCases {
		actual := IsZoneNumberInvalid(tc.zn, tc.min, tc.max)
		if actual != tc.expected {
			t.Errorf("IsZoneNumberValid(%v, %v, %v) failed: expected %v, got: %v", tc.zn, tc.min, tc.max, tc.expected, actual)
		}
	}
}

func TestIsWeightInvalid(t *testing.T) {
	testCases := []struct {
		weight   float64
		min      float64
		max      float64
		expected bool
	}{
		{0.5, 0, 100, false},
		{50.5, 1, 100, false},
		{100, 100, 200, false},
		{100, 50, 100, false},
		{3.5, 5, 10, true},
		{9, 10, 20, true},
	}

	for _, tc := range testCases {
		actual := IsWeightInvalid(tc.weight, tc.min, tc.max)
		if actual != tc.expected {
			t.Errorf("IsWeightInvalid(%v, %v, %v) failed: expected %v, got: %v", tc.weight, tc.min, tc.max, tc.expected, actual)
		}
	}
}

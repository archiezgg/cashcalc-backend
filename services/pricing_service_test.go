package services

import (
	"errors"
	"reflect"
	"testing"
)

func TestIncreaseWithVat(t *testing.T) {
	testCases := []struct {
		input    float64
		vat      float64
		expected float64
	}{
		{0, 27, 0},
		{1, 0, 1},
		{1, 27, 1.27},
		{50, 30, 65},
	}

	for _, tc := range testCases {
		output := IncreaseWithVat(tc.input, tc.vat)
		if output != tc.expected {
			t.Errorf("IncreaseWithVat(%v, %v) failed: expected: %v, got: %v", tc.input, tc.vat, tc.expected, output)
		}
	}
}
func TestValidateAirFaresZoneNumber(t *testing.T) {
	e := errors.New("")
	testCases := []struct {
		x   int
		err error
	}{
		{0, nil},
		{9, nil},
		{3, nil},
		{5, nil},
		{-1, e},
		{10, e},
	}

	for _, tc := range testCases {
		err := ValidateAirFaresZoneNumber(tc.x)
		if reflect.TypeOf(err) != reflect.TypeOf(tc.err) {
			t.Errorf("ValidateZoneNumber(%v) failed: expected type: %T, got: %T", tc.x, tc.err, err)
		}
	}
}

func TestValidateAirDocFaresZoneNumber(t *testing.T) {
	e := errors.New("")
	testCases := []struct {
		x   int
		err error
	}{
		{5, nil},
		{9, nil},
		{6, nil},
		{1, e},
		{10, e},
		{4, e},
		{-1, e},
		{11, e},
	}

	for _, tc := range testCases {
		err := ValidateAirDocFaresZoneNumber(tc.x)
		if reflect.TypeOf(err) != reflect.TypeOf(tc.err) {
			t.Errorf("ValidateAirDocFaresZoneNumber(%v) failed: expected type: %T, got: %T", tc.x, tc.err, err)
		}
	}
}

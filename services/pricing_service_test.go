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
			t.Errorf("ValidateAirFaresZoneNumber(%v) failed: expected type: %T, got: %T", tc.x, tc.err, err)
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

func TestValidateRoadFaresZoneNumber(t *testing.T) {
	e := errors.New("")
	testCases := []struct {
		x   int
		err error
	}{
		{1, nil},
		{5, nil},
		{3, nil},
		{0, e},
		{6, e},
		{-1, e},
		{11, e},
	}

	for _, tc := range testCases {
		err := ValidateRoadFaresZoneNumber(tc.x)
		if reflect.TypeOf(err) != reflect.TypeOf(tc.err) {
			t.Errorf("ValidateRoadFaresZoneNumber(%v) failed: expected type: %T, got: %T", tc.x, tc.err, err)
		}
	}
}

func TestValidateAirFaresWeight(t *testing.T) {
	e := errors.New("")
	testCases := []struct {
		x   float64
		err error
	}{
		{0.5, nil},
		{200, nil},
		{4.5, nil},
		{70, nil},
		{0, e},
		{201, e},
		{-1, e},
	}

	for _, tc := range testCases {
		err := ValidateAirFaresWeight(tc.x)
		if reflect.TypeOf(err) != reflect.TypeOf(tc.err) {
			t.Errorf("ValidateAirFaresWeight(%v) failed: expected type: %T, got: %T", tc.x, tc.err, err)
		}
	}
}

func TestValidateAirDocFaresWeight(t *testing.T) {
	e := errors.New("")
	testCases := []struct {
		x   float64
		err error
	}{
		{0.5, nil},
		{2, nil},
		{1, nil},
		{1.5, nil},
		{0, e},
		{2.5, e},
		{-1, e},
	}

	for _, tc := range testCases {
		err := ValidateAirDocFaresWeight(tc.x)
		if reflect.TypeOf(err) != reflect.TypeOf(tc.err) {
			t.Errorf("ValidateAirDocFaresWeight(%v) failed: expected type: %T, got: %T", tc.x, tc.err, err)
		}
	}
}

func TestValidateRoadFaresWeight(t *testing.T) {
	e := errors.New("")
	testCases := []struct {
		x   float64
		err error
	}{
		{1, nil},
		{100, nil},
		{55, nil},
		{32, nil},
		{0, e},
		{101, e},
		{-1, e},
	}

	for _, tc := range testCases {
		err := ValidateRoadFaresWeight(tc.x)
		if reflect.TypeOf(err) != reflect.TypeOf(tc.err) {
			t.Errorf("ValidateRoadFaresWeight(%v) failed: expected type: %T, got: %T", tc.x, tc.err, err)
		}
	}
}

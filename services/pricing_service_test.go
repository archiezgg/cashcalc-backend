package services

import (
	"errors"
	"reflect"
	"testing"
)

func TestValidateZoneNumber(t *testing.T) {
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
		err := ValidateZoneNumber(tc.x)
		if reflect.TypeOf(err) != reflect.TypeOf(tc.err) {
			t.Errorf("validateZoneNumber(%v) failed: expected type: %T, got: %T", tc.x, tc.err, err)
		}
	}
}

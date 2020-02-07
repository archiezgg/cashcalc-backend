package service

import "testing"

import "fmt"

import "reflect"

func TestValidateZoneNumber(t *testing.T) {
	errorMsg := "the zone number %v is invalid"
	tables := []struct {
		x   int
		err error
	}{
		{0, nil},
		{9, nil},
		{10, fmt.Errorf(errorMsg, 10)},
	}

	for _, table := range tables {
		err := validateZoneNumber(table.x)
		if reflect.TypeOf(table.x) != reflect.TypeOf(table.err) {
			t.Errorf("validateZoneNumber(%v) failed: expected: %T, got: %T", table.x, table.err, err)
		}
	}
}

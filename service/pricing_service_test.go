package service

import (
	"errors"
	"reflect"
	"testing"
)

func TestValidateZoneNumber(t *testing.T) {
	e := errors.New("")
	tables := []struct {
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

	for _, table := range tables {
		err := validateZoneNumber(table.x)
		if reflect.TypeOf(err) != reflect.TypeOf(table.err) {
			t.Errorf("validateZoneNumber(%v) failed: expected type: %T, got: %T", table.x, table.err, err)
		}
	}
}

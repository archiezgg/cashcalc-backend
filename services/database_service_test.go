/*
	Cashcalc 2020
	Copyright (C) 2019-2020 Istvan Nemeth
	mailto: nemethistvanius@gmail.com
*/

package services

import "testing"

func TestGetMongoDBNameFromURI(t *testing.T) {
	testCases := []struct {
		input    string
		expected string
	}{
		{"mongodb://user:pw@host:port/db", "db"},
		{"mongodb://user:pw@host:port//db", "db"},
		{"mongodb://host/db", "db"},
		{"mongo://token@host/db", "db"},
		{"db", "db"},
		{"/db", "db"},
		{"/host:port/db", "db"},
		{"mongodb://user:pw@host:port/db?ssl=true", "db"},
	}

	for _, tc := range testCases {
		actual := GetMongoDBNameFromURI(tc.input)
		if actual != tc.expected {
			t.Errorf("getDBnameFromURI(%v) failed: expected: %v, got: %v", tc.input, tc.expected, actual)
		}
	}
}

func TestGetPostgresDBSpecsFromURL(t *testing.T) {
	testCases := []struct {
		input, expectedUser, expectedPw, expectedHost, expectedPort, expectedDBName string
	}{
		{"protocol://user:pw@host:port/db", "user", "pw", "host", "port", "db"},
	}

	for _, tc := range testCases {
		actualHost, actualUser, actualPort, actualPw, actualDB := GetPostgresDBSpecsFromURL(tc.input)
		if actualHost != tc.expectedHost && actualUser != tc.expectedUser &&
			actualPort != tc.expectedPort && actualPw != tc.expectedPw && actualDB != tc.expectedDBName {
			t.Errorf("GetPostgresDBSpecsFromURL(%v) failed: exptected: protocol://%v:%v@%v:%v/%v, got: protocol://%v:%v@%v:%v/%v",
				tc.input, tc.expectedUser, tc.expectedPw, tc.expectedHost, tc.expectedPort, tc.expectedDBName,
				actualUser, actualPw, actualHost, actualPort, actualDB)
		}
	}
}

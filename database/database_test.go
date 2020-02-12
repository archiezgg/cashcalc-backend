package database

import "testing"

func TestGetDBNameFromURI(t *testing.T) {
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
	}

	for _, tc := range testCases {
		actual := getDBNameFromURI(tc.input)
		if actual != tc.expected {
			t.Errorf("getDBnameFromURI(%v) failed: expected: %v, got: %v", tc.input, tc.expected, actual)
		}
	}
}

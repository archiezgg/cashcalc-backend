package services

import "strings"

// GetDBNameFromURI splits the given URI and returns the name of the DB
func GetDBNameFromURI(uri string) string {
	splitURI := strings.SplitAfter(uri, "/")
	return splitURI[len(splitURI)-1]
}

/*
	Cashcalc 2020
	Copyright (C) 2019-2020 Istvan Nemeth
	mailto: nemethistvanius@gmail.com
*/

package services

import (
	"strings"
)

// GetMongoDBNameFromURI splits the given URI and returns the name of the DB
func GetMongoDBNameFromURI(uri string) string {
	splitURI := strings.SplitAfter(uri, "/")
	dbName := splitURI[len(splitURI)-1]

	if strings.Contains(dbName, "?") {
		dbName = strings.Split(dbName, "?")[0]
	}

	return dbName
}

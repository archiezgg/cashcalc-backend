/*
	Cashcalc 2020
	Copyright (C) 2019-2020 Istvan Nemeth
	mailto: nemethistvanius@gmail.com
*/

package models

// Country stores the countries with name and a zone number
type Country struct {
	Name       string `bson:"name" json:"name"`
	ZoneNumber int    `bson:"zoneNumber" json:"zoneNumber"`
}

// Countries stores both air and road lists as fields
type Countries struct {
	CountriesAir  []Country `bson:"countriesAir" json:"countriesAir"`
	CountriesRoad []Country `bson:"countriesRoad" json:"countriesRoad"`
}

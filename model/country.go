package model

// Country stores the countries with name and a zone number
type Country struct {
	Name       string `bson:"name"`
	ZoneNumber int    `bson:"zoneNumber"`
}

// Countries stores both air and road lists as fields
type Countries struct {
	CountriesAir  []Country `bson:"countriesAir"`
	CountriesRoad []Country `bson:"countriesRoad"`
}

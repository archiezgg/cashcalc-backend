package model

import (
	"log"
	"encoding/json"
	"io/ioutil"
	"os"
)

const (
	airCountriesJSON = "data/countries_air.json"
)

// Country stores the countries with name and a zone number
type Country struct {
	Name       string `json:"name"`
	ZoneNumber int    `json:"zone_number"`
}

// Countries has a list of country types
type Countries struct {
	Countries []Country `json:"countries"`
}

// GetAirCountriesFromJSON returns all countries from a data JSON file
func GetAirCountriesFromJSON() (countries Countries) {
	dataFile, err := os.Open(airCountriesJSON)
	if err != nil {
		log.Fatalln("error opening air country data:", err)
	}
	defer dataFile.Close()
	b, err := ioutil.ReadAll(dataFile)
	if err != nil {
		log.Fatalln("failed to read air country data into bytes:", err)
	}

	json.Unmarshal(b, &countries)
	return
}




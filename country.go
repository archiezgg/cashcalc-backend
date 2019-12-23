package main

import (
	"log"
	"encoding/json"
	"io/ioutil"
	"os"
)

const (
	airCountriesJSON = "data/countries_air.json"
)

type country struct {
	Name       string `json:"name"`
	ZoneNumber int    `json:"zone_number"`
}

type countries struct {
	Countries []country `json:"countries"`
}

func getAirCountriesFromJSON() (countries countries) {
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




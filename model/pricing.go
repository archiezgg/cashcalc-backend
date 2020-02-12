package model

import (
	"fmt"
	"os"

	"github.com/IstvanN/cashcalc-backend/database"
)

var (
	pricingsCollectionName = os.Getenv("PRICINGS_COLL")
)

// Pricing is the struct to store zone numbers and the matching fares
type Pricing struct {
	ZoneNumber int   `bson:"zoneNumber"`
	Fares      []int `bson:"fares"`
	DocFares   []int `bson:"docFares"`
}

//Pricings stores both air and road pricings as fields
type Pricings struct {
	AirPricings  []Pricing `bson:"airPricings"`
	RoadPricings []Pricing `bson:"roadPricings"`
}

// GetAirPricingsFromDB returns with a slice of all elements of the air pricings collection, or an error
func GetAirPricingsFromDB() ([]Pricing, error) {
	p, err := getPricingsFromDB()
	if err != nil {
		return nil, err
	}
	return p.AirPricings, nil
}

// GetRoadPricingsFromDB returns with a slice of all elements of the road pricings collection or an error
func GetRoadPricingsFromDB() ([]Pricing, error) {
	p, err := getPricingsFromDB()
	if err != nil {
		return nil, err
	}
	return p.RoadPricings, nil
}

func getPricingsFromDB() (Pricings, error) {
	coll := database.GetCollectionByName(pricingsCollectionName)

	var p Pricings
	err := coll.Find(nil).One(&p)
	if err != nil {
		return Pricings{}, fmt.Errorf("error while retrieving collection %v from database: %v", pricingsCollectionName, err)
	}

	return p, nil
}

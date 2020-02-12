package model

import (
	"fmt"
	"os"

	"github.com/IstvanN/cashcalc-backend/database"
)

var (
	airPricingsCollectionName  = os.Getenv("PRICINGS_AIR_COLL")
	roadPricingsCollectionName = os.Getenv("PRICINGS_ROAD_COLL")
)

// Pricing is the struct to store zone numbers and the matching fares
type Pricing struct {
	ZoneNumber      int
	Fares, DocFares []int
}

// GetAirPricingsFromDB returns with a slice of all elements of the air pricings collection, or an error
func GetAirPricingsFromDB() ([]Pricing, error) {
	coll := database.GetCollectionByName(airPricingsCollectionName)

	var airPricings []Pricing
	err := coll.Find(nil).All(&airPricings)
	if err != nil {
		return nil, fmt.Errorf("error while retrieving collection %v from database: %v", airPricingsCollectionName, err)
	}
	return airPricings, nil
}

// GetRoadPricingsFromDB returns with a slice of all elements of the road pricings collection or an error
func GetRoadPricingsFromDB() ([]Pricing, error) {
	coll := database.GetCollectionByName(roadPricingsCollectionName)

	var roadPricings []Pricing
	err := coll.Find(nil).All(&roadPricings)
	if err != nil {
		return nil, fmt.Errorf("error while retrieving collection %v from database: %v", roadPricingsCollectionName, err)
	}
	return roadPricings, nil
}

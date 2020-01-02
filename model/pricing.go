package model

import (
	"context"
	"fmt"
	"log"
	"os"

	"github.com/IstvanN/cashcalc-backend/database"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo/options"
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

// GetAirPricingsFromDB returns with a slice of all elements of the air pricings collection
func GetAirPricingsFromDB() ([]Pricing, error) {
	coll := database.GetCollectionByName(airPricingsCollectionName)

	cur, err := coll.Find(context.TODO(), bson.D{{}}, options.Find())
	defer cur.Close(context.TODO())
	if err != nil {
		return nil, fmt.Errorf("retrieving collection %v failed: %v", airPricingsCollectionName, err)
	}

	var airPricings []Pricing
	for cur.Next(context.TODO()) {
		var p Pricing
		err := cur.Decode(&p)
		if err != nil {
			return nil, fmt.Errorf("error while decoding air pricing: %v", err)
		}
		airPricings = append(airPricings, p)
	}

	return airPricings, nil
}

// GetAirPricingFaresByZoneNumber takes a zone number int as parameter and returns with the corresponding air pricing fares as slice of ints
func GetAirPricingFaresByZoneNumber(zn int) []int {
	ap := GetAirPricingsFromDB()

	for _, p := range ap {
		if p.ZoneNumber == zn {
			return p.Fares
		}
	}
	log.Printf("zone number '%v' is invalid\n", zn)
	return nil
}

// GetRoadPricingsFromDB returns with a slice of all elements of the road pricings collection
func GetRoadPricingsFromDB() ([]Pricing, error) {
	coll := database.GetCollectionByName(roadPricingsCollectionName)

	cur, err := coll.Find(context.TODO(), bson.D{{}}, options.Find())
	defer cur.Close(context.TODO())
	if err != nil {
		return nil, fmt.Errorf("retrieving collection %v failed: %v", roadPricingsCollectionName, err)
	}

	var roadPricings []Pricing
	for cur.Next(context.TODO()) {
		var p Pricing
		err := cur.Decode(&p)
		if err != nil {
			return nil, fmt.Errorf("error while decoding road pricing: %v", err)
		}
		roadPricings = append(roadPricings, p)
	}

	return roadPricings, nil
}

// GetRoadPricingFaresByZoneNumber takes a zone number int as parameter and returns with the corresponding road pricing fares as slice of ints
func GetRoadPricingFaresByZoneNumber(zn int) []int {
	rp := GetRoadPricingsFromDB()

	for _, p := range rp {
		if p.ZoneNumber == zn {
			return p.Fares
		}
	}
	log.Printf("zone number '%v' is invalid\n", zn)
	return nil
}

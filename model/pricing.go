package model

import (
	"context"
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
func GetAirPricingsFromDB() []Pricing {
	coll := database.GetCollectionByName(airPricingsCollectionName)
	cur, err := coll.Find(context.TODO(), bson.D{{}}, options.Find())
	if err != nil {
		log.Printf("retrieving collection %v failed: %v\n", airPricingsCollectionName, err)
	}

	var airPricings []Pricing
	for cur.Next(context.TODO()) {
		var p Pricing
		err := cur.Decode(&p)
		if err != nil {
			log.Println("error while decoding air pricing: ", err)
		} else {
			airPricings = append(airPricings, p)
		}
	}
	cur.Close(context.TODO())

	return airPricings
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

// GetRoadPricingsFromDB returns with a slice of all elements of the air pricings collection
func GetRoadPricingsFromDB() []Pricing {
	coll := database.GetCollectionByName(roadPricingsCollectionName)
	cur, err := coll.Find(context.TODO(), bson.D{{}}, options.Find())
	if err != nil {
		log.Printf("retrieving collection %v failed: %v\n", roadPricingsCollectionName, err)
	}

	var roadPricings []Pricing
	for cur.Next(context.TODO()) {
		var p Pricing
		err := cur.Decode(&p)
		if err != nil {
			log.Println("error while decoding road pricing: ", err)
		} else {
			roadPricings = append(roadPricings, p)
		}
	}
	cur.Close(context.TODO())

	return roadPricings
}

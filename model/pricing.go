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
	airPricingsCollectionName = os.Getenv("PRICINGS_AIR_COLL")
)

// Pricing is the struct to store zone numbers and the matching fares
type Pricing struct {
	ZoneNumber      int
	Fares, DocFares []int
}

// GetAirPricingsFromDB returns with a slice of all elements of the air pricings collection
func GetAirPricingsFromDB() (airPricings []Pricing) {
	coll := database.GetCollectionByName(airPricingsCollectionName)
	cur, err := coll.Find(context.TODO(), bson.D{{}}, options.Find())
	if err != nil {
		log.Printf("retrieving collection %v failed: %v\n", airCountriesCollectionName, err)
	}

	for cur.Next(context.TODO()) {
		var p Pricing
		err := cur.Decode(&p)
		if err != nil {
			log.Println("error while retrieving air pricings, ", err)
		} else {
			airPricings = append(airPricings, p)
		}
	}
	cur.Close(context.TODO())
	return
}

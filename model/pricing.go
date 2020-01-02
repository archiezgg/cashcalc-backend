package model

import (
	"os"
)

var (
	airPricingsCollectionName = os.Getenv("PRICINGS_AIR_COLL")
)

// Pricing is the struct to store zone numbers and the matching fares
type Pricing struct {
	ZoneNumber      int
	Fares, DocFares []int
}

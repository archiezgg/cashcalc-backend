package models

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

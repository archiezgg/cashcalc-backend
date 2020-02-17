package models

// Pricing is the struct to store zone numbers and the matching fares
type Pricing struct {
	ZoneNumber int   `bson:"zoneNumber"`
	Fares      []Fare `bson:"fares"`
	DocFares   []Fare `bson:"docFares"`
}

//Pricings stores both air and road pricings as fields
type Pricings struct {
	AirPricings  []Pricing `bson:"airPricings"`
	RoadPricings []Pricing `bson:"roadPricings"`
}

// Fare stores the weight and the corresponding base fare of the package
type Fare struct {
	Weight   float64
	BaseFare int
}

package models

// Pricing is the struct to store zone numbers and the matching fares
type Pricing struct {
	ZoneNumber int    `bson:"zoneNumber" json:"zoneNumber"`
	Fares      []Fare `bson:"fares" json:"fares"`
	DocFares   []Fare `bson:"docFares" json:"docFares"`
}

//Pricings stores both air and road pricings as fields
type Pricings struct {
	AirPricings  []Pricing `bson:"airPricings" json:"airPricings"`
	RoadPricings []Pricing `bson:"roadPricings" json:"roadPricings"`
}

// Fare stores the weight and the corresponding base fare of the package
type Fare struct {
	Weight   float64 `bson:"weight" json:"weight"`
	BaseFare int     `bson:"baseFare" json:"baseFare"`
}

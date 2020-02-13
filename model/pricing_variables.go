package model

//PricingVariables is the struct to store the variables that can be set by the admin of the application
type PricingVariables struct {
	VATPercent          float64 `bson:"vatPercent"`
	AirFuelFarePercent  float64 `bson:"airFuelFarePercent"`
	RoadFuelFarePercent float64 `bson:"roadFuelFarePercent"`
	Express9h           int     `bson:"express9h"`
	Express9hHungarian  int     `bson:"express9hHun"`
	Express12h          int     `bson:"express12h"`
	Express12hHungarian int     `bson:"express12hHun"`
	InsuranceLimit      int     `bson:"insuranceLimit"`
	MinInsurance        int     `bson:"minInsurance"`
	EXT                 int     `bson:"ext"`
	RAS                 int     `bson:"ras"`
	TK                  int     `bson:"tk"`
}

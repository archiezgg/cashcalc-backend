/*
	Cashcalc 2020
	Copyright (C) 2019-2020 Istvan Nemeth
	mailto: nemethistvanius@gmail.com
*/

package models

//PricingVariables is the struct to store the variables that can be set by the admin of the application
type PricingVariables struct {
	VATPercent          float64 `bson:"vatPercent" json:"vatPercent"`
	AirFuelFarePercent  float64 `bson:"airFuelFarePercent" json:"airFuelFarePercent"`
	RoadFuelFarePercent float64 `bson:"roadFuelFarePercent" json:"roadFuelFarePercent"`
	Express9h           int     `bson:"express9h" json:"express9h"`
	Express9hHungarian  int     `bson:"express9hHun" json:"express9hHun"`
	Express12h          int     `bson:"express12h" json:"express12h"`
	Express12hHungarian int     `bson:"express12hHun" json:"express12hHun"`
	InsuranceLimit      int     `bson:"insuranceLimit" json:"insuranceLimit"`
	MinInsurance        int     `bson:"minInsurance" json:"minInsurance"`
	EXT                 int     `bson:"ext" json:"ext"`
	RAS                 int     `bson:"ras" json:"ras"`
	TK                  int     `bson:"tk" json:"tk"`
	EmergencyFare       int     `bson:"emergencyFare" json:"emergencyFare"`
}

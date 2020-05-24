/*
	Cashcalc 2020
	Copyright (C) 2019-2020 Istvan Nemeth
	mailto: nemethistvanius@gmail.com
*/

package models

// CalcInputData is the structure to which the inpu json will be parsed to
type CalcInputData struct {
	ZoneNumber      int     `json:"zoneNumber"`
	Weight          float64 `json:"weight"`
	Insurance       int     `json:"insurance"`
	DiscountPercent float64 `json:"discountPercent"`
	ExpressType     Express `json:"expressType"`
	IsDocument      bool    `json:"isDocument"`
	IsExt           bool    `json:"isExt"`
	IsRas           bool    `json:"isRas"`
	IsTk            bool    `json:"isTk"`
}

// CalcOutputData is the structure that contains the end result and its elements
type CalcOutputData struct {
	BaseFare      float64 `json:"baseFare"`
	ExpressFare   float64 `json:"expressFare"`
	InsuranceFare float64 `json:"insuranceFare"`
	ExtFare       float64 `json:"extFare"`
	RasFare       float64 `json:"rasFare"`
	TkFare        float64 `json:"tkFare"`
	FuelFare      float64 `json:"fuelFare"`
	Result        float64 `json:"result"`
}

// Express represents the 3 express types: worldwide, 9h, 12h
type Express string

const (
	// ExpressWorldwide represents express type worldwide
	ExpressWorldwide = "worldwide"
	// Express9h represents express type 9h
	Express9h = "9h"
	// Express12h represents express type 12h
	Express12h = "12h"
)

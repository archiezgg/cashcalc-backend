package properties

import (
	"github.com/magiconair/properties"
)

var (
	propertiesFile = "app.properties"
)

const (
	pricingsCollectionProp    = "collection.pricings"
	countriesCollectionProp   = "collection.countries"
	pricingVarsCollectionProp = "collection.pricingvars"
	pricingsEndpointProp      = "endpoint.pricings"
	countriesEndpointProp     = "endpoint.countries"
	pricingVarsEndpointProp   = "endpoint.pricingvars"
	airFaresZnMinProp         = "air.fares.zn.min"
	airFaresZnMaxProp         = "air.fares.zn.max"
	airDocFaresZnMinProp      = "air.docfares.zn.min"
	airDocFaresZnMaxProp      = "air.docfares.zn.max"
	roadFaresZnMinProp        = "road.fares.zn.min"
	roadFaresZnMaxProp        = "road.fares.zn.max"
	airFaresWeightMinProp     = "air.fares.weight.min"
	airFaresWeightMaxProp     = "air.fares.weight.max"
	airDocFaresWeightMinProp  = "air.docfares.weight.min"
	airDocFaresWeightMaxProp  = "air.docfares.weight.max"
	roadFaresWeightMinProp    = "road.fares.weight.min"
	roadFaresWeightMaxProp    = "road.fares.weight.max"
)

var (
	// PricingsCollection is the name of the DB collection of pricings
	PricingsCollection string
	// CountriesCollection is the name of the DB collection of countries
	CountriesCollection string
	// PricingVarsCollection is the name of the DB collection of pricing variables
	PricingVarsCollection string
	// PricingsEndpoint is the endpoint for pricings
	PricingsEndpoint string
	// CountriesEndpoint is the endpoint for countries
	CountriesEndpoint string
	// PricingVarsEndpoint is the endpoint for pricing variables
	PricingVarsEndpoint string
	// AirFaresZnMin is the minimum zone number for air fares
	AirFaresZnMin int
	// AirFaresZnMax is the maximum zone number for air fares
	AirFaresZnMax int
	// AirDocFaresZnMin is the minimum zone number for document fares
	AirDocFaresZnMin int
	// AirDocFaresZnMax is the maximum zone number for document fares
	AirDocFaresZnMax int
	// RoadFaresZnMin is the minimum zone number for road fares
	RoadFaresZnMin int
	// RoadFaresZnMax is the maxmimum zone number for road fares
	RoadFaresZnMax int
	// AirFaresWeightMin is the minimum weight for air fares
	AirFaresWeightMin float64
	// AirFaresWeightMax is the maxmimum weight for air fares
	AirFaresWeightMax float64
	// AirDocFaresWeightMin is the minimum weight for air document fares
	AirDocFaresWeightMin float64
	// AirDocFaresWeightMax is the maxmimum weight for air document fares
	AirDocFaresWeightMax float64
	// RoadFaresWeightMin is the minimmum weight for road fares
	RoadFaresWeightMin float64
	// RoadFaresWeightMax is the maxmimum weight for road fares
	RoadFaresWeightMax float64
)

// This function gets executed automatically when the app initializes
func init() {
	p := properties.MustLoadFile(propertiesFile, properties.UTF8)
	PricingsCollection = p.MustGetString(pricingsCollectionProp)
	CountriesCollection = p.MustGetString(countriesCollectionProp)
	PricingVarsCollection = p.MustGetString(pricingVarsCollectionProp)
	PricingsEndpoint = p.MustGetString(pricingsEndpointProp)
	CountriesEndpoint = p.MustGetString(countriesEndpointProp)
	PricingVarsEndpoint = p.MustGetString(pricingVarsEndpointProp)
	AirFaresZnMin = p.MustGetInt(airFaresZnMinProp)
	AirFaresZnMax = p.MustGetInt(airFaresZnMaxProp)
	AirDocFaresZnMin = p.MustGetInt(airDocFaresZnMinProp)
	AirDocFaresZnMax = p.MustGetInt(airDocFaresZnMaxProp)
	RoadFaresZnMin = p.MustGetInt(roadFaresZnMinProp)
	RoadFaresZnMax = p.MustGetInt(roadFaresZnMaxProp)
	AirFaresWeightMin = p.MustGetFloat64(airFaresWeightMinProp)
	AirFaresWeightMax = p.MustGetFloat64(airFaresWeightMaxProp)
	AirDocFaresWeightMin = p.MustGetFloat64(airDocFaresWeightMinProp)
	AirDocFaresWeightMax = p.MustGetFloat64(airDocFaresWeightMaxProp)
	RoadFaresWeightMin = p.MustGetFloat64(roadFaresWeightMinProp)
	RoadFaresWeightMax = p.MustGetFloat64(roadFaresWeightMaxProp)
}

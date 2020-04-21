package properties

import (
	"os"

	"github.com/magiconair/properties"
)

var (
	// Prop is the proprties that gets loaded when the app inits
	Prop           *properties.Properties
	propertiesFile = os.Getenv("PROPERTIES_FILE")
)

const (
	// PricingsCollection is the name of the DB collection of pricings
	PricingsCollection = "collection.pricings"
	// CountriesCollection is the name of the DB collection of countries
	CountriesCollection = "collection.countries"
	// PricingVarsCollection is the name of the DB collection of pricing variables
	PricingVarsCollection = "collection.pricingvars"
	// PricingsEndpoint is the endpoint for pricings
	PricingsEndpoint = "endpoint.pricings"
	// CountriesEndpoint is the endpoint for countries
	CountriesEndpoint    = "endpoint.countries"
	AirFaresZnMin        = "air.fares.zn.min"
	AirFaresZnMax        = "air.fares.zn.max"
	AirDocFaresZnMin     = "air.docfares.zn.min"
	AirDocFaresZnMax     = "air.docfares.zn.max"
	RoadFaresZnMin       = "road.fares.zn.min"
	RoadFaresZnMax       = "road.fares.zn.max"
	AirFaresWeightMin    = "air.fares.weight.min"
	AirFaresWeightMax    = "air.fares.weight.max"
	AirDocFaresWeightMin = "air.docfares.weight.min"
	AirDocFaresWeightMax = "air.docfares.weight.max"
	RoadFaresWeightMin   = "road.fares.weight.min"
	RoadFaresWeightMax   = "road.fares.weight.max"
)

// This function gets executed automatically when the app initializes
func init() {
	Prop = properties.MustLoadFile(propertiesFile, properties.UTF8)
}

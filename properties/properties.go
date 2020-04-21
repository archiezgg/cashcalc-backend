package properties

import (
	"os"

	"github.com/magiconair/properties"
)

var (
	// Prop is the proprties that gets loaded when the app inits
	Prop           *properties.Properties
	propertiesFile = os.Getenv("GOPATH") + "/src/github.com/IstvanN/cashcalc-backend/app.properties"
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
	CountriesEndpoint = "endpoint.countries"
)

// This function gets executed automatically on initiation
func init() {
	Prop = properties.MustLoadFile(propertiesFile, properties.UTF8)
}

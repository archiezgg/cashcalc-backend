package services

// IncreaseWithVat takes a float64 and a percentage as parameter
// and returns with the vat-increased result
func IncreaseWithVat(num float64, vat float64) float64 {
	return num * (1 + (vat / 100))
}

// IsZoneNumberInvalid tests if a given zone number is between given min and max
func IsZoneNumberInvalid(zn, min, max int) bool {
	return zn < min || zn > max
}

// IsWeightInvalid tests if a given weight is between the given min and max
func IsWeightInvalid(weight, min, max float64) bool {
	return weight < min || weight > max
}

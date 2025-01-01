// Package weather is a package for condition.
package weather

// CurrentCondition is a value.
var CurrentCondition string

// CurrentLocation is a value.
var CurrentLocation string

// Forecast is a function.
func Forecast(city, condition string) string {
	CurrentLocation, CurrentCondition = city, condition
	return CurrentLocation + " - current weather condition: " + CurrentCondition
}

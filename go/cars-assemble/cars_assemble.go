package cars

// CalculateWorkingCarsPerHour calculates how many working cars are
// produced by the assembly line every hour.
func CalculateWorkingCarsPerHour(productionRate int, successRate float64) float64 {
	successCars := float64(productionRate) * successRate
	return float64(successCars / 100)
}

// CalculateWorkingCarsPerMinute calculates how many working cars are
// produced by the assembly line every minute.
func CalculateWorkingCarsPerMinute(productionRate int, successRate float64) int {
	successCars := CalculateWorkingCarsPerHour(productionRate, successRate) / 60
	return int(successCars)
}

// CalculateCost works out the cost of producing the given number of cars.
func CalculateCost(carsCount int) uint {
	groupsCarCount := carsCount / 10
	singleCarCount := carsCount % 10

	cost := groupsCarCount*95000 + singleCarCount*10000
	return uint(cost)
}

package purchase

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {

	if kind == "car" {
		return true
	} else if kind == "truck" {
		return true
	} else {
		return false
	}
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in lexicographical order.
func ChooseVehicle(option1, option2 string) string {
	if option1 > option2 {
		return option2 + " is clearly the better choice."
	} else {
		return option1 + " is clearly the better choice."
	}

}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
	if age < 3 {
		originalPrice = originalPrice * 80 / 100
		return originalPrice
	} else if age > 2 && age < 10 {
		originalPrice = originalPrice * 70 / 100
		return originalPrice
	} else {
		originalPrice = originalPrice * 50 / 100
		return originalPrice
	}
}

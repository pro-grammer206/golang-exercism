package vehiclePurchase

import "fmt"

func NeedsLicense(kind string) bool {
	if kind == "car" || kind == "truck" {
		return true
	} else {
		return false
	}
}

func ChooseVehicle(option1, option2 string) string {
	selection := ""
	if option1 < option2 {
		selection = option1
	} else {
		selection = option2
	}
	return fmt.Sprintf("%s is clearly the better choice.", selection)
}
func CalculateResellPrice(originalPrice, age float64) float64 {
	var estimatedPrice float64
	if age < 3 {
		estimatedPrice = originalPrice * 0.8
	} else if age > 3 && age < 10 {
		estimatedPrice = originalPrice * 0.7
	} else {
		estimatedPrice = originalPrice * 0.5
	}
	return estimatedPrice
}
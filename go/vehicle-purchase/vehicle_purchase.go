package purchase

import (
	"fmt"
	"strings"
)

const (
	Car     = "car"
	Truck   = "truck"
	Heavy   = "heavy"
	Light   = "Light"
	Message = "%s is clearly the better choice."
)

var Vehicles = map[string]string{
	Car:   Heavy,
	Truck: Heavy,
}

// NeedsLicense determines whether a license is needed to drive a type of vehicle. Only "car" and "truck" require a license.
func NeedsLicense(kind string) bool {
	kind = strings.ToLower(kind)
	if _, ok := Vehicles[kind]; ok {
		return true
	}
	return false
}

// ChooseVehicle recommends a vehicle for selection. It always recommends the vehicle that comes first in dictionary order.
func ChooseVehicle(option1, option2 string) string {
	if strings.Compare(option1, option2) < 0 {
		return fmt.Sprintf(Message, option1)
	}

	return fmt.Sprintf(Message, option2)

}

// CalculateResellPrice calculates how much a vehicle can resell for at a certain age.
func CalculateResellPrice(originalPrice, age float64) float64 {
	switch {
	case age < 3:
		return .80 * originalPrice
	case age >= 10:
		return .50 * originalPrice
	default:
		return .70 * originalPrice
	}
}

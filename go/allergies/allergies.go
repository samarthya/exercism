package allergies

import "log"

// The problem just needs a BIT to be manipulated
const (
	eggs = 1 << iota
	peanuts
	shellfish
	strawberries
	tomatoes
	chocolate
	pollen
	cats
)

var allergies = map[string]uint{
	"eggs":         eggs,
	"peanuts":      peanuts,
	"shellfish":    shellfish,
	"strawberries": strawberries,
	"tomatoes":     tomatoes,
	"chocolate":    chocolate,
	"pollen":       pollen,
	"cats":         cats,
}

// Just for debug purpose
func printConsts() {
	for k, v := range allergies {
		log.Printf(" %s %b=%d \n", k, v, v)
	}
}

func init() {
	// printConsts()
}

// Allergies - Computes all allergies by looping through
func Allergies(a uint) (r []string) {
	for k, _ := range allergies {
		if AllergicTo(a, k) {
			r = append(r, k)
		}
	}
	return
}

// AllergicTo Checks the allergy by &'ing the bits to confirm
func AllergicTo(a uint, allergen string) bool {
	return a&allergies[allergen] > 0
}

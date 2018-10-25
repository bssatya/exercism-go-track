// Package space impliments method to calculate space age
package space

type Planet string

// Age function takes seconds and planet inputs, calcualtes the space age for the
// given planet and returns age
func Age(seconds float64, planet Planet) float64 {
	earthYear := seconds / 31557600.00
	switch planet {
	case "Earth":
		return earthYear
	case "Mercury":
		return earthYear / 0.2408467
	case "Venus":
		return earthYear / 0.61519726
	case "Mars":
		return earthYear / 1.8808158
	case "Jupiter":
		return earthYear / 11.862615
	case "Saturn":
		return earthYear / 29.447498
	case "Uranus":
		return earthYear / 84.016846
	case "Neptune":
		return earthYear / 164.79132
	default:
		return -1
	}
}

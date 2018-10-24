// Package twofer implements a method returning twofer string
package twofer

// ShareWith accepts name argument and returns a twofer for the same
func ShareWith(name string) string {
	if name == "" {
		return "One for you, one for me."
	} else {
		return "One for " + name + ", one for me."
	}
}

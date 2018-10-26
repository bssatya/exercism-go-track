// Package Triangle impliments functions to evaluate kind of triangle and return the kind
package triangle

import (
	"math"
)

// Notice KindFromSides() returns this type. Pick a suitable data type.
type Kind int

const (
	// Pick values for the following identifiers used by the test program.
	NaT = iota // not a triangle
	Equ = iota // equilateral
	Iso = iota // isosceles
	Sca = iota // scalene
)

// KindFromSides take 3 sides of a triangle as float64 and evaluates the kind of the triange, and returns kind (int).
func KindFromSides(a, b, c float64) Kind {
	var k Kind

	if checkTriangeInequality(a, b, c) || isNotFloat(a) || isNotFloat(b) || isNotFloat(c) {
		k = NaT
		return NaT
	}

	k = Sca
	if a == b || b == c || c == a {
		k = Iso
	}

	if a == b && a == c {
		k = Equ
	}

	return k
}

func isNotFloat(n float64) bool {
	return math.IsInf(n, 1) || math.IsInf(n, -1) || math.IsNaN(n) || n <= 0
}

func checkTriangeInequality(a, b, c float64) bool {
	if (a+b) < c || (b+c) < a || (c+a) < b {
		return true
	}
	return false
}

// Package hamming impliments the function to find hamming distace of 2 dna strands
package hamming

import (
	"errors"
)

// Distance accepts 2 strings and returns the differences between them, if the 2
// string differ is size, -1 will be returned
func Distance(a, b string) (int, error) {
	if len(a) != len(b) {
		return -1, errors.New("DNA strands with different length")
	}

	diff := 0
	for i := 0; i < len(a); i++ {
		if a[i] != b[i] {
			diff++
		}
	}
	return diff, nil
}

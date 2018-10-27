// Prackage proverb impliments menthod to create proverb with the words given
package proverb

import (
	"fmt"
)

const (
	proverbLines = "For want of a %s the %s was lost."
	endLine      = "And all for the want of a %s."
)

// Proverb function takes slice of stings and returns a proverb as slice of strings
func Proverb(rhyme []string) []string {
	rhymeLen := len(rhyme)
	if rhymeLen == 0 {
		return []string{}
	}

	proverb := make([]string, rhymeLen)

	for i := 0; i < rhymeLen-1; i++ {
		proverb[i] = fmt.Sprintf(proverbLines, rhyme[i], rhyme[i+1])
	}
	proverb[rhymeLen-1] = fmt.Sprintf(endLine, rhyme[0])

	return proverb
}

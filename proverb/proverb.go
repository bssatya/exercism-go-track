// This is a "stub" file.  It's a little start on your solution.
// It's not a complete solution though; you have to write some code.

// Package proverb should have a package comment that summarizes what it's about.
// https://golang.org/doc/effective_go.html#commentary
package proverb

import (
	"fmt"
)

// Proverb should have a comment documenting it.
func Proverb(rhyme []string) []string {
	proverb := []string{}
	var temp string
	// Write some code here to pass the test suite.
	// Then remove all the stock comments.
	// They're here to help you get started but they only clutter a finished solution.
	// If you leave them in, reviewers may protest!
	for i := 0; i < len(rhyme) - 1; i++ {
		temp = ""
		fmt.Sprintln(temp, "For want of " + rhyme[i] + "the" + rhyme[i+1] + "was lost")
		proverb = append(proverb, temp)
	}
	fmt.Sprintln(temp, "And all for the want of a" + rhyme[0])
	proverb = append(proverb, temp)

	return proverb
}

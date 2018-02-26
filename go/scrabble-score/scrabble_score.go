// Package scrabble does stuff
package scrabble

import (
	"strings"
	"unicode"
)

// Score returns a scrabble score.
func Score(word string) int {
	counts := make(map[rune]int)
	for _, r := range word {
		counts[unicode.ToUpper(r)] += 1
	}

	value := 0
	for r, count := range counts {
		switch {
		case strings.ContainsRune("AEIOULNRST", r):
			value += 1 * count
		case strings.ContainsRune("DG", r):
			value += 2 * count
		case strings.ContainsRune("BCMP", r):
			value += 3 * count
		case strings.ContainsRune("FHVWY", r):
			value += 4 * count
		case strings.ContainsRune("K", r):
			value += 5 * count
		case strings.ContainsRune("JX", r):
			value += 8 * count
		case strings.ContainsRune("QZ", r):
			value += 10 * count
		}
	}
	return value
}

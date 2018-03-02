// Package isogram is cool.
package isogram

import "unicode"

// IsIsogram checks if input is an isogram.
func IsIsogram(input string) bool {
	chars := make(map[rune]bool)

	for _, r := range input {
		if !unicode.IsLetter(r) {
			continue
		}
		r = unicode.ToLower(r)
		if _, exists := chars[r]; exists {
			return false
		} else {
			chars[r] = true
		}
	}

	return true
}

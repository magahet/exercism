// Package isogram is cool.
package isogram

import "unicode"

// IsIsogram checks if input is an isogram.
func IsIsogramNested(input string) bool {

	for i := range input {
		for j := range input {
			if i == j || !unicode.IsLetter(rune(input[i])) || !unicode.IsLetter(rune(input[j])) {
				continue
			}
			if unicode.ToLower(rune(input[i])) == unicode.ToLower(rune(input[j])) {
				return false
			}
		}
	}

	return true
}

func IsIsogramSet(input string) bool {
	chars := make(map[rune]bool)

	for _, r := range input {
		if !unicode.IsLetter(r) {
			continue
		}
		r = unicode.ToLower(r)
		if _, exists := chars[r]; exists {
			return false
		}

		chars[r] = true
	}

	return true
}

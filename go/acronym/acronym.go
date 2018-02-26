// Package acronym helps make strings short.
package acronym

import (
	"unicode"
)

// Abbreviate makes strings short.
func Abbreviate(s string) string {
	lastNotLetter := true
	output := ""
	for _, r := range s {
		if lastNotLetter && unicode.IsLetter(r) {
			output += string(unicode.ToUpper(r))
			lastNotLetter = false
		} else if !unicode.IsLetter(r) {
			lastNotLetter = true
		}
	}
	return output
}

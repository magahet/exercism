// Package reverse
package reverse

// String function
func String(input string) string {
	chars := []rune(input)
	for i := len(chars)/2 - 1; i >= 0; i-- {
		chars[i], chars[len(chars)-1-i] = chars[len(chars)-1-i], chars[i]
	}
	return string(chars)
}

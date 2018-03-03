// Package luhn validates numbers.
package luhn

import (
	"strconv"
	"unicode"
)

// Valid checks if number is valid.
func Valid(input string) bool {
	var numbers []int
	for _, r := range input {
		switch {
		case unicode.IsSpace(r):
			continue
		case unicode.IsNumber(r):
			n, _ := strconv.Atoi(string(r))
			numbers = append(numbers, n)
		default:
			return false
		}
	}

	if len(numbers) < 2 {
		return false
	}

	var sum, n int
	for i, _ := range numbers {
		n = numbers[len(numbers)-1-i]
		if i%2 == 1 {
			n = n * 2
			if n > 9 {
				n -= 9
			}
		}
		sum += n
	}
	return sum%10 == 0
}

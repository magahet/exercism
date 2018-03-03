// Package diffsquares provides funcs to calc sum/squares.
package diffsquares

// Pow does simple integer exponentiation.
func Pow(x, y int) int {
	z := 1
	for i := 0; i < y; i++ {
		z *= x
	}
	return z
}

// SquareOfSums calcs the sum of squares.
func SquareOfSums(n int) int {
	return Pow(n*(n+1)/2, 2)
}

// SumOfSquares calcs the sum of squares.
func SumOfSquares(n int) int {
	return n * (n + 1) * (2*n + 1) / 6
}

// Difference calcs the diff between the two.
func Difference(n int) int {
	return SquareOfSums(n) - SumOfSquares(n)
}

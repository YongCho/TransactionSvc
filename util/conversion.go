package util

import "math"

// DollarToCents converts dollar amount in float to cents amount in integer.
// It rounds the input to the 2 decimal digits (1 cent) to handle the float precision issue.
//
// Example:
//
//	59.99 -> 5999
//	0.99999994 -> 100
func DollarToCents(dollar float64) int64 {
	return int64(math.Round(dollar * 100))
}

// CentsToDollar converts cent amount (int) to dollar amount (float).
func CentsToDollar(cents int64) float64 {
	return float64(cents) / 100
}

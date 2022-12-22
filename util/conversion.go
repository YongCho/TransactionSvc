// This file contains the logic for converting dollars to cents and vice versa.
// We use cents (int) instead of dollars (float) for storing the money amount to avoid the
// float precision issues.

package util

import "math"

// DollarToCents converts dollar (float) to cent (int).
// It rounds the input to the 2 decimal digits (1 cent) to handle the float precision issue.
//
// Example:
//
//	0.99999994 -> 100
func DollarToCents(dollar float64) int64 {
	return int64(math.Round(dollar * 100))
}

// CentsToDollar converts cent (int) to dollar (float).
func CentsToDollar(cents int64) float64 {
	return float64(cents) / 100
}

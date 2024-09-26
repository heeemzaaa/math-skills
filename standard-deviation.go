package ms

import (
	"math"
)

// This function calculates the standart deviation
// The SD is the square root of the variance
func Standard_Deviation(variance float64) float64 {
	deviation := math.Sqrt(variance)
	return deviation
}

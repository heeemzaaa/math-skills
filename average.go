package ms

import "math"

// This function return the average of the data
// The average is the sum of all the numbers there divided by the number of them (length)
func Average(data []float64) float64 {
	var sum float64
	length := len(data)
	var result float64

	for i := 0; i < length; i++ {
		sum += float64(data[i])
	}
	result = sum / float64(length)
	return float64(int(math.Round(float64(result))))
}

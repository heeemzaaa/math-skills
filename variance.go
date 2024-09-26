package ms

import "math"

// This function calculates the variance
// To calculate the variance we have to get through some operations
// First of all we calculate the average
// Then we substract the average from each number in the data
// After that we get the squared of each number of the data
// Then we sum all the squared number in the data
// Finally we divide the result by the length of the data
func Variance(data []float64) float64 {
	var mean float64
	var variance float64
	length := len(data)
	var sum float64

	for i := 0; i < length; i++ {
		sum += data[i]
	}
	mean = sum / float64(length)
	for i := 0; i < length; i++ {
		data[i] = data[i] - mean
	}
	for i := 0; i < length; i++ {
		data[i] = data[i] * data[i]
	}
	sum = 0
	for i := 0; i < length; i++ {
		sum += data[i]
	}
	variance = sum / float64(length)
	return float64(int(math.Round(variance)))
}

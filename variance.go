package ms

import "math"

func Variance(data []float64) int {
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
	return int(math.Round(variance))
}

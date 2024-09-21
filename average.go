package ms

import "math"

func Average(data []float64) int {
	var sum float64
	length := len(data)
	var result float64

	for i := 0; i < length; i++ {
		sum += float64(data[i])
	}
	result = sum / float64(length)
	return int(math.Round(float64(result)))
}

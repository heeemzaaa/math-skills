package ms

// This function return the average of the data
// The average is the sum of all the numbers there divided by the number of them (length)
func Average(data []float64) float64 {
	var sum float64
	length := len(data)
	var average float64

	for i := 0; i < length; i++ {
		sum += float64(data[i])
	}
	average = sum / float64(length)
	return average
}

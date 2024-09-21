package ms

import "math"

//this function is for sorting the data's file in ascending order
// I used the bublle sorting
func SortData(data []float64) []float64 {
	for i := 0; i < len(data); i++ {
		for j := 0; j < len(data)-1-i; j++ {
			if data[j] > data[j+1] {
				data[j], data[j+1] = data[j+1], data[j]
			}
		}
	}
	return data
}

// This function finds the midean number in the data file
// The median number is the middle number in the data after we sorted it
func Median(data []float64) int {
	sortedData := SortData(data)
	var median float64
	if len(sortedData)%2 != 0 {
		median = sortedData[len(sortedData)/2]
	} else {
		median = (sortedData[(len(sortedData)/2)-1] + sortedData[(len(sortedData)/2)]) / 2
	}
	return int(math.Round(median))
}

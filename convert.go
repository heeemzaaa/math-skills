package ms

import (
	"log"
	"strconv"
)

// This function in converting the data passed from the file , it was in format string
// it converted to float64 to calculate also the cases of .
func ConvertingToFLoat(slice []string) []float64 {
	result := []float64{}
	for i := 0; i < len(slice); i++ {
		holder, err := strconv.ParseFloat(slice[i], 64)
		if err != nil {
			log.Println("Error: ", err)
			return nil
		}
		result = append(result, holder)
		holder = 0
	}
	return result
}

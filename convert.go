package ms

import (
	"log"
	"strconv"
)

func ConvertingToInt(slice []string) []float64 {
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

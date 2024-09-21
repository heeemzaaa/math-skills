package ms

import (
	"math"
)

func Standard_Deviation(variance float64) int {
	deviation := math.Sqrt(variance)
	return int(math.Round(deviation))
}

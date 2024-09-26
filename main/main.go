package main

import (
	"fmt"
	"log"
	"os"

	"ms"
)

func main() {
	if len(os.Args) != 2 {
		log.Println("Argument Error !")
		return
	}
	data := ms.ReadFile(os.Args[1])
	if data == nil {
		log.Println("Error reading data !")
		return
	}
	splitted_data := ms.ConvertingToFLoat(data)
	average := ms.Average(splitted_data)
	median := ms.Median(splitted_data)
	variance := ms.Variance(splitted_data)
	deviation := ms.Standard_Deviation(variance)

	fmt.Printf("Average: %.0f\n", average)
	fmt.Printf("Median: %.0f\n", median)
	fmt.Printf("Variance: %.0f\n", variance)
	fmt.Printf("Standard Deviation: %.0f\n", deviation)
}

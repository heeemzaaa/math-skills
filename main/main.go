package main

import (
	"fmt"
	"log"
	"ms"
	"os"
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
	deviation := ms.Standard_Deviation(float64(variance))
	fmt.Println(variance)
	fmt.Println(float64(variance))

	fmt.Printf("Average: %d\n", average)
	fmt.Printf("Median: %d\n", median)
	fmt.Printf("Variance: %d\n", variance)
	fmt.Printf("Standard Deviation: %d\n", deviation)
}

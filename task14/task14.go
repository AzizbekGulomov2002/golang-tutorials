package main

import (
	"fmt"
	"math"
)

func main() {
	var s float64
	fmt.Print("Enter surface = ")
	fmt.Scan(&s)
	diameter := 2 * math.Sqrt(s/3.14)
	fmt.Printf("Diameter = %.2f\n", diameter) 
	length := math.Pi * diameter
	fmt.Printf("Length = %.2f\n", length)
}

package main

import (
	"fmt"
)

func main() {
	var x, y int
	fmt.Print("Enter x = ")
	fmt.Scanln(&x)
	fmt.Print("Enter y = ")
	fmt.Scanln(&y)

	if (x+y)%2 != 0 {
		fmt.Println(true) // Black
	} else {
		fmt.Println(false) // White
	}
}

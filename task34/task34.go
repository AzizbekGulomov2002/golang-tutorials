package main

import (
	"fmt"
)

func main() {
	var x1, y1, x2, y2 int
	fmt.Print("Enter x1, y1 = ")
	fmt.Scanln(&x1, &y1)
	fmt.Print("Enter x2, y2 = ")
	fmt.Scanln(&x2, &y2)

	if (x1+y1)%2 == (x2+y2)%2 {
		fmt.Println(true) // The same
	} else {
		fmt.Println(false) // Seperate
	}
}

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

	if (x > 0 && y > 0) || (x < 0 && y < 0) {
		fmt.Println(true)
	} else {
		fmt.Println(false)
	}
}

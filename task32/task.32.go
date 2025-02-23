package main

import (
	"fmt"
)

func main() {
	var a, b, c int
	fmt.Print("Enter a = ")
	fmt.Scanln(&a)
	fmt.Print("Enter b = ")
	fmt.Scanln(&b)
	fmt.Print("Enter c = ")
	fmt.Scanln(&c)

	if a+b > c && a+c > b && b+c > a {
		fmt.Println(true)
	} else {
		fmt.Println(false)
	}
}

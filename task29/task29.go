package main

import (
	"fmt"
)

func condition(a, b, c int) bool {
	if a+b > c && a+c > b && b+c > a {
		return a == b && b == c
	}
	return false
}

func main() {
	var a, b, c int
	fmt.Print("Enter a = ")
	fmt.Scanln(&a)
	fmt.Print("Enter b = ")
	fmt.Scanln(&b)
	fmt.Print("Enter c = ")
	fmt.Scanln(&c)

	fmt.Println(condition(a, b, c))
}

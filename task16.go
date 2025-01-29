// Given a,b,c name points. Find a*b+a*c+b*c
package main

import "fmt"
func main()  {
	var a,b,c float64
	fmt.Print("Enter A = ")
	fmt.Scan(&a)

	fmt.Print("Enter B = ")
	fmt.Scan(&b)

	fmt.Print("Enter c = ")
	fmt.Scan(&c)

	sum := (a*b+a*c+b*c)
	fmt.Println("Summ of these numbers = ", sum)
}
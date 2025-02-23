package main

import (
	"fmt"
)
func main()  {
	var x int
	fmt.Print("Enter x = ")
	fmt.Scanln(&x)

	if x<1000 || x>9999{
		fmt.Println(false)
		return
	}

	d1 := x / 1000
	d2 := (x/100)%10
	d3 := (x/10)%10
	d4 := x%10

	if d1 == d4 && d2==d3{
		fmt.Println(true)
	}else{
		fmt.Println(false)
	}
}
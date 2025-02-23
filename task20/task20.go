package main

import (
	"fmt"
)
func main()  {
	var x,y,z int
	fmt.Print("Enter x = ")
	fmt.Scanln(&x)
	fmt.Print("Enter y = ")
	fmt.Scanln(&y)
	fmt.Print("Enter z = ")
	fmt.Scanln(&z)

	if x==y || x==z || z==y{
		fmt.Println(false)
	}else{
		fmt.Println(true)
	}
}
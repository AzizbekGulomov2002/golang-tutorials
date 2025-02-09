// Given three-digit number. Replace it's unit and decimal places
package main
import "fmt"
func main()  {
	var num int // 324 -> 342
	fmt.Print("Enter three-digit number = ")
	fmt.Scanln(&num)
	if num < 100 || num > 999{
		fmt.Println("Please enter only tree-digit number")
		return
	}

	a := (num / 100) * 100
	b := ((num % 10) % 10) * 10
	c := ((num / 10) % 10)

	fmt.Println(a+b+c)
}
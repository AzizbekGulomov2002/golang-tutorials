// Given three digit number. Find it's decimal and unit part of it
package main
import "fmt"
func main()  {
	var num int
	fmt.Print("Enter three digit number = ")
	fmt.Scanln(&num)

	if num < 100 || num > 999{
		fmt.Println("Please enter only three digit")
		return
	}

	decimal := (num / 10 ) % 10
	unit  := num % 10

	fmt.Println(decimal, unit)

}
// Given three digit number. Find the sum of its numbers
package main
import "fmt"
func main()  {
	var num int
	fmt.Print("Enter three-digit number = ")
	fmt.Scanln(&num)

	if num < 100 || num > 999{
		fmt.Println("Please enter only three-digit number")
		return
	}

	hundred := num / 100
	decimal := (num / 10) % 10
	unit := (num % 10) % 10

	fmt.Println(hundred+decimal+unit)

}
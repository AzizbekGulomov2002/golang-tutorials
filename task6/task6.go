// Given decimal number. Find multiplication and sum of it's unit and decimal parts
package main
import "fmt"
func main()  {
	var num int
	fmt.Print("Enter decimal number = ")
	fmt.Scanln(&num)

	decimal := num / 10
	unit := num % 10

	sum := decimal + unit
	multiplied := decimal * unit
	
	fmt.Println("Sum of units = ", sum)
	fmt.Println("Multiplication of units = ", multiplied)
}
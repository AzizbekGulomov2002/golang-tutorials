// Given decimal number. Return its decimal and unit parts
package main
import "fmt"
func main()  {
	var num int
	fmt.Print("Enter decimal number = ")
	fmt.Scanln(&num)

	decimal := num / 10
	unit := num % 10

	fmt.Println("Decimal part:", decimal, "Unit part:", unit)
}b
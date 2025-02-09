// Give that not equal two numbers. Find its sum, difference, multiplication, division

package main
import "fmt"
var a,b uint
func main()  {
	fmt.Print("Enter a = ")
	fmt.Scan(&a)

	fmt.Print("Enter b = ")
	fmt.Scan(&b)

	sum := a+b
	difference := a-b
	multiplication := a*b
	division := a/b

	fmt.Println("Sum = ", sum, "Difference = ", difference, "Multiplication = ", multiplication,"Division = ", division)
}
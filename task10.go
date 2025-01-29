// Given 2 numbers that not equal 0. Find their difference's module result, sum, multiplication, division

package main
import (
	"fmt"
	"math"
)
func main()  {
	var a,b float64
	fmt.Print("Enter a  = ")
	fmt.Scan(&a)
	for a==0{
		fmt.Print("a can't be 0. Enter validate number = ")
		fmt.Scan(&a)
	}

	fmt.Print("Enter b = ")
	fmt.Scan(&b)
	for b == 0 {
		fmt.Print("b can't be 0. Enter validate number = ")
		fmt.Scan(&b)
	}

	sum := a+b
	multiplication := a*b
	division := math.Abs(math.Floor(a/b))
	difference := a-b

	fmt.Println("Sum = ", sum, "Multiplication = ", multiplication, "Division = ", division, "Difference = ", difference)
}

// Given x number. Find second degree, third and fourth degree
package main
import "fmt"
func main()  {
	var x float64
	fmt.Print("Enter x = ")
	fmt.Scan(&x)

	second := x*x
	third := x*x*x
	fourth := x*x*x*x

	fmt.Println("Second degree = ", second)
	fmt.Println("Third degree = ", third)
	fmt.Println("Fourth degree = ", fourth)
}
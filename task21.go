// Given x. Find y = 36x-16x-7
package main
import "fmt"
func main()  {
	var x float64
	fmt.Print("Enter x = ")
	fmt.Scan(&x)

	y := 36*x-16*x-7
	fmt.Println("Result y = ", y)
}
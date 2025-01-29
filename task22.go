// Given x. Find y = 4(x-3)*6 - 7*(x-3)*3-2
package main
import "fmt"
func main()  {
	var x float64
	fmt.Print("Enter x = ")
	fmt.Scan(&x)

	y := 4*(x-3)*6 - 7*(x-3)*3-2
	fmt.Println("Value y = ", y)
}
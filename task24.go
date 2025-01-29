// Given x. Find degree2, degree2, degree5, degree10, degree15
package main
import "fmt"
func main()  {
	var a float64
	fmt.Print("Enter a = ")
	fmt.Scan(&a)

	a2 := a*a
	a3 := a2*a
	a5 := a3*a2
	a10 := a5*a5
	a15 := a10*a5

	fmt.Println("A15 = ", a15)
}
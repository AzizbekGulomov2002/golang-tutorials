// Given a and b sides of the right triangle. Find it's c side and Perimeter
package main
import (
	"fmt"
	"math"
)
func main()  {
	var a,b int
	fmt.Print("Enter a = ")
	fmt.Scan(&a)

	fmt.Print("Enter b = ")
	fmt.Scan(&b)

	hypotenuse := math.Sqrt(float64(a*a + b*b))
	perimeter := a+b+int(hypotenuse)
	fmt.Println("Hypotenuse = ", hypotenuse)
	fmt.Println("Perimeter = ",perimeter)
}
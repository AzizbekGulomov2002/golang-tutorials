// Given a and b sides of the rectangle. Find the surface and perimeter
package main
import "fmt"
func main()  {
	var (
		a int
		b int
	)
	fmt.Print("Enter the a point of rectangle = ")
	fmt.Scan(&a)
	fmt.Print("Enter the b point of rectangle = ")
	fmt.Scan(&b)

	p := 2*(a+b)
	s := a*b
	fmt.Println("The perimeter of the rectangle is = ", p)
	fmt.Println("The surface of the rectangle is = ", s)

}
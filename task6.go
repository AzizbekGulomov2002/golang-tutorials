// Given the Radius of the sphere. Find the Length and Surface of Circle
package main
import "fmt"
func main()  {
	var r float64

	fmt.Print("Enter the radius of sphere = ")
	fmt.Scan(&r)

	l := 2*3.14*r
	s := r*r*3.14

	fmt.Println("Length of the sphere = ", l)
	fmt.Println("Surface of the sphere = ", s)
}
// Given length of the circle named L. Find Radius and Surface
package main
import "fmt"
func main()  {
	var l float64
	fmt.Print("Enter length = ")
	fmt.Scan(&l)

	radius := (2*3.14)/l
	surface := radius*radius*3.14
	fmt.Println("Radius = ",radius)
	fmt.Println("Surface = ", surface)

}
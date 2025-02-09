// Given 2 spheres radiuses. Find their surfaces and surface differences
package main
import "fmt"
func main()  {
	var r1, r2 float64
	fmt.Print("Enter Radius 1 = ")
	fmt.Scan(&r1)

	fmt.Print("Enter Radius 2 = ")
	fmt.Scan(&r2)

	surface1 := 3.14*r1*r1
	surface2 := 3.14*r2*r2
	difference := surface1-surface2

	fmt.Println("Surface 1 = ", surface1, "Surface 2 = ", surface2, "Difference = ", difference)
}
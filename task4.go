// Given a side of the cubes named "a". Find volume and surface of the cube
package main
import "fmt"
func main()  {
	var a int
	fmt.Print("Enter the side of the cube = ")
	fmt.Scanln(&a)

	v := a*a*a
	s := (a*a)*6

	fmt.Println("Cube's volume is = ", v)
	fmt.Println("Cube's surface is = ", s)

}
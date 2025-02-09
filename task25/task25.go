// Given Tempirature named "t". Find degree of Celsius. Tc = ((t-32)*5) / 9

package main
import "fmt"
func main()  {
	var t int
	fmt.Print("Enter tempirature = ")
	fmt.Scan(&t)

	tc := (5*(t-32))/9
	println("Tempirature Celsisus degree = ", tc)
}
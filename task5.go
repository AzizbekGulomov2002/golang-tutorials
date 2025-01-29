// Given Rectangle shaped paralellogram's edges "a, b, c". Find Volume, Surface
package main
import "fmt"
func main()  {
	var (
		a float64
		b float64
		c float64
	)
	fmt.Print("Enter side a = ")
	fmt.Scan(&a)

	fmt.Print("Enter side b = ")
	fmt.Scan(&b)

	fmt.Print("Enter side c = ")
	fmt.Scan(&c)

	v := a*b*c
	s := 2*(a*b+a*c+b*c)

	fmt.Println("Volume is = ", v)
	fmt.Println("Surface is = ", s)
}
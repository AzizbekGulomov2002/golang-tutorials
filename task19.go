// Given 2 numbers. Result should be change it's places
package main
import "fmt"

func main()  {
	var a,b float64
	fmt.Print("Enter a = ")
	fmt.Scan(&a)

	fmt.Print("Enter b = ")
	fmt.Scan(&b)

	a,b = b,a
	fmt.Println("Changed.", "a =", a, "b =",b)
}
// Give a and b named numbers. Find normal arithmetic value
package main
import "fmt"
func main()  {
	var (
		a float64
		b float64
	)

	fmt.Print("Enter side a = ")
	fmt.Scan(&a)

	fmt.Print("Enter side b = ")
	fmt.Scan(&b)

	arithmetic := (a+b) / 2
	fmt.Println("The normal arithmetic value is = ", arithmetic)
}
// Given three numbers. Result should be descending order
package main
import "fmt"
func main()  {
	var a,b,c float64
	fmt.Print("Enter a = ")
	fmt.Scan(&a)

	fmt.Print("Enter b = ")
	fmt.Scan(&b)

	fmt.Print("Enter c = ")
	fmt.Scan(&c)

	// Bubble sort type
	if a<b{
		a,b=b,a
	}
	if a<c{
		a,c = c,a
	}
	if b<c{
		b,c = c,b
	}

	fmt.Println("Descending format = ", a,b,c)
}

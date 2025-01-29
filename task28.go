// Given v1,v2 and t. Find S difference
package main
import "fmt"

func main()  {
	var v1,v2,t float64
	fmt.Print("Enter v1 = ")
	fmt.Scan(&v1)

	fmt.Print("Enter v2 = ")
	fmt.Scan(&v2)

	fmt.Print("Enter t = ")
	fmt.Scan(&t)

	s := t*(v2-v1)
	fmt.Println("S = ", s)
}
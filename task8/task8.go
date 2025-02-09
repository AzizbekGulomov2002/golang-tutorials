// Given unsigned a and b numbers. Find normal geometric value
package main
import (
	"fmt"
	"math"
)

func main()  {
	var a,b uint
	fmt.Print("Enter side a = ")
	fmt.Scan(&a)

	fmt.Print("Enter side b = ")
	fmt.Scan(&b)

	geometric := math.Sqrt(float64(a * b))

	fmt.Println("Normal geometric = ",geometric)
}

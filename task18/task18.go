// Given between of two coordinate points named (x1,x2) and (y1,y2)
package main
import (
	"fmt"
	"math"
)

func main()  {
	var x1,x2,y1,y2 float64
	fmt.Print("Enter x1 = ")
	fmt.Scan(&x1)

	fmt.Print("Enter x2 = ")
	fmt.Scan(&x2)

	fmt.Print("Enter y1 = ")
	fmt.Scan(&y1)

	fmt.Print("Enter y2 = ")
	fmt.Scan(&y2)
	a := math.Sqrt(math.Pow(x2-x1, 2) + math.Pow(y2-y1, 2))

	fmt.Println("Result of the value = ", a)
}
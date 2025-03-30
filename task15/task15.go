package main
import (
	"fmt"
	"math"
)

func main()  {
	var x int
	fmt.Print("Enter x = ")
	fmt.Scanln(&x)

	sqrt := int(math.Sqrt(float64(x)))
	switch sqrt*sqrt{
	case x:
		fmt.Println("Perfect square = ", sqrt)
	default:
		fmt.Println("Not a perfect square")
	}
}
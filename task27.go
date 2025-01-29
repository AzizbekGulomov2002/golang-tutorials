// Give x kg choco with A sum. y kg choco2 with B sum. Find prices and differences between with
package main
import "fmt"
func main()  {
	var x,A,y,B float64
	fmt.Print("x = ")
	fmt.Scan(&x)

	fmt.Print("sum of ", x, "= ")
	fmt.Scan(&A)

	fmt.Print("y = ")
	fmt.Scan(&y)

	fmt.Print("sum of ", y, "= ")
	fmt.Scan(&B)

	sumX := x*A
	sumY := y*B
	difference := sumX-sumY
	fmt.Println("Sum of", x, " = ", sumX)
	fmt.Println("Sum of", y, " = ", sumY)
	fmt.Println("Difference = ", difference)
}
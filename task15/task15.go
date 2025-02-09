// Given coordinate numbers with x1 and x2. Find distance their difference between
package main
import "fmt"
func main()  {
	var x1, x2 float64
	fmt.Print("Enter x1 = ")
	fmt.Scan(&x1)

	fmt.Print("Enter x2 = ")
	fmt.Scan(&x2)

	difference := (x2-x1)
	fmt.Println("Difference = ", difference)
}
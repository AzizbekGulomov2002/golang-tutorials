// Given 2 length of cross lines. Divide each other
package main
import "fmt"
func main()  {
	var a float64
	var b float64
	fmt.Print("Enter a = ")
	fmt.Scanln(&a)
	fmt.Print("Enter b = ")
	fmt.Scanln(&b)

	result := a/b
	fmt.Println("Result = ", result)
}
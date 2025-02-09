// Given n second, Generate it minute
package main
import "fmt"
func main()  {
	var num int
	fmt.Print("Enter second = ")
	fmt.Scanln(&num)

	fmt.Println(num/60)
}
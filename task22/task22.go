// Given n seconds, generate it's hours and remaining seconds
package main
import "fmt"
func main()  {
	var num uint
	fmt.Print("Enter number = ")
	fmt.Scanln(&num)

	fmt.Println(num%100)
}

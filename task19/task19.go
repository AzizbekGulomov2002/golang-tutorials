// Given n second, generate it hours
package main
import "fmt"
func main()  {
	var num int
	fmt.Print("Enter second = ")
	fmt.Scanln(&num)

	fmt.Println(num / 3600)
}
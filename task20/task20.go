// Given n seconds, generate it minutes and find it's remain
package main
import "fmt"
func main()  {
	var num uint
	fmt.Print("Enter seconds = ")
	fmt.Scanln(&num)

	remaining := num % 60

	fmt.Println(remaining)
}
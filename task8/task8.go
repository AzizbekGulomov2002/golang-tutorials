// Given three digit number. Find it's hundred unit
package main
import "fmt"
func main()  {
	var num int
	fmt.Print("Enter three digit num = ")
	fmt.Scanln(&num)
	if num <100 || num > 999 {
		fmt.Println("Enter only three digit number")
		return
	}
	
	result := num / 100
	fmt.Println(result)

}
// Given four-digit number. Find it's thousand's place
package main
import "fmt"
func main()  {
	var num int // 1234 -> 1
	fmt.Print("Enter four-digit number = ")
	fmt.Scanln(&num)
	if num < 1000 || num > 9999{
		fmt.Println("Please enter four-digit number")
		return
	}
	a := (num / 1000)
	fmt.Println(a)
}
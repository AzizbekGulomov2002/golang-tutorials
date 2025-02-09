// Given four-digit number. Find it's hundred's point
package main
import "fmt"
func main()  {
	var num int // 1234 -> 2
	fmt.Print("Enter four-digit number = ")
	fmt.Scanln(&num)

	if num < 1000 || num > 9999{
		fmt.Println("Please enter four-digit number")
		return
	}

	a := (num / 100)%10
	fmt.Println(a)

}
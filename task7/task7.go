// Given decimal number. Find it's replace digits 
package main
import "fmt"
func main()  {
	var num int
	fmt.Print("Enter decimal number = ")
	fmt.Scanln(&num)

	if num < 10 || num > 99{
		fmt.Println("Enter only two digit number")
		return
	}

	decimal := num/10
	unit := num%10

	result := unit*10+decimal

	fmt.Println(result)
}
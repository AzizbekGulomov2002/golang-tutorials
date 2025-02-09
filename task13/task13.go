// Given three-digit number. Replace it's first and last units
package main
import "fmt"
func main()  {
	var num int
	fmt.Print("Enter number = ")
	fmt.Scanln(&num)

	if num < 100 || num > 999{
		fmt.Println("Please enter three-digit number")
		return
	}

	a := ((num % 10)%10) *100
	b := ((num/10)%10) * 10
	c := (num/100)


	fmt.Println(a+b+c)
}
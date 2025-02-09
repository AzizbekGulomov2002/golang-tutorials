// Given three-digit number. Replace it's hundred and decimal points
package main
import "fmt"
func main()  {
	var num int // 324 -> 234
	fmt.Print("Enter three-digit number = ")
	fmt.Scanln(&num)
	if num < 100 || num > 999{
		fmt.Println("Please enter only three-digit num")
		return
	}

	a := ((num/10)%10) * 100
	b := ((num/100)) * 10
	c := ((num%10)%10)

	fmt.Println(a+b+c)

}
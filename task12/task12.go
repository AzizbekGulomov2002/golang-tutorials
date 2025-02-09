// Given three-digit number. Replace it's second and first units
package main
import "fmt"
func main()  {
	var num int // 324 = 432
	fmt.Print("Enter number  = ")
	fmt.Scanln(&num)

	if num <100 || num > 999{
		fmt.Println("Please enter only three-digit number")
		return
	}
	a := ((num%100)%10) * 100
	b := (num/100) * 10
	c := ((num / 10) % 10) 

	fmt.Println(a+b+c)

}
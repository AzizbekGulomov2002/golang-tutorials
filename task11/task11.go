// Given three-digit number. Replace it 
package main
import "fmt"
func main()  {
	var num int
	fmt.Print("Enter three-digit num = ")
	fmt.Scanln(&num)
	if num < 100 || num > 999{
		fmt.Println("Please enter only three-digit number ")
		return
	}

	a := ((num%10)%10)*100
	b := ((num/10)%10)*10
	c := (num/100)

	fmt.Println(a+b+c)

}
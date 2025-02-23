package main
import "fmt"
func main()  {
	var x int
	fmt.Print("Enter x = ")
	fmt.Scanln(&x)

	if x<100 || x>999{
		fmt.Println(false)
		return
	}
	digit1 := x/100
	digit2 := (x/10)%10
	digit3 := x%10

	if digit1 < digit2 && digit2 < digit3{
		fmt.Println(true)
	} else{
		fmt.Println(false)
	}
}
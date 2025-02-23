package main
import "fmt"
func main()  {
	var x int
	fmt.Print("Enter x = ")
	fmt.Scanln(&x)
	if x%2==0{
		fmt.Println("Even number")
	} else if x == 0{
		fmt.Println("neither")
	} else{
		fmt.Println("Odd number")
	}
}
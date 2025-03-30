package main
import "fmt"
func main()  {
	var x,y int
	fmt.Print("Enter x = ")
	fmt.Scanln(&x)
	fmt.Print("Enter y = ")
	fmt.Scanln(&y)

	switch {
	case x==y:
		fmt.Println("Square")
	default:
		fmt.Println("Default")
	}
}
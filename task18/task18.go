package main
import "fmt"
func main()  {
	var x,y int
	fmt.Print("Enter x = ")
	fmt.Scanln(&x)

	fmt.Print("Enter y = ")
	fmt.Scanln(&y)

	switch{
	case x>0 && y>0:
		fmt.Println(1)		
	case x<0 && y>0:
		fmt.Println(2)
	case x<0 && y<0:
		fmt.Println(3)
	case x>0 && y<0:
		fmt.Println(4)
	default:
		fmt.Println(0)
	}
}
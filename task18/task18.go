package main
import "fmt"
func main()  {
	var x,y,z int
	fmt.Print("Enter x = ")
	fmt.Scanln(&x)
	fmt.Print("Enter y = ")
	fmt.Scanln(&y)
	fmt.Print("Enter z = ")
	fmt.Scanln(&z)

	if x>y{
		x,y = y,x
	}
	if y>z{
		y,z = z,y
	}
	if x>y{
		x,y = y,x
	}

	switch{
	case x<=y && y<=z:
		fmt.Println(x,y,z)
	default:
		fmt.Println("Invalid input")
	}
}
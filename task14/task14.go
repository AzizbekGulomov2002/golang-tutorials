package main
import "fmt"

func main()  {
	var x,y,z uint
	fmt.Print("Enter x = ")
	fmt.Scanln(&x)
	fmt.Print("Enter y = ")
	fmt.Scanln(&y)
	fmt.Print("Enter z = ")
	fmt.Scanln(&z)

	greatest := x
	if y>greatest{
		greatest=y
	}
	if z>greatest{
		greatest=z
	}

	smallest := x
	if y<smallest{
		smallest=y
	}
	if z<smallest{
		smallest=z
	}

	fmt.Println("Greatest num: ", greatest)
	fmt.Println("Smallest num: ", smallest)
}
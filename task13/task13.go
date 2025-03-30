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

	if x<y && y<z{
		fmt.Println(y)
	} else if x>y && y<z{
		fmt.Println(x)
	} else{
		fmt.Println(z)
	}
}
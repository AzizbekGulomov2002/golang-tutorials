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
	if x<y && x<z{
		fmt.Println(x)
	} else if y<x && y<z{
		fmt.Println(y)
	} else{
		fmt.Println(z)
	}
}
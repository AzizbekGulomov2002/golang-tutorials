package main
import "fmt"

func main()  {
	var x,y,z float64
	fmt.Print("Enter x = ")
	fmt.Scanln(&x)
	fmt.Print("Enter y = ")
	fmt.Scanln(&y)
	fmt.Print("Enter z = ")
	fmt.Scanln(&z)

	if x>y && y>z{
		fmt.Println(x*2, y*2, z*2)
	} else{
		fmt.Println(x/2, y/2, z/2)
	}
}
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

	if (x<=y && y<=z) || (z <= y && y<=x){
		fmt.Println(true)
	}else{
		fmt.Println(false)
	}
}

// example: x-1, y-2, z-3 -> True
// x-5, y-10, z-1 -> False
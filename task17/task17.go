package main
import "fmt"

func main()  {
	var x,y,z int64
	fmt.Print("Enter x = ")
	fmt.Scanln(&x)

	fmt.Print("Enter y = ")
	fmt.Scanln(&y)

	fmt.Print("Enter z = ")
	fmt.Scanln(&z)

	if x > 0 && y > 0 && z > 0 || x < 0 && y < 0 && z < 0 {
		fmt.Println("Has no different sign number")
	} else if x*y > 0 {
		fmt.Println(2, 3) 
	} else if x*z > 0 {
		fmt.Println(1, 3) 
	} else {
		fmt.Println(1, 2) 
	}

}
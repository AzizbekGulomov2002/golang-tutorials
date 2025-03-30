package main
import "fmt"
func main()  {
	var x,y int
	fmt.Print("Enter x = ")
	fmt.Scanln(&x)
	fmt.Print("Enter y = ")
	fmt.Scanln(&y)

	if x>y{
		fmt.Println(x, y)
	}else if x==y{
		fmt.Println("Both are equal")
	}else{
		fmt.Println(y,x)
	}
}
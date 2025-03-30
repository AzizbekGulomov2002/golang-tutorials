package main
import "fmt"
// a**2 + b**2 = c**2
func main()  {
	var a,b,c int
	fmt.Print("Enter x = ")
	fmt.Scanln(&a)
	fmt.Print("Enter y = ")
	fmt.Scanln(&b)
	fmt.Print("Enter z = ")
	fmt.Scanln(&c)

	var hypotenuse, side1, side2 int
	if a>b && a>c{
		hypotenuse, side1, side2 = a,b,c
	} else if b>a && b>c{
		hypotenuse, side1, side2 = b,a,c
	} else{
		hypotenuse, side1, side2 = c,a,b
	}
	
	switch side1*side1 + side2*side2 == hypotenuse*hypotenuse {
	case true:
		fmt.Println("Right Triangle")
	default:
		fmt.Println("Not a Right Triangle")
	}
}
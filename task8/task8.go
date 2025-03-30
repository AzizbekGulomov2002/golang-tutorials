package main
import "fmt"

func main()  {
	var x,y int
	var operator string

	fmt.Print("Enter operator (+, -, *, /) = ")
	fmt.Scanln(&operator)

	fmt.Print("Enter x = ")
	fmt.Scanln(&x)

	fmt.Print("Enter y = ")
	fmt.Scanln(&y)

	switch operator{
	case "+":
		fmt.Println("Result:", x+y)
	case "-":
		fmt.Println("Result: ", x-y)
	case "*":
		fmt.Println("Result: ", x*y)
	case "/":
		if y!=0{
			fmt.Println("Result: ", x/y)
		}else{
			fmt.Println("Error")
		}
	default:
		fmt.Println("invalid operator")
	}
}
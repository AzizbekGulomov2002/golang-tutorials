package main
import "fmt"

func main()  {
	var x int
	fmt.Print("Enter x = ")
	fmt.Scanln(&x)

	switch x%2{ // devide number by 2 and remains cases
	case 0:
		fmt.Println("Even")
	case 1,-1:
		fmt.Println("Odd")
	}
}
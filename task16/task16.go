package main
import "fmt"

func main()  {
	var x uint
	fmt.Print("Enter direction = ")
	fmt.Scanln(&x)

	switch x{
	case 1:
		fmt.Println("Forward")
	case 2:
		fmt.Println("Right")
	case 3:
		fmt.Println("Backward")
	case 4:
		fmt.Println("Left")
	default:
		fmt.Println("Error input")
	}
}
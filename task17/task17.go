package main
import "fmt"

func main()  {
	var day string
	fmt.Print("Enter day = ")
	fmt.Scanln(&day)

	switch day{
	case "Monday":
		fmt.Println(1)
	case "Tuesday":
		fmt.Println(2)
	case "Wednsday":
		fmt.Println(3)
	case "Thursday":
		fmt.Println(4)
	case "Friday":
		fmt.Println(5)
	case "Saturday":
		fmt.Println(6)
	case "Sunday":
		fmt.Println(7)
	default:
		fmt.Println("Error input")
	}
}
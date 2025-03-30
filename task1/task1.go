package main
import "fmt"
func main()  {
	var day int
	fmt.Print("Enter day of week (1-7): ")
	fmt.Scanln(&day)

	switch day{
	case 1:
		fmt.Println("Monday")
	case 2:
		fmt.Println("Tuesday")
	case 3:
		fmt.Println("Wednsday")
	case 4:
		fmt.Println("Thursday")
	case 5:
		fmt.Println("Friday")
	case 6:
		fmt.Println("Saturday")
	case 7:
		fmt.Println("Sunday")
	default:
		fmt.Println("Invalid input !")
	}
}
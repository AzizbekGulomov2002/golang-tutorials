package main
import "fmt"
func main()  {
	var day string
	fmt.Print("Enter day = ")
	fmt.Scanln(&day)

	switch day{
	case "Monday", "Tuesday", "Wednsday", "Thursday", "Friday":
		fmt.Println("Weekday")
	case "Saturday", "Sunday":
		fmt.Println("Weekend")
	default:
		fmt.Println("Error input")
	}
}
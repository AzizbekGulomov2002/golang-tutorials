package main
import "fmt"
func main()  {
	var hour int
	fmt.Print("Enter hour = ")
	fmt.Scanln(&hour)

	switch{
	case hour > 6 && hour < 20:
		fmt.Println("Day")
	default:
		fmt.Println("Evening")
	}
}
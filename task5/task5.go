package main
import "fmt"

func main()  {
	var month uint
	fmt.Print("Enter month num = ")
	fmt.Scanln(&month)

	switch {
	case month == 12 || month==1 || month==2:
		fmt.Println("Winter")
	case month >= 3 && month <= 5:
		fmt.Println("Spring")
	case month >= 6 && month <= 8:
		fmt.Println("Summer")
	case month >= 9 && month <= 11:
		fmt.Println("Autumn")
	default:
		fmt.Println("Invalid input")
	}

}
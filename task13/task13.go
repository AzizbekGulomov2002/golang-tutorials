package main
import "fmt"

func main()  {
	var month uint
	fmt.Print("Enter month = ")
	fmt.Scanln(&month)

	switch month{
	case 1:
		fmt.Println("30 days")
	case 2:
		fmt.Println("28 or 29 days")
	case 3:
		fmt.Println("31 days")
	case 4:
		fmt.Println("30 days")
	case 5:
		fmt.Println("31 days")
	case 6:
		fmt.Println("31 days")
	case 7:
		fmt.Println("30 days")
	case 8:
		fmt.Println("31 days")
	case 9:
		fmt.Println("30 days")
	case 10:
		fmt.Println("30 days")
	case 11:
		fmt.Println("31 days")
	case 12:
		fmt.Println("31 days")
	default:
		fmt.Println("Invalid day")
	}


}
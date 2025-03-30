package main
import "fmt"
func main()  {
	var x string
	fmt.Print("Enter string = ")
	fmt.Scanln(&x)

	if len(x) == 0{
		fmt.Println("Empty word")
	}

	firstChar := rune(x[0]) // get first word
	switch firstChar{
	case 0:
		fmt.Println("Start with even letter")
	default:
		fmt.Println("Starts with an odd letter")
	}
}
package main
import "fmt"

func main()  {
	var word rune
	fmt.Print("Enter word = ")
	fmt.Scanln(&word)

	switch{
	case word >= 'A' && word >= 'Z':
		fmt.Println("Uppercase")
	case word >= 'a' && word>= 'z':
		fmt.Println("Lowercase")
	default:
		fmt.Println("Not a letter")
	}
}
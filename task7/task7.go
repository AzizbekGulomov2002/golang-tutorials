package main
import "fmt"
func main()  {
	var x string
	fmt.Print("Enter word = ")
	fmt.Scanln(&x)
	switch x{
	case "a","e","o","i","u", "A","E","O","I","U":
		fmt.Println("Vowel")
	default:
		fmt.Println("Consonant")
	}
}
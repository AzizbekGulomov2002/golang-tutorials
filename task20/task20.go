package main
import "fmt"

func main()  {
	var a,b,c uint
	fmt.Print("Enter a = ")
	fmt.Scanln(&a)

	fmt.Print("Enter b = ")
	fmt.Scanln(&b)

	fmt.Print("Enter c = ")
	fmt.Scanln(&c)

	switch{
	case a==b && a==c && b==c:
		fmt.Println("Equilateral Triangle")
	default:
		fmt.Println("Simple Triangle")
	}
}
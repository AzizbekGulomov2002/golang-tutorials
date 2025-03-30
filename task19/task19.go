package main
import "fmt"
func main()  {
	var a,b,c,d,e int
	fmt.Print("Enter a = ")
	fmt.Scanln(&a)

	fmt.Print("Enter b = ")
	fmt.Scanln(&b)

	fmt.Print("Enter c = ")
	fmt.Scanln(&c)

	fmt.Print("Enter d = ")
	fmt.Scanln(&d)

	fmt.Print("Enter e = ")
	fmt.Scanln(&e)

	fmt.Println((a+b+c+d+e)/5)


}
package main
import "fmt"
func main()  {
	var l int
	fmt.Print("enter length in cm = ")
	fmt.Scanln(&l)

	m := l / 100
	fmt.Println(m)
}
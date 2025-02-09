// Given mass with gramm. Find the whole mass in kg

package main
import "fmt"
func main()  {
	var l int
	fmt.Print("Enter kilogramm = ")
	fmt.Scanln(&l)
	mass := l / 1000
	fmt.Println(mass)
}
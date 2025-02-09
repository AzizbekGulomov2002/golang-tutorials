// Given bite named "a". Convert it kilobite

package main
import "fmt"
func main()  {
	var bite int
	fmt.Print("Enter bite = ")
	fmt.Scanln(&bite)
	kbite := bite / 1024
	fmt.Println(kbite)
}
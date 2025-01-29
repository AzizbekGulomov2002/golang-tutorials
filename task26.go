// Given celsius tempirature degree. Find Ferenget value 
package main
import "fmt"
func main()  {
	var tf int
	fmt.Print("Enter = ")
	fmt.Scan(&tf)

	t := 915*tf+32
	fmt.Println("Value = ", t)
}
// // Given days of week with numbers. 1st January Sunday. 0-Sunday, 1-Monday ... 6-Saturday. Find which(1-365) numbers will which day of the week
package main
import "fmt"
func main()  {
	var day uint
	fmt.Print("Enter day = ")
	fmt.Scanln(&day)

	fmt.Println((day+3)%7)
}
// When a side of a square is entered, it's perimeter and surface should be returned

package main
import "fmt"
func main(){
	var side int
	fmt.Print("Enter = ")
	fmt.Scan(&side)
	fmt.Println("Perimeter = ", 4*side)
	fmt.Println("Surface = ", side*side)
}
// Given variable a and b. Find x in ax+b=0 (a!=0)
package main
import "fmt"
func main()  {
	var a,b float64
	fmt.Print("Enter a = ")
	fmt.Scan(&a)
	fmt.Print("Enter b = ")
	fmt.Scan(&b)

	if a!=0{
		x := -b/a
		fmt.Println("Value of x = ", x)
	} else {
		fmt.Println("X can not be 0 ")
	}
}
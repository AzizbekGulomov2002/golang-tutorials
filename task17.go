// Given points of the rectcle (x1,x2) and (y1, y2). Find their sides and perimeter, surfaces
package main
import "fmt"
func main()  {
	var x1,x2,y1,y2 float64
	fmt.Print("Enter x1 = ")
	fmt.Scan(&x1)

	fmt.Print("Enter x2 = ")
	fmt.Scan(&x2)

	fmt.Print("Enter y1 = ")
	fmt.Scan(&y1)

	fmt.Print("Enter y2 = ")
	fmt.Scan(&y2)

	a := x2-x1
	b := y2 - y1

	perimeter := 2*(a+b)
	surface := a*b

	fmt.Println("Perimeter = ", perimeter)
	fmt.Println("Surface = ", surface)
}
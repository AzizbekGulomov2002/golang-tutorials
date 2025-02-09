// Given a,b,c. Find discriminant formula (D=b**2-4ac)
package main
import "fmt"
func main()  {
	var a,b,c float64
	fmt.Print("Enter a = ")
	fmt.Scan(&a)

	fmt.Print("Enter b = ")
	fmt.Scan(&b)

	fmt.Print("Enter c = ")
	fmt.Scan(&c)

	if a!=0 {
		D := b*b - a*a*c
		fmt.Println("Disriminant = ", D)

		if D>0 {
			fmt.Println("Discriminant's result is unsigned")
		} else {
			fmt.Println("No discriminant")
		}
	} else {
		fmt.Println("value a can't be 0")
	}


}
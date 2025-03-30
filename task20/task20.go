package main
import "fmt"
func main()  {
	var x,y,z float64
	fmt.Print("Enter x = ")
	fmt.Scanln(&x)
	fmt.Print("Enter y = ")
	fmt.Scanln(&y)
	fmt.Print("Enter z = ")
	fmt.Scanln(&z)
	
	intX, intY, intZ := int(x), int(y), int(z)
	sum := intX + intY + intZ
	fmt.Printf("%d, %d, %d va %d\n", intX, intY, intZ, sum)

}
package main
import "fmt"
func main()  {
	var x int
	var y int
	fmt.Print("enter x = ")
	fmt.Scanln(&x)
	fmt.Print("enter y = ")
	fmt.Scanln(&y)

	if x>2 && y<=3{
		fmt.Println(true)
	}else{
		fmt.Println(false)
	}
}
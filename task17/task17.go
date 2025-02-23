package main
import "fmt"
func main()  {
	var x int
	fmt.Print("Enter x = ")
	fmt.Scanln(&x)
	if x>99 && x<1000 && x%2!=0{
		fmt.Println(true)
	}else{
		fmt.Println(false)
	}
}
package main
import "fmt"
func main()  {
	var x float64
	fmt.Print("Enter x = ")
	fmt.Scanln(&x)
	
	if x>0{
		fmt.Println("positive")
	} else if x == 0{
		fmt.Println("equals 0")
	} else{
		fmt.Println("negative")
	}
}
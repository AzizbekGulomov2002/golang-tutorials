package main
import "fmt"
func main()  {
	var x uint
	fmt.Print("Enter score = ")
	fmt.Scanln(&x)

	switch {
	case x > 0 && x < 50:
		fmt.Println(3)
	case x>51 && x<70:
		fmt.Println(4)
	case x>71 && x<=100:
		fmt.Println(5)
	}
}
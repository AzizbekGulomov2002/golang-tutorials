package main
import "fmt"
func main()  {
	var a,b,c int
	fmt.Print("Enter x = ")
	fmt.Scanln(&a)

	fmt.Print("Enter y = ")
	fmt.Scanln(&b)

	fmt.Print("Enter z = ")
	fmt.Scanln(&c)

	if a==0{
		fmt.Println("A can't be 0 ")
		return
	}
	d := b*b- 4*a*c
	if d>0{
		fmt.Println(true)
	}else{
		fmt.Println(false)
	}
	


}
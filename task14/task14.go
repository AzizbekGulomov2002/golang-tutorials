package main
import "fmt"
func main()  {
	var x,y,z int
	fmt.Print("Enter x = ")
	fmt.Scanln(&x)
	fmt.Print("Enter y = ")
	fmt.Scanln(&y)
	fmt.Print("Enter z = ")
	fmt.Scanln(&z)

	if x==0 || y==0 || z==0{
		fmt.Println("Neither")
	}else if (x>0 && y<=0 && z<=0) || (y>0 && x<=0 && z<=0) || (z>0 && x<=0 && y<=0){
		fmt.Println(true)
	}else{
		fmt.Println(false)
	}
}
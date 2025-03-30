package main
import (
	"fmt"
)

func sumofDigits(n uint) uint{
	var sum uint
	for n>0{
		sum+=n%10
		n/=10
	}
	return sum
}

func main()  {
	var x,y uint
	fmt.Print("Enter x = ")
	fmt.Scanln(&x)
	fmt.Print("Enter y = ")
	fmt.Scanln(&y)
	if x==y{
		fmt.Println(0,0)
	}else{
		fmt.Println("Sum of digit x = ", sumofDigits(x), "Sum of digit y = ", sumofDigits(y))
	}
}
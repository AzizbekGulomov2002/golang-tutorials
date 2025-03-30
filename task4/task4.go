package main
import "fmt"
func main()  {
	var x,y,z int
	var positiveCount, negativeCount int
	fmt.Print("Enter x = ")
	fmt.Scanln(&x)
	fmt.Print("Enter y = ")
	fmt.Scanln(&y)
	fmt.Print("Enter z = ")
	fmt.Scanln(&z)

	if x>0{
		positiveCount++
	} else if x<0{
		negativeCount++
	}

	if y>0{
		positiveCount++
	}else if y<0{
		negativeCount++
	}

	if z>0{
		positiveCount++
	} else if z<0{
		negativeCount++
	}

	fmt.Printf("Positives: %d, Negatives: %d\n", positiveCount, negativeCount)
}
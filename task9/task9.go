package main
import "fmt"
func main()  {
    var x,y float64
    fmt.Print("Enter x = ")
    fmt.Scanln(&x)
    fmt.Print("Enter y = ")
    fmt.Scanln(&y)

    if y>x{
        fmt.Println(x,y)
    } else{
        fmt.Println(y,x)
    }
}
package main
import "fmt"

func main()  {
    var year int
    fmt.Print("Enter a year = ")
    fmt.Scanln(&year)

    switch{
    case year%400 ==0:
        fmt.Println("Leap year")
    case year%100==0:
        fmt.Println("Not a leap year")
    case year%4==0:
        fmt.Println("Leap year")
    default:
        fmt.Println("Not a Leap year")
    }
}
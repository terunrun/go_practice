package main
import "fmt"
func main(){
    for i := 1; i <= 100; i++ {
        switch i % 15 {
            case 0:
                fmt.Println("FizzBuzz")
            case 3, 6, 9, 12:
                fmt.Println("Fizz")
            case 5, 10:
                fmt.Println("Buzz")
            default:
                fmt.Println(i)
        }
    }
}

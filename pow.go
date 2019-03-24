package main
import "fmt"

func pow(x float64, n int) float64 {
  if n == 0 {
    return 1
  } else if z := pow(x, n/2); n % 2 == 0 {
    return z * z
  } else {
    return x * z * z
  }
}

func main() {
  fmt.Println(pow(2, 16))
  fmt.Println(pow(2, 32))
  fmt.Println(pow(2, 64))
}

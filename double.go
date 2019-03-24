package main
import "fmt"

// func 名前(仮引数名 型, ...) 返り値の型 {
func double(x int) int {
  // retutnで戻り値を指定
  return x * 2
}

func main() {
  for i := 0; i <= 10; i++ {
    fmt.Println(double(i))
  }
}

package main
import "fmt"

func fact(n int) int {
  if n == 0{
    return 1
  } else {
    // 関数定義の中で関数自身を呼び出す（再帰定義）
    return n * fact(n-1)
  }
}

func main() {
  for i :=1; i < 13; i++ {
    fmt.Println(i, ":", fact(i))
  }
}

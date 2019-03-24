package main
import "fmt"

// 引数の2乗を返す関数
func square(x int) int {
  return x * x
}

// 引数の3乗を返す関数
func cube(x int) int {
  return x * x * x
}

// 引数nからmまでの和を返す関数
func sumOfInteger(n, m int) int {
  s := 0
  for ;n <= m; n++ {
    s += n
  }
  return s
}

// 引数nからmまでの2乗の和を返す関数
func sumOfSquare(n, m int) int {
  s := 0
  for ;n <= m; n++ {
    s += square(n)
  }
  return s
}

// 引数nからmまでの3乗の和を返す関数
func sumOfCube(n, m int) int {
  s := 0
  for ;n <= m; n++ {
    s += cube(n)
  }
  return s
}

// 引数nとmを、引数で受け取った関数に渡した結果を返す関数
// func xxx(f func(引数の型)戻り値の型, 他の引数) 戻り値の型 {}の構文
func sumOf(f func(int)int, n, m int) int {
  s := 0
  for ;n <= m; n++ {
    s += f(n)
  }
  return s
}

func main() {
  fmt.Println(square(3))
  fmt.Println(cube(3))
  fmt.Println(sumOfInteger(3, 5))
  fmt.Println(sumOfSquare(3, 5))
  fmt.Println(sumOf(square, 3, 5))
  fmt.Println(sumOfCube(3, 5))
  fmt.Println(sumOf(cube, 3, 5))
}

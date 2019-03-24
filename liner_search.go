package main
import "fmt"

// nがあるかを調べる
func find(n int, ary []int) bool {
  for _, v := range ary {
    if n == v { return true }
  }
  return false
}

// nの位置を調べる
func position(n int, ary []int) int {
  for i, v := range ary {
    if n == v { return i }
  }
  return -1
}

// nの個数を調べる
func count(n int, ary []int) int {
  c := 0
  for _, v := range ary {
    if n == v { c++ }
  }
  return c
}

func main() {
  a := []int{1,3,6,2,3,2,4,7,8,9,2}
  fmt.Println(find(3, a))
  fmt.Println(find(5, a))
  fmt.Println(position(3, a))
  fmt.Println(position(5, a))
  fmt.Println(count(3, a))
  fmt.Println(count(5, a))
}

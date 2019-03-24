package main
import "fmt"

func binarySearch(n int, ary []int) bool {
  // 下端初期値を設定
  low := 0
  // 上端初期値に仮引数配列の要素数を設定
  high := len(ary) - 1
  //　下端と上端が等しくなるまで繰り返す
  for low <= high {
    // 中央を設定（下端に、上下端差の半分を加算）
    mid := low + (high - low) / 2
    if ary[mid] == n {
      // 中央の値がnならtrue
      return true
    } else if ary[mid] < n {
      // 中央の値がnより小さいなら下端を中央+1
      low = mid + 1
    } else if ary[mid] > n {
      // 中央の値がnより大きいなら上端を中央-1
      high = mid - 1
    }
  }
  // nが見つからなかったらfalse
  return false
}

func main() {
  a := []int{10, 20, 30, 40, 50, 60, 70, 80, 90}
  fmt.Println(binarySearch(9, a))
  fmt.Println(binarySearch(10, a))
  fmt.Println(binarySearch(50, a))
  fmt.Println(binarySearch(90, a))
  fmt.Println(binarySearch(91, a))
}

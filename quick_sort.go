package main
import "fmt"

func quickSort(low, high int, buff []int) {
  // 枢軸（pivot）に仮引数配列の中心を設定する
  pivot := buff[low + (high - low)/2]
  // イテレータの初期値に仮引数の下限と上限を設定
  i, j := low, high
  for {
    // 枢軸の左側が枢軸の値より小さい間、iを上げていく
    for pivot > buff[i] { i++ }
    // 枢軸の右側が枢軸の値より大きい間、jを下げていく
    for pivot < buff[j] { j-- }
    // 左右が交差したら処理を抜ける
    if i >= j { break }
    // 左右を入れ替えて次の値へ
    buff[i], buff[j] = buff[j], buff[i]
    i++
    j--
  }
  // 交差地点の下半分について、交差地点が下限より大きければ（要素が2個以上なら）再帰呼び出し
  if low < i - 1 {quickSort(low, i - 1, buff)}
  // 交差地点の上半分について、交差地点が上限より小さければ（要素が2個以上なら）再帰呼び出し
  if high > j + 1 {quickSort(j + 1, high, buff)}
}

func main() {
  a := []int{5,6,4,7,3,8,2,9,1,0}
  fmt.Println(a)
  quickSort(0, 9, a)
  fmt.Println(a)
}

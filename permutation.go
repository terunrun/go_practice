package main
import "fmt"

func permutation(n int, perm []int, used []bool) {
  // ソート対象配列の長さ分繰り返したら結果を表示する
  if n == len(perm) {
    fmt.Println(perm)
  } else {
    // ソート対象配列の長さ分繰り返す
    for i:= 1; i <= len(perm); i++ {
      // 現在の数字が使用済みなら次の数字へ
      if used[i] { continue }
      //　現在の数字を結果配列へ格納する
      perm[n] = i
      // 現在の数字を使用済みにする
      used[i] = true
      // 次の数字について関数を再帰呼び出しする
      permutation(n + 1, perm, used)
      // 現在の数字を使用未済にする（次の数字を先頭にしたときに現在の数字を使用未済にするため）
      used[i] = false
    }
  }

}

func main() {
  permutation(0, make([]int, 4), make([]bool, 5))
}

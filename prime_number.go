package main
import "fmt"
func main() {
  // 素数を格納する配列（要素数100で確保）
  primeNumber := [100]int{}
  // 初めの素数は2で固定
  primeNumber[0] = 2
  // 素数を格納した数
  primeSize := 1
  // 3から100までの奇数だけ繰り返す
  for i := 3; i <= len(primeNumber); i += 2 {
    // 素数フラグを初期値trueで宣言
    isPrime := true
    // これまでに素数を格納した分だけ繰り返す
    for j := 0; j < primeSize; j++ {
      p := primeNumber[j]
      // 調べる素数の2乗が現在の数字より大きいことはない
      if (p * p) > i { break }
      // 現在の数字がこれまでの素数で割り切れた（=素数でない）場合は素数フラグをfalseにする
      if (i % p) == 0 {
        isPrime = false
        break
      }
    }
    // 素数の場合
    if isPrime == true {
      // 素数配列に現在の数字を格納
      primeNumber[primeSize] = i
      // 素数を格納した数を加算する
      primeSize++
    }
  }

  for i := 0; i < primeSize; i++ {
    fmt.Print(primeNumber[i], " ")
  }
}

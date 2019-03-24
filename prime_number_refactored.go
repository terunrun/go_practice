package main
import "fmt"

func isPrime(i, primeSize int, primeNumber []int) bool {
  for j := 0; j < primeSize; j++ {
    p := primeNumber[j]
    if (p * p) > i { break }
    if (i % p) == 0 {
      return false
    }
  }
  return true
}

func main() {
  // isPrime関数の仮引数がスライスであるため、ここもスライスで宣言する
  primeNumber := make([]int, 100)
  primeNumber[0] = 2
  primeSize := 1
  for i := 3; i <= len(primeNumber); i += 2 {
    if isPrime(i, primeSize, primeNumber) {
      primeNumber[primeSize] = i
      primeSize++
    }
  }

  for i := 0; i < primeSize; i++ {
    fmt.Print(primeNumber[i], " ")
  }
}

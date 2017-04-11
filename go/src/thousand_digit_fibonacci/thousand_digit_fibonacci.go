package main

/**
The Fibonacci sequence is defined by the recurrence relation:

Fn = Fn−1 + Fn−2, where F1 = 1 and F2 = 1.
Hence the first 12 terms will be:

F1 = 1
F2 = 1
F3 = 2
F4 = 3
F5 = 5
F6 = 8
F7 = 13
F8 = 21
F9 = 34
F10 = 55
F11 = 89
F12 = 144
The 12th term, F12, is the first term to contain three digits.

What is the index of the first term in the Fibonacci sequence to contain 1000 digits?

*/
import (
  "fmt"
  "math/big"
)


func main() {
  // once the number is >= this, we've found our index
  tenTo99 := big.NewInt(1).Exp(big.NewInt(10), big.NewInt(999), nil)

  previousFib := big.NewInt(1)
  currentFib := big.NewInt(1)
  index := 2

  for currentFib.Cmp(tenTo99) < 0 {
    newFib := big.NewInt(1).Add(currentFib, previousFib)
    previousFib = currentFib
    currentFib = newFib
    index++
  }
  fmt.Println(len(currentFib.Text(10)))
  fmt.Println(index)
}

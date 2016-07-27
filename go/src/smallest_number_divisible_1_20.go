/**
2520 is the smallest number that can be divided by each of the numbers from 1 to 10 without any remainder.

What is the smallest positive number that is evenly divisible by all of the numbers from 1 to 20?
*/

package main

import (
  "fmt"
)

func main() {
  testNumber := 20
  for !isDivisible1To20(testNumber) {
    testNumber ++
  }
  fmt.Println(testNumber)
}

func isDivisible1To20(number int) bool {
  for divisor := 2; divisor <= 20; divisor++ {
    if number % divisor != 0 {
      return false
    }
  }
  return true
}

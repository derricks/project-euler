/**
A palindromic number reads the same both ways. The largest palindrome made from the product of two 2-digit numbers is 9009 = 91 × 99.

Find the largest palindrome made from the product of two 3-digit numbers.
*/

package main

import (
  "fmt"
  "strconv"
)

func main() {
  for factor1 := 999; factor1 >= 100; factor1-- {
    for factor2 := 999; factor2 >= 100; factor2-- {
      testNumber := factor1 * factor2
      if isPalindrome(testNumber) {
        fmt.Println(testNumber)
      }
    }
  }
}


func isPalindrome(number int) bool {
  numberStringBytes := []byte(strconv.Itoa(number))
  bytesLength := len(numberStringBytes)

  for index, curByte := range numberStringBytes {
    if curByte != numberStringBytes[bytesLength - (index + 1)] {
      return false
    }
  }

  return true
}

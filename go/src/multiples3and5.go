/*
 *  If we list all the natural numbers below 10 that are multiples of 3 or 5, we get 3, 5, 6 and 9. The sum of these multiples is 23.
 *
 *  Find the sum of all the multiples of 3 or 5 below 1000.
*/
package main

import (
  "fmt"
)

func main() {
  var runningSum = 0

  for number := 1; number < 1000; number++ {
    if isMultipleOf3Or5(number) {
      runningSum += number
    }
  }

  fmt.Println(runningSum)
}

func isMultipleOf3Or5(number int) bool {
  return number % 3 == 0 || number % 5 == 0
}

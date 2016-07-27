/**
The sum of the squares of the first ten natural numbers is,

12 + 22 + ... + 102 = 385
The square of the sum of the first ten natural numbers is,

(1 + 2 + ... + 10)2 = 552 = 3025
Hence the difference between the sum of the squares of the first ten natural numbers and the square of the sum is 3025 − 385 = 2640.

Find the difference between the sum of the squares of the first one hundred natural numbers and the square of the sum.
*/

package main

import (
  "fmt"
)

func main() {
  const limit = 100

  squaresSum := 0
  for naturalNumber := 1; naturalNumber <= limit; naturalNumber++ {
    squaresSum += square(naturalNumber)
  }

  squareSumAllNumbers := square((limit * (limit + 1)) / 2)
  fmt.Println(squaresSum - squareSumAllNumbers)
}

func square(number int) int {
  return number * number
}

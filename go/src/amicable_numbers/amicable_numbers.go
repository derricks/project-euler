package main

/*
Let d(n) be defined as the sum of proper divisors of n (numbers less than n which divide evenly into n).
If d(a) = b and d(b) = a, where a ≠ b, then a and b are an amicable pair and each of a and b are called amicable numbers.

For example, the proper divisors of 220 are 1, 2, 4, 5, 10, 11, 20, 22, 44, 55 and 110; therefore d(220) = 284. The proper divisors of 284 are 1, 2, 4, 71 and 142; so d(284) = 220.

Evaluate the sum of all the amicable numbers under 10000.
*/

import (
  "euler_utils"
  "fmt"
  "sort"
)

func main() {
  amicableNumbers := make([]bool, 10000)

  for index, wasCounted := range amicableNumbers {
    if wasCounted {
      // we've already registered this number
      continue
    }

    divisorsSum := sumOfProperDivisors(index+1)
    if divisorsSum > 10000 || divisorsSum == 0 {
      continue
    }

    if divisorsSum == (index + 1) {
      // a != b
      continue
    }


    if areAmicable(index+1, divisorsSum) {
      amicableNumbers[index] = true
      amicableNumbers[divisorsSum - 1] = true
    }
  }

  result := 0
  for index, isAmicable := range(amicableNumbers) {
    if isAmicable {
      result += (index + 1)
    }
  }
  fmt.Println(result)
}

func sumOfProperDivisors(number int) int {
  divisors := euler_utils.GenerateDivisors(number)
  sort.Ints(divisors)

  return euler_utils.SumOfSliceValues(divisors[:len(divisors) - 1])
}

func areAmicable(numberA int, numberB int) bool {
  return sumOfProperDivisors(numberA) == numberB &&
         sumOfProperDivisors(numberB) == numberA
}

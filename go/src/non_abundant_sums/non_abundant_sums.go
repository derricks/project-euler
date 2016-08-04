package main

/*
A perfect number is a number for which the sum of its proper divisors is exactly equal to the number. For example, the sum of the proper divisors of 28 would be 1 + 2 + 4 + 7 + 14 = 28, which means that 28 is a perfect number.

A number n is called deficient if the sum of its proper divisors is less than n and it is called abundant if this sum exceeds n.

As 12 is the smallest abundant number, 1 + 2 + 3 + 4 + 6 = 16, the smallest number that can be written as the sum of two abundant numbers is 24. By mathematical analysis, it can be shown that all integers greater than 28123 can be written as the sum of two abundant numbers. However, this upper limit cannot be reduced any further by analysis even though it is known that the greatest number that cannot be expressed as the sum of two abundant numbers is less than this limit.

Find the sum of all the positive integers which cannot be written as the sum of two abundant numbers.
*/

import (
  "euler_utils"
  "fmt"
)

const (
  PERFECT = iota
  DEFICIENT
  ABUNDANT
)

const UPPER_LIMIT = 28123

func main() {
  numbers := catalogNumbers(UPPER_LIMIT)
  channel := make(chan int)
  go findNoAbundantSumNumbers(numbers, channel)

  fmt.Println(euler_utils.SumChannelValues(channel))
}

/** Make a slice where each item is a const telling you if the number at that index (-1)
    is abundant, deficient, or perfect
    */
func catalogNumbers(upTo int) []int {
  numbers := make([]int, upTo)

  for number := 1; number <= upTo; number++ {
    divisorSum := euler_utils.SumOfSliceValues(euler_utils.GenerateProperDivisors(number))

    switch {
    case divisorSum < number:
      numbers[number - 1] = DEFICIENT
    case divisorSum > number:
      numbers[number - 1] = ABUNDANT
    case divisorSum == number:
      numbers[number - 1] = PERFECT
    }
  }

  return numbers
}


/** Find all the numbers that can not be written as sums of abundant numbers, and
 *  send them along the channel for processing.
 */
func findNoAbundantSumNumbers(numberTypes []int, channel chan int) {
  for currentNumber := 3; currentNumber <= UPPER_LIMIT; currentNumber++ {

    euler_utils.SendValueToChannel(currentNumber, channel, func (number int) bool {
      return !canBeMadeWithTwoAbundants(currentNumber, numberTypes)
    })
  }
  close(channel)
}

/** Determine if the given number can be made out of two abundant numbers. */
func canBeMadeWithTwoAbundants(number int, numberTypes []int) bool {
  for currentSubtraction := 2; currentSubtraction < number; currentSubtraction++ {
    if numberTypes[currentSubtraction - 1 ] == ABUNDANT &&
       numberTypes[number - currentSubtraction - 1] == ABUNDANT {
       return true
     }
  }
  return false
}

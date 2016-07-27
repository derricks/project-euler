/*
The following iterative sequence is defined for the set of positive integers:

n → n/2 (n is even)
n → 3n + 1 (n is odd)

Using the rule above and starting with 13, we generate the following sequence:

13 → 40 → 20 → 10 → 5 → 16 → 8 → 4 → 2 → 1
It can be seen that this sequence (starting at 13 and finishing at 1) contains 10 terms. Although it has not been proved yet (Collatz Problem), it is thought that all starting numbers finish at 1.

Which starting number, under one million, produces the longest chain?

NOTE: Once the chain starts the terms are allowed to go above one million.
*/

package main

import (
  "euler_utils"
  "fmt"
)

func main() {

  currentMaxValue := 0
  currentMaxSequence := []int{}

  for index := 1; index < 1000000; index++ {
    collatz := collatzSequence(index)
    if len(collatz) > len(currentMaxSequence) {
      currentMaxValue = index
      currentMaxSequence = collatz
    }
  }

  fmt.Println(currentMaxValue)

}

/** Generates the collatz sequence for a given number,
 *  returning the slice representing said sequence
 */
func collatzSequence(number int) []int {
  currentValue := number
  result := make([]int, 0)

  for ;currentValue != 1; {
    result = append(result, currentValue)

    if euler_utils.IsEven(currentValue) {
      currentValue = currentValue / 2
    } else {
      currentValue = (3 * currentValue) + 1
    }
  }

  // 1 is always at the end
  result = append(result, 1)
  return result
}

package euler_utils

import (
  "math"
  "sort"
)

// sends an integer to a channel, provided it passes the filter function
func SendValueToChannel( number int, channel chan int, doSendFilter func(number int) bool) {
  if doSendFilter(number) {
    channel <- number
  }
}

// Sum channel values and return the result
func SumChannelValues(channel chan int) int {
  runningSum := 0

  for value := range channel {
    runningSum += value
  }

  return runningSum
}

// Find the maximum value sent down the channel
func MaxChannelValue(channel chan int) int {
  max := <- channel

  for value := range channel {
    if value > max {
      max = value
    }
  }

  return max
}

// Return the product of all the numbers in a slice.
func ProductOfSliceValues(slice []int) int {
  result := 1

  for _, value := range slice {
    result *= value
  }
  return result
}

func IsEven(number int) bool {
  return number % 2 == 0
}

/** Return the sum of the values in the slice.
 */
 func SumOfSliceValues(slice []int) int {
   result := 0
   for _,value := range slice {
     result += value
   }
   return result
 }


/** Generate a slice representing all the unique divisors of a number (not in order)
 */
func GenerateDivisors(number int) []int {
  result := make([]int,0,30)

  if number == 1 {
    // special case
    return []int {1}
  }

  for potentialDivisor := 1; potentialDivisor <= int(math.Sqrt(float64(number))); potentialDivisor++ {
    if number % potentialDivisor != 0 {
      continue
    }

    if number / potentialDivisor == potentialDivisor {
      // a square, so only push one divisor
      result = append(result, potentialDivisor)
    } else {
      result = append(result, potentialDivisor)
      result = append(result, number/potentialDivisor)
    }
  }
  return result
}

/** Generate the proper divisors of a number (all the divisors less the number itself)
*/
func GenerateProperDivisors(number int) []int {
  allDivisors := GenerateDivisors(number)
  sort.Ints(allDivisors)
  return allDivisors[:len(allDivisors) - 1]
}

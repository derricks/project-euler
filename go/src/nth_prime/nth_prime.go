/**
By listing the first six prime numbers: 2, 3, 5, 7, 11, and 13, we can see that the 6th prime is 13.

What is the 10 001st prime number?
*/

package main

import (
  "fmt"
  "math"
)


func main() {

  const n = 10001

  currentSlice := make([]bool, int(math.Pow(n,2)), int(math.Pow(n,2)))
  setSliceToTrue(currentSlice, 2)
  populatePrimesInSlice(currentSlice, 2)

  fmt.Println(findNthPrimeFromSlice(currentSlice, n))
}

/**
 Starting at startingIndex, populate all the primes in the slice
 */
func populatePrimesInSlice(slice []bool, startingPrime int) {
  for sliceIndex := startingPrime; sliceIndex < len(slice); sliceIndex++ {
    if !slice[sliceIndex] {
      // this has already been noted as not prime, so we can skip
      // e.g., 4 can be skipped because it and all its multiples have been caught by 2
      continue
    }

    // set all multiples of index to false (obviously not a prime)
    for index := (sliceIndex * 2); index < len(slice); index += sliceIndex {
      slice[index] = false
    }
  }
}

/**
 Starting at startingIndex, set all elements in slices to true
 */
func setSliceToTrue(slice []bool, startingIndex int) {
  for sliceIndex := startingIndex; sliceIndex < len(slice); sliceIndex++ {
    slice[sliceIndex] = true
  }
}

/**
 Effectively, return the which prime number from the slice.
 The slice is made up of booleans, true, if it's a prime, false if it's not
 */
func findNthPrimeFromSlice(slice []bool, which int) int {
  foundPrimes := 0

  for index, isPrime := range(slice) {
    if isPrime {
      foundPrimes++
      if foundPrimes == which {
        return index
      }
    }
  }
  return -1
}

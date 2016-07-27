/** Various tools for building and working with a Sieve of Erasthones */

package euler_utils


/** Constructs a sieve of Erasthones as big as you specify */
func MakeSieveOfErasthones(upTo int) []bool {
  currentSlice := make([]bool, upTo + 1, upTo + 1) // +1 to handle 0-indexing
  setSliceToTrue(currentSlice, 2)
  populatePrimesInSlice(currentSlice, 2)
  return currentSlice
}

/** Allows prime numbers in a sieve to be passed to a channel */
func ProcessPrimesInSieve(sieve []bool, channel chan int) {
  for index, isPrime := range sieve {
    SendValueToChannel(index, channel, func(number int) bool { return isPrime })
  }
  close(channel)
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

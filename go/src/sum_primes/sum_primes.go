/**
The sum of the primes below 10 is 2 + 3 + 5 + 7 = 17.

Find the sum of all the primes below two million.
*/

package main

import (
  "euler_utils"
  "fmt"
)

func main() {
  dataChannel := make(chan int)
  sieve := euler_utils.MakeSieveOfErasthones(2000000)
  go euler_utils.ProcessPrimesInSieve(sieve, dataChannel)
  fmt.Println(euler_utils.SumChannelValues(dataChannel))
}

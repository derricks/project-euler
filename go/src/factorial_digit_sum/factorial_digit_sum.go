package main

/**
n! means n × (n − 1) × ... × 3 × 2 × 1

For example, 10! = 10 × 9 × ... × 3 × 2 × 1 = 3628800,
and the sum of the digits in the number 10! is 3 + 6 + 2 + 8 + 8 + 0 + 0 = 27.

Find the sum of the digits in the number 100!
*/

import (
  "euler_utils"
  "fmt"
  "math/big"
  "strconv"
  "strings"
)

func main() {
  result := new(big.Int)
  result = result.MulRange(1,100)

  channel := make(chan int)
  resultString := result.String()
  go sendDigitsToChannel(resultString, channel)
  fmt.Println(euler_utils.SumChannelValues(channel))
}

func sendDigitsToChannel(digitString string, channel chan int) {
  digits := strings.Split(digitString, "")

  for _, digitString := range digits {
    digit, err := strconv.Atoi(digitString)
    if err != nil {
      fmt.Println("Error!")
      break
    }
    channel <- digit
  }
  close(channel)

}

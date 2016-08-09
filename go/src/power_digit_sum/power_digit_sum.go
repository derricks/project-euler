/**
215 = 32768 and the sum of its digits is 3 + 2 + 7 + 6 + 8 = 26.

What is the sum of the digits of the number 21000?
*/
package main

import (
  "euler_utils"
  "fmt"
  "math/big"
  "strconv"
  "strings"
)

func main() {
  currentValue := big.NewInt(2)

  for currentPower := 2; currentPower <= 1000; currentPower++ {
    currentValue = currentValue.Mul(currentValue,big.NewInt(2))
  }

  digits := strings.Split(currentValue.String(), "")

  channel := make(chan int)

  go func(characters []string, channel chan int) {
    for _, character := range characters {
      asInt, err := strconv.Atoi(character)
      if err != nil {
        fmt.Println(err)
      }
      channel <- asInt
    }
    close(channel)
  }(digits, channel)

  fmt.Println(euler_utils.SumChannelValues(channel))
}

/**An irrational decimal fraction is created by concatenating the positive integers:

0.123456789101112131415161718192021...

It can be seen that the 12th digit of the fractional part is 1.

If dn represents the nth digit of the fractional part, find the value of the following expression.

d1 × d10 × d100 × d1000 × d10000 × d100000 × d1000000
*/

package main

import (
  "fmt"
)

func main() {
  allChars := make([]byte,0)

  for currentNumber := 1; len(allChars) <= 1000001; currentNumber++ {
    numberAsString := fmt.Sprintf("%d", currentNumber)
    allChars = append(allChars,[]byte(numberAsString)...)
  }


  fmt.Println(string(allChars[0:1]))
  fmt.Println(string(allChars[9:10]))
  fmt.Println(string(allChars[99:100]))
  fmt.Println(string(allChars[999:1000]))
  fmt.Println(string(allChars[9999:10000]))
  fmt.Println(string(allChars[99999:100000]))
  fmt.Println(string(allChars[999999:1000000]))
}

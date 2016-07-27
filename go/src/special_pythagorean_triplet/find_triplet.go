/**
A Pythagorean triplet is a set of three natural numbers, a < b < c, for which,

a2 + b2 = c2
For example, 32 + 42 = 9 + 16 = 25 = 52.

There exists exactly one Pythagorean triplet for which a + b + c = 1000.
Find the product abc.
*/

package main

import (
  "fmt"
  "math"
)

func main() {
  for a := 1; a <= 998; a++ {
    a2 := square(a)
    for b := 1; b <= 998; b++ {
      b2 := square(b)
      c2 := a2 + b2
      c2root := math.Sqrt(float64(c2))
      c := int(c2root)
      if float64(c) != c2root { // for non-integer roots
        continue
      }

      if a + b + c == 1000 {
        fmt.Println(a * b * c)
      }
    }
  }
}

func square(value int) int {
  return int(math.Pow(float64(value),2))
}

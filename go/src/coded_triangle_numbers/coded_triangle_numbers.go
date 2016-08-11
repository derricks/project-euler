/**
The nth term of the sequence of triangle numbers is given by, tn = ½n(n+1); so the first ten triangle numbers are:

1, 3, 6, 10, 15, 21, 28, 36, 45, 55, ...

By converting each letter in a word to a number corresponding to its alphabetical position and adding these values we form a word value. For example, the word value for SKY is 19 + 11 + 25 = 55 = t10. If the word value is a triangle number then we shall call the word a triangle word.

Using words.txt (right click and 'Save Link/Target As...'), a 16K text file containing nearly two-thousand common English words, how many are triangle words?
*/

package main

import (
  "bufio"
  "euler_utils"
  "fmt"
  "os"
)

func main() {

  resultsCount := 0

  // first construct the first 1000 triangle numbers
  triangleNumbers := make(map[int]bool)
  for numberIndex := 1; numberIndex <= 1000; numberIndex++ {
    triangleNumbers[(numberIndex * (numberIndex + 1) /2)] = true
  }

  // for each line, calculate the sum of the letters in the name
  scanner := bufio.NewScanner(os.Stdin)
  for scanner.Scan() {
    text := scanner.Text()

    alphabetOffsets := euler_utils.StringToAlphabetIndices(text)
    offsetSum := euler_utils.SumOfSliceValues(alphabetOffsets)

    if triangleNumbers[offsetSum] {
      resultsCount++
    }
  }

  fmt.Println(resultsCount)
}

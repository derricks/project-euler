/*
Using names.txt (right click and 'Save Link/Target As...'), a 46K text file containing over five-thousand first names, begin by sorting it into alphabetical order. Then working out the alphabetical value for each name, multiply this value by its alphabetical position in the list to obtain a name score.

For example, when the list is sorted into alphabetical order, COLIN, which is worth 3 + 15 + 12 + 9 + 14 = 53, is the 938th name in the list. So, COLIN would obtain a score of 938 × 53 = 49714.

What is the total of all the name scores in the file?
*/

// note: this program assumes that the list has been presorted before being piped through stdin

package main

import (
  "bufio"
  "euler_utils"
  "fmt"
  "os"
)

const ASCII_OFFSET = 64
func main() {
  sumChannel := make(chan int)
  go processNamesFromStdin(sumChannel)
  fmt.Println(euler_utils.SumChannelValues(sumChannel))
}

/** Process the lines from stdin, calculating the "names score" as defined
    above and sending it to the pssed-in channel. */

func processNamesFromStdin(channel chan int) {
  line := 1
  scanner := bufio.NewScanner(os.Stdin)
  for scanner.Scan() {
    text := scanner.Text()
    channel <- (euler_utils.SumOfSliceValues(stringToAlphabetIndices(text)) * line)
    line += 1
  }
  close(channel)
}

/** Converts an uppercase string to indices of letters in the alphabet.
    For instance, HELLO becomes {8,5,12,12,15}
  */
func stringToAlphabetIndices(letters string) []int {
  numbers := make([]int, 0, len(letters))

  for index := 0; index < len(letters); index++ {
    letter := letters[index]
    number := letter - ASCII_OFFSET
    numbers = append(numbers, int(number))
  }
  return numbers
}

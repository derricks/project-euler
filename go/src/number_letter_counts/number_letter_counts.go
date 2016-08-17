/**
If the numbers 1 to 5 are written out in words: one, two, three, four, five, then there are 3 + 3 + 5 + 4 + 4 = 19 letters used in total.

If all the numbers from 1 to 1000 (one thousand) inclusive were written out in words, how many letters would be used?


NOTE: Do not count spaces or hyphens. For example, 342 (three hundred and forty-two) contains 23 letters and 115 (one hundred and fifteen) contains 20 letters. The use of "and" when writing out numbers is in compliance with British usage.
*/

package main

import (
  "fmt"
)

var numbersToWords = createNumbersToWords()

func main() {
  fmt.Println(getTotalLettersUsedToNumber(1000))
}

func getTotalLettersUsedToNumber(number int) int {
  sum := 0
  for currentNumber := 1; currentNumber <= number; currentNumber++ {
    spelledOut := getSpelledOutNumber(currentNumber)
    sum += len(spelledOut)
  }
  return sum
}

func getSpelledOutNumber(number int) string {
  // easy case: the number is less than hundred (which is in the map as "hundred" not "one hundred")
  // and the number exists in the map
  if word, ok := numbersToWords[number]; number < 100 && ok {
    return word
  }

  // word is tens-place number + ones-place number
  if number > 20 && number < 100 {
    ones_place_digit := number % 10
    tens_place_digit := number - ones_place_digit
    return getSpelledOutNumber(tens_place_digit) + getSpelledOutNumber(ones_place_digit)
  }

  // one-off cases
  if number == 1000 {
    return "onethousand"
  }

  if number == 100 {
    return "onehundred"
  }

  tens_and_ones := number % 100
  number_of_hundreds := (number - tens_and_ones) / 100
  tens_and_ones_string := ""
  if tens_and_ones > 0 {
    tens_and_ones_string = "and" + getSpelledOutNumber(tens_and_ones)
  }
  return getSpelledOutNumber(number_of_hundreds) + "hundred" + tens_and_ones_string
}

func createNumbersToWords() map[int]string {
   numbersToWords := make(map[int]string)
   numbersToWords[1] = "one"
   numbersToWords[2] = "two"
   numbersToWords[3] = "three"
   numbersToWords[4] = "four"
   numbersToWords[5] = "five"
   numbersToWords[6] = "six"
   numbersToWords[7] = "seven"
   numbersToWords[8] = "eight"
   numbersToWords[9] = "nine"
   numbersToWords[10] = "ten"
   numbersToWords[11] = "eleven"
   numbersToWords[12] = "twelve"
   numbersToWords[13] = "thirteen"
   numbersToWords[14] = "fourteen"
   numbersToWords[15] = "fifteen"
   numbersToWords[16] = "sixteen"
   numbersToWords[17] = "seventeen"
   numbersToWords[18] = "eighteen"
   numbersToWords[19] = "nineteen"
   numbersToWords[20] = "twenty"
   numbersToWords[30] = "thirty"
   numbersToWords[40] = "forty"
   numbersToWords[50] = "fifty"
   numbersToWords[60] = "sixty"
   numbersToWords[70] = "seventy"
   numbersToWords[80] = "eighty"
   numbersToWords[90] = "ninety"
   numbersToWords[100] = "hundred"
   numbersToWords[1000] = "thousand"
   return numbersToWords
}

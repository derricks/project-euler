package main

import (
  "testing"
)

func TestCatalogNumbers(test *testing.T) {
  numbers := catalogNumbers(12)

  for number := 0; number < 11; number++ {
      if numbers[number] == ABUNDANT {
        test.Fatalf("%d should not have been abundant", number+1)
      }
  }

  if numbers[11] != ABUNDANT {
    test.Fatal("12 should be abundant")
  }
}

func TestCanBeMadeWithTwoAbundants(test *testing.T) {
  numbers := catalogNumbers(25)

  for number := 3; number < 24; number++ {
    if canBeMadeWithTwoAbundants(number, numbers) {
      test.Fatalf("%d should not be makeable with two abundants", number)
    }
  }

  if !canBeMadeWithTwoAbundants(24, numbers) {
    test.Fatal("24 can be made with two abundants")
  }
}

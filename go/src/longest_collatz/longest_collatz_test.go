package main

import (
  "testing"
)

func TestCollatzSequence(test *testing.T) {
  sequence := collatzSequence(13)

  expected := []int {13,40,20,10,5,16,8,4,2,1}

  for index, value := range expected {
    if sequence[index] != value {
      test.Fatalf("Expected sequence step %d to be %d but was %d", index, expected[index], sequence[index])
    }
  }
}

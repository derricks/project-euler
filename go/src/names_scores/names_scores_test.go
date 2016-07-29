package main

import (
  "testing"
  "reflect"
)

func TestStringToAlphabetIndices(test *testing.T) {
  expected := []int{8, 5, 12, 12, 15}
  actual := stringToAlphabetIndices("HELLO")
  if !reflect.DeepEqual(expected, actual) {
    test.Fatalf("expected %v but got %v", expected, actual)
  }
}

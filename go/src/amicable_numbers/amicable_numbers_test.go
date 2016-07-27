package main

import (
  "testing"
)

func TestAreAmicable(test *testing.T) {
  if !areAmicable(220, 284) {
    test.Fatal("220 and 284 should have been amicable, were not")
  }
}

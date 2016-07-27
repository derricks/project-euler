package main

import (
  "testing"
)

func TestNthTriangle(t *testing.T) {

  firstTenTriangles := []int{1, 3, 6, 10, 15, 21, 28, 36, 45, 55 }

  for test := 0; test < 10; test++ {
    if nthTriangle(test + 1) != firstTenTriangles[test] {
      t.Fatalf("Expected %d got %d", firstTenTriangles[test], nthTriangle(test+1))
    }
  }
}

func TestCountDivisors(t *testing.T) {
  if countDivisors(9) != 3 {
    t.Fatalf("9 should have had 3 divisors, had %d", countDivisors(9))
  }


  if countDivisors(1) != 1 {
    t.Fatalf("1 should have had 1 divisor, had %d", countDivisors(1))
  }

  if countDivisors(28) != 6 {
    t.Fatalf("28 should have had 6 divisors, had %d", countDivisors(28))
  }
}

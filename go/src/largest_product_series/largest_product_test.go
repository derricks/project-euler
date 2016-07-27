package main

import (
  "testing"
)

func TestMultiplySliceVaules(t *testing.T) {
   slice := []int{1,2,3,4,5,6}
   product := multiplySliceValues(slice)

   if product != 720 {
     t.Fatalf("Product should have been 720, but was %d", product)
   }
}

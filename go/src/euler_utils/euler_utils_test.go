/**
 Testing support for nth_prime
 */
package euler_utils

import (
  "testing"
  "reflect"
)

func TestSetSliceToTrueAll(t *testing.T) {
  currentSlice := make([]bool, 6, 6)
  setSliceToTrue(currentSlice, 0)

  for index, value := range currentSlice {
    if !value {
      t.Errorf("Value at index %d is false!", index)
    }
  }
}

func TestSetSliceToTruePartial(t *testing.T) {
  const trueStartingIndex = 3
  currentSlice := make([]bool, 6,6)
  setSliceToTrue(currentSlice, trueStartingIndex)

  for index, value := range currentSlice {
    if value && index < trueStartingIndex - 1 {
      t.Fatalf("Value at %v is true but should be false", index)
    }
  }
}

func TestPrimeSlicesAll(t *testing.T) {
  currentSlice := make([]bool, 800, 800)
  currentSlice[0] = false
  setSliceToTrue(currentSlice, 2)
  populatePrimesInSlice(currentSlice, 2)

  // 1st prime number should be 1
  assertShouldNotBePrime(currentSlice, 1, t)

  // 2 should be a prime number
  assertShouldBePrime(currentSlice, 2, t)

  // 3 should be a prime number
  assertShouldBePrime(currentSlice, 3, t)

  // 4 should not be a prime number
  assertShouldNotBePrime(currentSlice, 4, t)

  // 5 should be a prime number
  assertShouldBePrime(currentSlice, 5, t)

  // 6 should not be prime
  assertShouldNotBePrime(currentSlice, 6, t)

  // 103 should be prime
  assertShouldBePrime(currentSlice, 103, t)

  // 107 should be prime
  assertShouldBePrime(currentSlice, 107, t)
}

func TestIsEven(t *testing.T) {
  if IsEven(1) {
    t.Fatal("1 should have been odd!")
  }

  if !IsEven(2) {
    t.Fatal("2 should have been even!")
  }
}


func TestProductOfSlice(t *testing.T) {
  slice := []int{1,2,3,4,5,6}
  product := ProductOfSliceValues(slice)
  if product != 720 {
    t.Fatalf("Expected 720. Was %d",product)
  }
}

func TestSumOfSlice(test *testing.T) {
  slice := []int{1,2,3,4}
  failOnTest(SumOfSliceValues(slice) == 10, "Sum of {1,2,3,4} should have been 10", test)
}

func TestGenerateDivisors(test *testing.T) {
  failOnTest(reflect.DeepEqual(GenerateDivisors(1), []int{1}), "1's divisors should be {1}", test)
  failOnTest(reflect.DeepEqual(GenerateDivisors(9), []int{1,9,3}), "Wrong divisors for 9", test)
  failOnTest(reflect.DeepEqual(GenerateDivisors(12), []int{1,12,2,6,3,4}), "Wrong divisors for 12", test)
}

func TestGenerateProperDivisors(test *testing.T) {
  failOnTest(reflect.DeepEqual(GenerateProperDivisors(9), []int{1,3}), "9s proper divisors should be 1,3", test)
}


func failOnTest(testResult bool, message string, test *testing.T) {
  if (!testResult) {
    test.Fatal(message)
  }
}

func assertShouldBePrime(primes []bool, index int, t *testing.T) {
  if !primes[index] {
    t.Fatalf("%d should be a prime", index)
  }
}

func assertShouldNotBePrime(primes []bool, index int, t *testing.T) {
  if primes[index] {
    t.Fatalf("%d should not be prime", index)
  }
}

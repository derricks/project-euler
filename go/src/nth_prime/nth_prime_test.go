/**
 Testing support for nth_prime
 */
package main

import (
  "testing"
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

func TestFindNthPrime(t *testing.T) {
  currentSlice := make([]bool, 800, 800)
  setSliceToTrue(currentSlice, 2)
  populatePrimesInSlice(currentSlice, 2)

  assertNthPrime(currentSlice, 1, 2, t)
  assertNthPrime(currentSlice, 2, 3, t)
  assertNthPrime(currentSlice, 3, 5, t)
  assertNthPrime(currentSlice, 4, 7, t)
  assertNthPrime(currentSlice, 5, 11, t)
  assertNthPrime(currentSlice, 6, 13, t)
  assertNthPrime(currentSlice, 7, 17, t)
  assertNthPrime(currentSlice, 8, 19, t)
  assertNthPrime(currentSlice, 9, 23, t)
  assertNthPrime(currentSlice, 10, 29, t)
  assertNthPrime(currentSlice, 11, 31, t)
  assertNthPrime(currentSlice, 12, 37, t)
  assertNthPrime(currentSlice, 13, 41, t)
  assertNthPrime(currentSlice, 14, 43, t)
  assertNthPrime(currentSlice, 15, 47, t)
  assertNthPrime(currentSlice, 16, 53, t)
  assertNthPrime(currentSlice, 17, 59, t)
  assertNthPrime(currentSlice, 18, 61, t)
  assertNthPrime(currentSlice, 19, 67, t)
  assertNthPrime(currentSlice, 20, 71, t)
  assertNthPrime(currentSlice, 21, 73, t)
  assertNthPrime(currentSlice, 22, 79, t)
  assertNthPrime(currentSlice, 23, 83, t)
  assertNthPrime(currentSlice, 24, 89, t)
  assertNthPrime(currentSlice, 25, 97, t)
  assertNthPrime(currentSlice, 26, 101, t)
  assertNthPrime(currentSlice, 27, 103, t)
  assertNthPrime(currentSlice, 28, 107, t)
}

func assertNthPrime(primes []bool, which int, expected int, t *testing.T) {
  if result := findNthPrimeFromSlice(primes, which); result != expected {
    t.Fatalf("Prime number %d should be %d, is actually %d", which, expected, result)
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

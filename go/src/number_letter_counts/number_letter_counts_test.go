package main

import (
  "testing"
)

func TestSpellOutNumber(test *testing.T) {
  if (getSpelledOutNumber(1) != "one" ) {
    test.Fatal("spelled out 1 should be one")
  }

  if (getSpelledOutNumber(342) != "threehundredandfortytwo") {
    test.Fatalf("342 should be threehundredandfortytwo but is %s", getSpelledOutNumber(342))
  }

  if (getSpelledOutNumber(115) != "onehundredandfifteen") {
    test.Fatal("115 should be onehundredandfifteen")
  }
}

func TestGetTotalLettersUsed(test *testing.T) {
  sum := getTotalLettersUsedToNumber(5)
  if (sum != 19) {
    test.Fatalf("Number of letters 1-5 should be 19. Was %v", sum )
  }
}

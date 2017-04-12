package main
// You are given the following information, but you may prefer to do some research for yourself.
//
// 1 Jan 1900 was a Monday.
// Thirty days has September,
// April, June and November.
// All the rest have thirty-one,
// Saving February alone,
// Which has twenty-eight, rain or shine.
// And on leap years, twenty-nine.
// A leap year occurs on any year evenly divisible by 4, but not on a century unless it is divisible by 400.
// How many Sundays fell on the first of the month during the twentieth century (1 Jan 1901 to 31 Dec 2000)?

import (
  "fmt"
  "os"
  "time"
)

// this goes against the spirit of the problem, but does give me exposure to
// the language's time functions
func main() {
  firstSundaysCount := 0

  oneDay, err := time.ParseDuration("24h")

  if err != nil {
    fmt.Println("Invalid duration format")
    os.Exit(1)
  }

  endDate := time.Date(2000, time.December, 31, 23, 59, 0, 0, time.UTC )
  currentDate := time.Date(1901, time.January, 1, 0, 0, 0, 0, time.UTC )

  for currentDate.Before(endDate) {
    if isFirst(currentDate) && isSunday(currentDate) {
      firstSundaysCount += 1
    }
    currentDate = currentDate.Add(oneDay)

  }
  fmt.Println(firstSundaysCount)
}

func isSunday(date time.Time) bool {
  return date.Weekday() == time.Sunday
}

func isFirst(date time.Time) bool {
  return date.Day() == 1
}

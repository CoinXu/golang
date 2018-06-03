package main

import (
  "fmt"
  "time"
)

func main() {
  today := time.Now().Weekday()

  switch time.Saturday {
  case today + 0:
    fmt.Println("Today")
  case today + 1:
    fmt.Println("Tomorrow")
  case today + 2:
    fmt.Println("In tow days")
  default:
    fmt.Println("Too far away")
  }
}
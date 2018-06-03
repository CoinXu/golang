package main

import (
  "fmt"
)

func add(x int, y int) int {
  return x + y
}

func add_a(x, y int) int {
  return x + y
}

func main() {
  fmt.Println(add(42, 13))
  fmt.Println(add_a(42, 13))
}

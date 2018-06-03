package main

import "fmt"

func slip(sum int) (x, y int) {
  x = sum * 4 / 9
  y = sum - x
  return
}

func main() {
  fmt.Println(slip(17)) // 7 10
}
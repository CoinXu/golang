package main

import "fmt"

func main() {
  const total int = 10

  sum := 0
  i := 0

  for ; i < total; i++ {
    sum += i
  }

  fmt.Println(sum)

  for ; sum < 100; {
    sum += 10
  }

  fmt.Println(sum)

  for sum < 200 {
    sum += 10
  }

  fmt.Println(sum)

  for {
    if sum > 300 {
      break
    }
    sum += 10
  }

  fmt.Println(sum)
}
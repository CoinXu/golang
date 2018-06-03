package main

import "fmt"

const Pi float32 = 3.14

func main() {
  const World string = "World"
  fmt.Println("Hello", World)
  fmt.Println("Happy", Pi, "Day")

  const Trust = true
  fmt.Println("Go rules?", Trust)
}
package main

import "fmt"

type Vertex struct {
  X int
  Y int
}

func main() {
  var v Vertex = Vertex{1, 2}
  fmt.Println(v.X)
  fmt.Println(v)

  // struct pointer
  var p *Vertex = &v
  p.X = 1e9
  fmt.Println(v)
}

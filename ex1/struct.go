package main

import "fmt"

type Position struct {
  x int
  y int
}

func main() {
  pos :=Position{1, 1}
  fmt.Printf("%+v\n", pos)
  pos = Position{2, 4}
  fmt.Printf("%#v\n", pos)
  pos = Position{5, 6}
}

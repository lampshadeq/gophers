package main

import "fmt"

const (
  A = iota        // 0
  B = iota * 2    // 2
  C = iota * 5    // 10
  D = iota * iota // 9
)

func main() {
  fmt.Println(A)
  fmt.Println(B)
  fmt.Println(C)
  fmt.Println(D)
}
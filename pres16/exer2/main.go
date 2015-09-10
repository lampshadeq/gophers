package main

import "fmt"

func greatest(list ...int) int {
  great := list[0]
  list = list[1:]
  for _, v := range list {
    if v > great {
      great = v
    }
  }
  return great
}

func main() {
  first := []int{9, 2, 7, 12, 2, 11}
  second := []int{-5, -5, -12, -474}

  fmt.Println(greatest(first...))
  fmt.Println(greatest(second...))
}
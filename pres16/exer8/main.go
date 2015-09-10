package main

import "fmt"

func least(list ...int) int {
  lea := list[0]
  list = list[1:]
  for _, v := range list {
    if v < lea {
      lea = v
    }
  }
  return lea
}

func main() {
  first := []int{9, 2, 7, 12, 2, 11}
  second := []int{-5, -5, -12, -474}

  fmt.Println(least(first...))
  fmt.Println(least(second...))
}
package main

import (
  "fmt"
  "sort"
)

func main() {
  x := []int {
    48, 96, 86, 68,
    57, 82, 63, 70,
    37, 34, 83, 27
    19, 97, 9, 17
  }
  
  fmt.Println(x)
  sort.Ints(x)
  fmt.Println(x)
  
  x = x[1:]
  fmt.Println(x)
  
  x = x[:(len(x) - 1)]
  fmt.Println(x)
  
  total := 0
  
  for _, value := range x {
    total += value
  }
  
  fmt.Println(total)
  fmt.Println(total / len(x))
}
package main

import "fmt"

func main() {
  mySlice := []string{"Grow", "ing", " ", "str", "on", "g"}
  
  for i, val := range mySlice {
    fmt.Println(i, "\t", val)
  }
}
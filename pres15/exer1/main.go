package main

import "fmt"

func main() {
  mySlice := []int{6, 5, 2, 77, 11, -345, 390}
  
  for i, val := range mySlice {
    fmt.Println(i, "\t", val)
  }
}
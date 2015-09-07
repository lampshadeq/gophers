package main

import (
  "fmt"
  "math"
)

func main() {
  var x float64
  fmt.Print("Please enter a float64: ")
  fmt.Scan(&x)
  
  x = math.Ceil(x)
  fmt.Println(x)
}
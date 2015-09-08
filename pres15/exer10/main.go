package main

import "fmt"

func main() {
  var myVar *bool = new(bool)
  fmt.Println(myVar)
  fmt.Println(*myVar)
  
  // True, new returned a pointer
}
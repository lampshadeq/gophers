package main

import "fmt"

func main() {
  var myVar *int = new(int)
  fmt.Println(myVar)
  fmt.Println(*myVar)
  
  // True, new returned a pointer
}
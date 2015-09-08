package main

import "fmt"

func main() {
  var myVar *string = new(string)
  fmt.Println(myVar)
  fmt.Println(*myVar)
  
  // True, new returned a pointer
}
package main

import "fmt"

func main() {
  var myVar map[int]string = make(map[int]string, 10)
  fmt.Println(myVar)
  
  // False, make did not return a pointer
}
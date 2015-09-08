package main

import "fmt"

func main() {
  var myVar []int = make([]int, 10)
  fmt.Println(myVar)
  
  // False, make did not return a pointer
}
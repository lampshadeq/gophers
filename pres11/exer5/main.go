package main

import "fmt"

func main() {
  firstVar := "Heyyyy"
  fmt.Println(firstVar)
  
  var myPoint *string = &firstVar
  fmt.Println(*myPoint)
}
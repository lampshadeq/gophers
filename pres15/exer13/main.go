package main

import "fmt"

func main() {
  var myInt *int = new(int)
  var myString *string = new(string)
  var myBool *bool = new(bool)
  
  fmt.Println("Zeroed value for int:", *myInt)
  fmt.Println("Zeroed value for string:", *myString)
  fmt.Println("Zeroed value for bool:", *myBool)
}
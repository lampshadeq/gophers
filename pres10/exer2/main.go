package main

import "fmt"
import "github.com/lampshadeq/gophers/pres10/exer2/names"

func main() {
  fmt.Println(names.MyName)
  
  // Causes an error
  fmt.Println(yourName)
}
package main

import "fmt"

func main() {
  greeting := func() {
    fmt.Println("Hello ????")
  }
  
  greeting()
  fmt.Printf("%T\n", greeting)
}
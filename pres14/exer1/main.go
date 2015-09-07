package main

import "fmt"

func main() {
  fmt.Print("Hello there user!\n")
  fmt.Print("Here are some numbers:\n")
  for i := 0; i < 10; i++ {
    fmt.Print("\t", i, "\n")
  }
  fmt.Print("Pretty cool, right?")
}
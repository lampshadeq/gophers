package main

import "fmt"

func makeGreeter() func() string {
  return func() string {
    return "Growing strong"
  }
}

func main() {
  tyrells := makeGreeter()
  fmt.Println(tyrells())
}
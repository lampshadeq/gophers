package main

import "fmt"

func hello() {
  fmt.Print("Growing ")
}

func world() {
  fmt.Println("strong")
}

func main() {
  defer world()
  hello()
}
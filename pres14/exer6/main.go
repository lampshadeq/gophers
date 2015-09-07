package main

import (
  "fmt"
  "strconv"
)

func main() {
  i, _ := strconv.Atoi("136")
  s := strconv.Itoa(136)
  
  fmt.Println(i)
  fmt.Println(s)
}
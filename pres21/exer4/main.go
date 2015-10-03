package main

import (
  "strings"
  "fmt"
)

func main() {
  exx := "Growing strong!!!\n"
  ex := strings.Split(exx, " ")
  
  ch := make(chan string, len(ex))
  
  for _, e := range ex {
    ch <- e
  }
  
  close(ch)
  
  for m := range ch {
    fmt.Print(m + " ")
  }
}
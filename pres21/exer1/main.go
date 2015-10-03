package main

import (
  "fmt"
  "time"
)

func main() {
  godur, _ := time.ParseDuration("10ms")
  
  go func() {
    for i := 0; i < 100; i++ {
      fmt.Println("Yo")
      time.Sleep(godur)
    }
  }()
  
  go func() {
    for i := 0; i < 100; i++ {
      fmt.Println("Get shwifty!")
      time.Sleep(godur)
    }
  }()
  
  dur, _ := time.ParseDuration("1s")
  time.Sleep(dur)
}
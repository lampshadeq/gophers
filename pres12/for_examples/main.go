package main

import "fmt"

func main() {
  sum := 0
  for i := 0; i < 100; i++ {
    sum += i
  }
  fmt.Println(sum)
  
  for sum > 0 {
    sum -= 4
  }
  fmt.Println(sum)
  
  for {
    sum *= 2
    
    if sum < -1000 {
      break
    }
  }
  fmt.Println(sum)
}
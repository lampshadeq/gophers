package main

import "fmt"

func main() {
  myNum := make([]int, 5, 10)
  
  for i := 0; i < 5; i++ {
    myNum[i] = i
  }
  
  myNum = append(myNum, 5)
  
  for i, val := range myNum {
    fmt.Println(i, "\t", val)
  }
}
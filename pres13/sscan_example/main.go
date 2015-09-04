package main

import "fmt"

func main() {
  myStr := "15 53"
  var firstNum int
  var secondNum int
  
  fmt.Sscan(myStr, &firstNum, &secondNum)
  fmt.Print(firstNum, secondNum)
}

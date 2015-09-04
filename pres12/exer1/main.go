package main

import "fmt"

func main() {
  // Declare the user variables
  var firstNum int
  var secondNum int
  
  // Get input
  fmt.Scan(&firstNum)
  fmt.Scan(&secondNum)
  
  // Display remainder
  answer := firstNum % secondNum
  fmt.Println(answer)
}
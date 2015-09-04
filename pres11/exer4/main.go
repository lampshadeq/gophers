package main

import "fmt"

func main() {
  var userInput int
  fmt.Scan(&userInput)
  
  userInput = userInput << 5
  
  fmt.Println(userInput)
}
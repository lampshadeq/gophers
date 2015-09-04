package main

import "fmt"

func main() {
  var userInput string
  fmt.Print("Enter your desired string: ")
  fmt.Scanln(&userInput)
  
  userRune := rune(userInput[0])
  fmt.Println(string(userRune))
}
package main

import "fmt"

func main() {
  // Initial
  var littleNumber uint8 = 69
  fmt.Println(&littleNumber)
  fmt.Println(littleNumber)
  fmt.Printf("%b\n", littleNumber)
  
  // Shift left by 3
  littleNumber = littleNumber << 3
  fmt.Println(&littleNumber)
  fmt.Println(littleNumber)
  fmt.Printf("%b\n", littleNumber)
  
  // Shift right by 2
  littleNumber = 69
  littleNumber = littleNumber >> 2
  fmt.Println(&littleNumber)
  fmt.Println(littleNumber)
  fmt.Printf("%b\n", littleNumber)
  
  // AND with mask
  littleNumber = 69
  var mask uint8 = 241
  littleNumber = littleNumber & mask
  fmt.Println(&littleNumber)
  fmt.Println(littleNumber)
  fmt.Printf("%b\n", littleNumber)
}
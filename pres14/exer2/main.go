package main

import "fmt"

func main() {
  str1 := "This is a string."
  str2 := "Here is another."
  str3 := "æ"
  
  fmt.Println("Len (bytes):", len(str1))
  fmt.Println("Len (bytes):", len(str2))
  fmt.Println("Len (bytes):", len(str3))
}
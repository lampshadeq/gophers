package main

import "fmt"

func main() {
  firstSlice := []string{"Gro", "wing", " "}
  secondSlice := []string{"s", "tr", "on", "g"}
  
  firstSlice = append(firstSlice, secondSlice...)
  fmt.Println(firstSlice)
  
  firstSlice = append(firstSlice[:2], firstSlice[3:]...)
  fmt.Println(firstSlice)
}
package main

import "fmt"

func main() {
  var mySlice *[]byte = new([]byte)
  var myMap *map[int]int = new(map[int]int)
  
  fmt.Println("Zeroed value for slice:", *mySlice)
  fmt.Println("Zeroed value for map:", *myMap)
}
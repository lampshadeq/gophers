package main

import "fmt"

func main() {
  // Create a map with shorthand notation
  myMap := map[string]int {
    "Sand": 124,
    "Water": 11,
    "Brick": 77,
    "Grass": 90,
    "Portal": 100,
  }
  
  // Add an entry
  myMap["Stone"] = 69
  
  // Change an entry
  myMap["Water"] = 5
  
  // Delete an entry
  delete(myMap, "Brick")
  
  // Print all the entries with range
  for key, val := range myMap {
    fmt.Println(key, "\t", val)
  }
  
  // Print the length of the map
  fmt.Println(len(myMap))
  
  // Use comma ok idiom
  if _, exist := myMap["Portal"]; exist {
    fmt.Println("Portal exists!")
  } else {
    fmt.Println("Portal does not exist... :(")
  }
}